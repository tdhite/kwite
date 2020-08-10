# Knative Exmaple of a Kwite Service
This example provides a simple example of deploying a Kwite via [Knative
Service](https://knative.dev/docs/serving/knative-kubernetes-services/) custom
resource declaration. There are two items of interest, the ConfigMap and the
Knative Service. Both are present in the [example
manifest](knative-kwite.yaml).

## The ConfigMap
Kwites pull certain configurations, for example the template to render, from
the /configs directory. In order to create the relevant files, a
[ConfigMap](../kubernetes/base/configmap.yaml) is mounted as a volume in the
Kwite container.

When using [Kwite-operator](https://github.com/tdhite/kwite-operator) to manage
Kwite deployments, all necessary Kubernetes resources like ConfigMaps and
related [Volume](https://kubernetes.io/docs/concepts/storage/volumes/) mounts
are automatically created and managed by that operator implementation.

When using Knative, no such operator implementation exists. In this example,
the [Knative Serving](https://knative.dev/docs/serving/) component is used to
manage Kwites. However, Knative only creates and manages the custom resources
specified in its own CRDs and they have no knowledge at all about the nature of
what is necesary to configure Kwites.

Long story short, one must create and manage, independently of Knative,
ConfigMaps that contain Kwite configurations needed by Kwites served up by
Knative Serving. As well, the Knative Service declaration needs the Volume and
related mount information.

## The Knative Service
The Knative Service example herein contains the necessary info to run a Kwite
container and serve up the templates provided in the ConfigMap. Take a look at
the [example manifest](knative-kwite.yaml), particularly the volumeMounts and
volumes statements.

## Try it out
To run the example, you need a valid Kubernetes cluster with [Knative
installed](https://knative.dev/docs/install/). Once running and you have
[kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) setup and
running with your cluster, just run the following command and all should be
working with your first Kwite deployed via Knative.

    kubectl apply -f knative-kwite.yaml

After that, your Kwite should be available at
`http://your_cluster_address/mykwite`. For example, execute the following
command.

    curl -H 'content-type: application/json' -d '{"x": 2}' https://kubernetes.io/docs/tasks/tools/install-kubectl/mykwite.

To find the actual address (to replace `your\_cluster\_address`, run this
command and it will show up under the kwite-1 entry, field URL.

