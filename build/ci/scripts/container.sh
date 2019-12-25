#!/bin/bash
# kwite container.sh
#
# Copyright (c) 2019-2020 VMware, Inc.
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

set -e -x

# Save current directory
TOP="$(pwd)"

# Show current setup
echo "TOP is: " $TOP
echo ""

# Copy build results and static content to the output
echo "Copy build artifacts to output"
cp -a ${TOP}/build/kwite ${TOP}/container/

# Copy the docker build and ancillary files
cp -a ${TOP}/sources/build/image/Dockerfile ${TOP}/container/
cp -a ${TOP}/sources/build/image/etc ${TOP}/container/

# List what got laid down
echo "List out the container directory"
ls -laRt ${TOP}/container
echo ""
