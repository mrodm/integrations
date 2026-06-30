// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package packages

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/elastic/integrations/dev/citools"
)

// DetectPackages maps a list of changed file paths to the package names they belong to.
// packagesDir is the path to the packages/ directory (e.g. "packages").
//
// It first calls citools.ListPackages to discover all package roots, then checks
// each changed file against those roots. This handles both flat structures
// (packages/<name>/) and nested structures (packages/<technology>/<name>/) without
// any custom tree-walking logic. Files not under any known package root are silently
// skipped.
//
// Returns a deduplicated list of package names in the order they are first encountered.
func DetectPackages(files []string, packagesDir string) ([]string, error) {
	pkgPaths, err := citools.ListPackages(packagesDir)
	if err != nil {
		return nil, fmt.Errorf("listing packages in %s: %w", packagesDir, err)
	}

	type pkg struct {
		prefix string // path/to/package/ (with trailing separator)
		name   string
	}
	known := make([]pkg, 0, len(pkgPaths))
	for _, p := range pkgPaths {
		manifest, err := citools.ReadPackageManifest(filepath.Join(p, citools.ManifestFileName))
		if err != nil {
			return nil, fmt.Errorf("reading manifest for %s: %w", p, err)
		}
		known = append(known, pkg{
			prefix: p + string(filepath.Separator),
			name:   manifest.Name,
		})
	}

	seen := make(map[string]struct{})
	var result []string
	for _, f := range files {
		for _, k := range known {
			if strings.HasPrefix(f, k.prefix) {
				if _, ok := seen[k.name]; !ok {
					seen[k.name] = struct{}{}
					result = append(result, k.name)
				}
				break
			}
		}
	}
	return result, nil
}
