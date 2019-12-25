# Kwite Functions
Kwites allow the use of various functions within templates. As Go [template
based](https://golang.org/pkg/text/template/) all the capabilities of that
package are automatically available.

However, Kwites add functionality from the core Go packages for use as well as
various additional functions the [Kwite](https://github.com/tdhite/kwite)
itself exposes.

## Kwite HTTP Functions
Kwite exposes [REST](https://restfulapi.net) helper functions. These are as
follows:

### httpGet url json headers...
httpGet initiates a [HTTP GET](https://restfulapi.net/http-methods/#get) call.
Parameters are as follows:

| Parameter | Description |
| --- | --- |
| url | is a string value specifying the HTTP URL to call. |
| json | The [HTTP BODY](https://en.wikipedia.org/wiki/HTTP_message_body) data, always JSON format and can be an empty string (""). |
| headers | A list of HTTP Header pairs to include. These must be two strings per header, such as: "Content-type" "application/json" and there can be any numbrer of such pairs. |

For example:
    httpGet "http://mykwite.myns.svc.cluster.local:8080/mykwite" "{\"x\": 2}" "Content-type" "application/json" "Accept" "text/plain"

### httpDelete url json
httpDelete initiates a [HTTP DELETE](https://restfulapi.net/http-methods/#delete)
call. Parameters are as follows:

| Parameter | Description |
| --- | --- |
| url | is a string value specifying the HTTP URL to call. |
| json | The [HTTP BODY](https://en.wikipedia.org/wiki/HTTP_message_body) data, always JSON format and can be an empty string (""). |
| headers | A list of HTTP Header pairs to include. These must be two strings per header, such as: "Content-type" "application/json" and there can be any numbrer of such pairs. |

For example:
    httpDelete "http://mykwite.myns.svc.cluster.local:8080/mykwite" "{\"x\": 2}" "Content-type" "application/json" "Accept" "text/plain"

### httpPatch url json
httpPatch initiates a [HTTP PATCH](https://restfulapi.net/http-methods/#patch)
call. Parameters are as follows:

| Parameter | Description |
| --- | --- |
| url | is a string value specifying the HTTP URL to call. |
| json | The [HTTP BODY](https://en.wikipedia.org/wiki/HTTP_message_body) data, always JSON format and can be an empty string (""). |
| headers | A list of HTTP Header pairs to include. These must be two strings per header, such as: "Content-type" "application/json" and there can be any numbrer of such pairs. |

For example:
    httpPatch "http://mykwite.myns.svc.cluster.local:8080/mykwite" "{\"x\": 2}" "Content-type" "application/json" "Accept" "text/plain"

### httpPatch url json
httpPost initiates a [HTTP POST](https://restfulapi.net/http-methods/#post)
call. Parameters are as follows:

| Parameter | Description |
| --- | --- |
| url | is a string value specifying the HTTP URL to call. |
| json | The [HTTP BODY](https://en.wikipedia.org/wiki/HTTP_message_body) data, always JSON format and can be an empty string (""). |
| headers | A list of HTTP Header pairs to include. These must be two strings per header, such as: "Content-type" "application/json" and there can be any numbrer of such pairs. |

For example:
    httpPost "http://mykwite.myns.svc.cluster.local:8080/mykwite" "{\"x\": 2}" "Content-type" "application/json" "Accept" "text/plain"

## String Functions
Kwites have access to most of the [Go strings](https://golang.org/pkg/strings)
package. In particular, Kwites can call all functions from the strings package
other than those related to strings.Builder, strings.Reader or those having a
function as a parameter (e.g., FieldFunc).

In order to call those functions, however, the prefix "str" must be appended to
the function call. This is done to prevent any potential name collisions, while
allowing direct access to the strings package, in future additions to the
[Kwite](https://github.com/tdhite/kwite) project.

For example, the following template would return `1` and that's it.

    {{ strCompare "bbb" "aaa" }}

Note also that because strNewReplacer returns a valid Go type with associated
methods, the methods supported by that type need no prefixes. For example, the
following is a valid template:

    {{ $r := strNewReplacer "<" "&lt;" ">" "&gt;" }}
    {{- $r.Replace "This is some italicized <i>HTML</i>!" }}

See also [../examples/configs/template](https://github.com/tdhite/kwite/blob/master/examples/configs/template).

## Math Functions
Kwites have access to most of the [Go math](https://golang.org/pkg/math)
package. In particular, Kwites can call all math functions other than Frexp,
Lgamma, Modf and Sincos.

In addition, Kwites can obtain any value from the [math
constants](https://golang.org/pkg/math/#pkg-constants) by calling a function
with the same name as a constant. The functions returns the value of the
constant. At this time, the "limit value" constants are not included.

For example, this template will work to print the value of sin(pi):

    The value of the sin(pi) is {{ Sin Pi }}.

For more examples, see [../examples/configs/template](https://github.com/tdhite/kwite/blob/master/examples/configs/template).
