#!/bin/bash

source .buildkite/scripts/common.sh

set -euo pipefail

add_bin_path
with_mage

mage -v check

check_git_diff

if mage isElasticPackageGreaterThan 0.113.0 ; then
    # links require at least v0.113.0
    use_elastic_package
    ${ELASTIC_PACKAGE_BIN} links check
fi
