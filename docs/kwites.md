# What are Kwites?
Kwites are microservices that execute, depending on HTTP request headers,
either [Go text templates](https://golang.org/pkg/text/template/) or [Go HTML
templates](https://golang.org/pkg/html/template/). Briefly, Kwite microservice
instances are web servers that return the results of the template execution,
but respond to one and only one URL. That might seem a bit odd at first blush
considering most [HTTP](https://www.w3.org/Protocols/) web servers respond to a
myriad of URLs.

Some reasons much larger web servers, like the
[Apache](https://httpd.apache.org), [Nginx](https://nginx.org) and others,
respond to many URLs might be to allow site administrators: to centralize the
content (URL responses) they serve; and vertically scale the servers so each
server can handle additional traffic load and content size. This Kwite project
and its sister
[Kwite-operator](https://github.com/tdhite/kwite-operator) seek
to have [Kubernetes](https://kubernetes.io) do all that kind of work by scaling
Kwites (effectively individual URLs) automatically and appropriately, thus
obviating the need for more traditional web servers. Kwites, in this way,
provide (indirectly via Kubernetes) very fine grained, even if automatic,
scaling.

## Kwites as Sites
Does all that above seem a crazy? If so, consider this: a Kwite is implemented
via a Go template that is infused with [many capabilities](funcs.md), including
the ability to retrieve or post data to other URLs directly within the
template.  So, indirectly, a single, correctly programmed Kwite can serve
content from many other Kwites (pages) from a single URL, though with site
navigation formed via the JSON message body, not the content and structure of,
for example, a file system that makes up the multi-page site on a traditional
server.
