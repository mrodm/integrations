// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package citools

import (
	"fmt"
	"os"

	"github.com/Masterminds/semver/v3"
	"golang.org/x/mod/modfile"
)

func PackageVersionGoMod(modulePath string) (*semver.Version, error) {
	contents, err := os.ReadFile("go.mod")
	if err != nil {
		return nil, fmt.Errorf("failed to read file go.mod: %w", err)
	}
	f, err := modfile.Parse("go.mod", contents, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse go.mod: %w", err)
	}
	epVersion := ""
	for _, r := range f.Require {
		if r.Indirect {
			continue
		}
		if r.Mod.Path != modulePath {
			continue
		}
		epVersion = r.Mod.Version
	}
	if epVersion == "" {
		return nil, fmt.Errorf("not found elastic-package dependency")
	}

	foundVersion, err := semver.NewVersion(epVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to parse version %q: %w", epVersion, err)
	}
	return foundVersion, nil
}
