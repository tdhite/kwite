#!/bin/bash
#
# Copyright (c) 2019-2020 VMware, Inc.
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

CONTAINER=${CONTAINER:-kwite}

# Grab the latest build output
cp -a ../../cmd/kwite/kwite .

# Build and push the container
docker build --rm -t "${CONTAINER}" .

# Cleanup
rm -f kwite

