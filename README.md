# Kwite
A quite small web service project that depends on
[Kubernetes](https://kubernetes.io) for
[URL](https://www.w3.org/Addressing/URL/url-spec.txt) routing and scale.

NOTE: This is a work in progress, so don't expect all things to work yet! There
are also many more capabilities to add, such is HTTPS support, more functions
and whatnot. They are coming soon enough.

## Overview
The original concept of this project was to create exactly the above --
basically a Kubernetes as a web server kind of thing that was quite small yet
quite powerful. What a terrible name would be Quite Small Web Server! So a good
name was necessary. Since Kubernetes was the target, 'K' always seems the
thing, so it is now _Kwite_ -- Kubernetes Web Integrated Template Engine. [Look
here](docs/whykwite.md) for the real background on that name.

### What Are Kwites? 
The basic premise of Kwites are to build a web server where Kubernetes itself
is the actual workhorse server the way one might think of with [Apache HTTP
Server Project](https://httpd.apache.org). While not an exact comparison, the
idea is that a Kwite represents a single
[URL](https://www.w3.org/Addressing/URL/url-spec.txt).

Kwites are microservices that only respond to one, and only one URL. For the
reasoning and backdrop on why only one URL, [read more here](docs/kwites.md).

A small sample Kwite template is available in a [sample
ConfigMap](examples/kubernetes/base/configmap.yaml), which is part of the
[kustomize](https://kustomize.io) based [example Kubernetes
manifests](examples/kubernetes) for deploying a Kwite. See the [ConfigMap
documentation](docs/configmap.md) for a description of the fields. Most,
however, will use the
[Kwite-operator](https://github.com/tdhite/kwite-operator) to manage Kwite
deployments.

In any event, note in the example template that it is using a variable ".x" to
fill it out. One way to call such a Kwite (assuming the Kwite microservice
address provided) might be:

    curl -H 'Content-Type: application/json' -d '{"x": 2}' http://mykwite.mynamespace.svc.cluster.local:8080/kwite

### Using Data in Kwite Templates
It should be noted from the example above that passing data for use in Kwite
templates is always via a [JSON](https://www.json.org/json-en.html) [HTTP
Message Body](https://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html).

Because Kwites are just Go templates under the hood, with lots of [callable
functionality added](docs/funcs.md), web programmers can accomplish quite a bit
of what they might want with a Kwite. That includes creating [entire
sites](docs/kwites.md#kwites-as-sites).

For example, a template similar to that in the sample below, which conforms to
[Kwite-operator](https://github.com/tdhite/kwite-operator), would deliver the
content of the first Kwite's template followed thereafter by the content
returned by a second Kwite at the URL specified in the httpGet call:

```yaml
apiVersion: web.kwite.io/v1beta1
kind: Kwite
metadata:
  name: kwite-1
spec:
  url: "/kwite"
  port: 8080
  image: "concourse.corp.local/kwite:0.0.11"
  targetcpu: 50
  minreplicas: 1
  maxreplicas: 10
  template: |
    this is a kwite template x is {{ .x }}
    {{ httpGet "http://kwite-2.kwiteop-system.svc.cluster.local:8081/kwite" "{\"x\": 2}" }}
  ready: "OK!"
  alive: "OK!"
```

### Added Functionality in Kwite Templates
Kwites inject additional functionality into the Go templates by way of
additional Go methods. The functions all are directly callable rather accessed
from the underlying Go template '.' operator, which would hide them from
validating parsing.

For the most part, users need not fully understand this, but by providing only
top level functions, Kwites allow other systems to more fully validate templates
before admitting them. For example, the
[Kwite-operator](https://github.com/tdhite/kwite-operator) validates Kwite
templates via its [Admission
Webhook](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#what-are-admission-webhooks).
That operator won't even allow a Kwite into a Kubernetes cluster if it will not
compile, including the extended functionality Kwite templates contain.

### Kwite Self-Reformation (future)
To close the loop of Kwites hopefully being less nuts-ish than at first blush,
Kwites work with the
[Kwite-operator](https://github.com/tdhite/kwite-operator)
project to auto-resolve URLs within templates so as to 'self-form' overall
sites easily for web developers. For example, where a URL in a template httpGet
call takes the form of `httpGet "kwite://thisotherkwite.thatnamespace" ""`, the
Kwite will automatically convert that to the full cluster address appropriate
for "thisotherkwite" in "thatnamespace" and issue the HTTP GET appropriately.
That might resolve to something like
`http://mykwite.thatnamespace.svc.cluster.local:8080/thisotherkwite`.

## Try it out
To try out Kwite, it must be built and is usually executed in a suitable
[Kubernetes](https://kubernetes.io) environment. Details are [further
below](#build-and-run).

For testing, you can also [run Kwites independently](#running-standalone).

### Prerequisites
There are some basic requirements in order build and use Kwite
microservices:

1. A Kubernetes cluster sufficient to run kwite instances;
1. A container registry such as [Harbor](https://goharbor.io) and relevant
   credentials for pushing and pulling containers;
1. A [Concourse](https://concourse-ci.org) or similar setup, for example
   [Argo CD](https://argoproj.github.io/argo-cd/), in order to run tests, build
   and deploy Kwites.

### Kubernetes Cluster
Deploying Kwites targets a Kubernetes cluster. They will deploy into the
default namespace unless otherwise specified.

In addition, as with any Kubernetes cluster, credentials should be setup.
Secrets for pushing and pulling containers from the Docker registry should
exist before building the microservice. One way to add such credentials is a
command simillar to the following:

    kubectl create secret docker-registry kwite-registry-creds --docker-server=<your-registry-server> --docker-username=<your-name> --docker-password=<your-pword> --docker-email=<your-email>

appropriately setup for your own registry and deployment setup.

### Build and Run
Building the code involves very few steps and there are some CI/CD sample setup
options to help out.

This project includes pipeline and related task declarations in the
[build/ci](build/ci) directory. Those can be used assuming Concourse is
available. See the [build/ci/README](build/ci/README.md) file for details.
Adapt those as needed for your own CI/CD setup.

#### Run In-Cluster
So long as the pipeline runs successfully the Kwite container will be pushed to
the container image registry specified in the
[params.yaml](build/ci/examples/params.yaml) file modified to match the local
environment. As well, a sample Kwite will be deployed via
[kustomize](https://github.com/kubernetes-sigs/kustomize) from
[examples/kubernetes](examples/kubernetes).

### Running Standalone
The microservice will run standalone for testing. To do that just run something
of the form:

    make
    cmd/kwite/kwite -c ./examples/configs --port 8081

Thereafter, use something like curl to access it, as in:

    curl -H 'Content-Type: application/json' -d '{"x": 2}' http://localhost:8081/kwite

## Contributing

The kwite project team welcomes contributions from the community. Before you
start working with kwite, please read our [Developer Certificate of
Origin](https://cla.vmware.com/dco). All contributions to this repository must
be signed as described on that page. Your signature certifies that you wrote
the patch or have the right to pass it on as an open-source patch. For more
detailed information, refer to [CONTRIBUTING.md](CONTRIBUTING.md).

## License and Copyright
Copyright: Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: [https://spdx.org/licenses/MIT.html](https://spdx.org/licenses/MIT.html)
