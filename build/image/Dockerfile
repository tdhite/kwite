# kwite container image
#
# Copyright (c) 2019-2020 VMware, Inc.
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#
FROM scratch
LABEL maintainer="Tom Hite <thite@vmware.com>"
COPY --chown=65534:65534 kwite /kwite
ADD etc /etc
USER 65534:65534
EXPOSE 8080
ENTRYPOINT ["/kwite"]
