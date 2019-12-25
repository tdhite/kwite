#!/bin/bash
# kwite build.sh
#
# Copyright (c) 2019-2020 VMware, Inc.
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

set -e -x

# Save current directory
TOP="$(pwd)"

# Show current setup
echo "GOPATH is: " $GOPATH
echo "TOP is: " $TOP
echo ""

# Build the beast
cd sources
make cmd/kwite/kwite

# Check static linked binary
echo "Check static link status:"
if ldd cmd/kwite/kwite; then
    echo "The kwite binary is dynamically linked, cannot use it."
    exit 1
fi

# Copy build artifacts to the output directory
cp -a cmd/kwite/kwite ${TOP}/build/
