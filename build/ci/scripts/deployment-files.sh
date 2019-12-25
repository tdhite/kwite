#!/bin/bash
# kwite deployment-files.sh
#
# Copyright (c) 2019-2020 VMware, Inc.
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

set -e -x

echo "------Environment Variables------"
set

# Save current directory
TOP="$(pwd)"

ret=0
if [ -z "${container}" ]; then
    echo "ERROR: container not supplied. Aborting!"
    ret=1
fi
if [ $ret -ne 0 ]; then
    exit $ret
fi

tag="$(cat version/version)"

# Make the output area if it does not exist
mkdir -p ${TOP}/kubernetes

# create the kubernetes deployment and service manifests
cp -a sources/examples/kubernetes/base ${TOP}/kubernetes/
cp -a sources/examples/kubernetes/overlays ${TOP}/kubernetes/
cat >${TOP}/kubernetes/base/kustomization.yaml <<EOF
resources:
- configmap.yaml
- deployment.yaml
- service.yaml
images:
- name: kwite:latest
  newTag: ${tag}
  newName: ${container}
EOF

# show the kustomization for potential debugging
cat ${TOP}/kubernetes/base/kustomization.yaml

# Check what's here
echo "List out the output directory:"
ls -laRt ${TOP}/kubernetes
echo ""
