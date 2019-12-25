# Kwite ConfigMap Settings
Kwite instance make use of a ConfigMap for the purposes of configuration. While
it is possible to specify a container that has the proper built-in file system
Kwites expect, it is much easier to update Kwite templates by updating the
ConfigMap.

The settings are as follows:

* `metadata.name`:
This can be set to any name consistent with Kubernetes naming standards. The
name uniquely identifies the ConfigMap within a Kubernetes namespace. For
example, `kwite-1`. Note that the name of the ConfigMap will normally be
referenced in the Pod spec in order to mount the map on the Kwite containers'
filesystems.

* `data.url`:
The path portion of the URL to which the kwite will respond. For example, a url
of `/kwite` would cause the Kwite to respond to
http://\<cluster-address\>/kwite.

* `data.ready`:
The [Go template](https://golang.org/pkg/text/template/) that the Kwite should
execute to determine if it is ready to begin servicing inbound HTTP requests.
The Kwite will make available `/kwiteready` as the path portion of the probe
URL. For details, see Kubernetes
[probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/).
See also the [Kwite documentation](kwites.md) regarding its use of Go
templating.

* `data.alive`:
The [Go template](https://golang.org/pkg/text/template/) that the Kwite should
execute to determine if it is still alive.
The Kwite will make available `/kwitealive` as the path portion of the probe
For details, see Kubernetes
[probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/).
See also the [Kwite documentation](kwites.md) regarding its use of Go
templating.

* `data.template`:
The [Go template](https://golang.org/pkg/text/template/) that the Kwite should
execute as the response to HTTP requests on the Kwite. See also the [Kwite
documentation](kwites.md)
regarding its use of Go templating.
