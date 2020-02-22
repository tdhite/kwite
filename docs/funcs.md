# Kwite Functions
Kwites allow the use of various functions within templates. As Go [template
based](https://golang.org/pkg/text/template/) all the capabilities of that
package are automatically available.

However, Kwites add functionality from the core Go packages for use as well as
various additional functions the [Kwite](https://github.com/tdhite/kwite)
itself exposes.

## HTTP Functions
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

## JSON Functions
Kwites have access to JSON encoding functionality. These are wrappers around
various calls from the [Go JSON
package](https://golang.org/pkg/encoding/json/).

### jsonHTMLEscape jsonData
Escapes JSON data for HTML to prepare it for \<script\> tags via wrapper around
the [Go language json.HTMLEscape](https://godoc.org/encoding/json#HTMLEscape)
method.

| Parameter | Description |
| --- | --- |
| jsonData | A string value containing the *valid* JSON. Use [jsonValid](#jsonValid-jsonData) to assure valid jSON. |

For example:
    $myJSON := "{\\"c\\": \\"This is <b>bold</b> and this is not\\"}"
    jsonHtmlEscape $myJSON

### jsonIndent jsonData prefix indent
Returens intended JSON (string) via wrapper around the [Golang
json.Indent](https://godoc.org/encoding/json#Indent) method.  func

| Parameter | Description |
| --- | --- |
| jsonData | A string value containing the *valid* JSON. Use [jsonValid](#jsonValid-jsonData) to assure valid jSON. |

### jsonToInterface jsonData
Unmarshals JSON string into a Go type (e.g., map[string]interface{}) via
wrapper around [json.Marshal](https://godoc.org/encoding/json#Marshal).

| Parameter | Description |
| --- | --- |
| jsonData | A string value containing the *valid* JSON. Use [jsonValid](#jsonValid-jsonData) to assure valid jSON. |

For example:
    $myJSON := "{\\"c\\": \\"This is <b>bold</b> and this is not\\"}"
    $map := jsonToInterface $myJSON

would create a Golang [map](https://golang.org/pkg/go/types/#Map). From there
access within a template is as normal, such as {{ $myJSON.c | jsonHTMLEscape}}.

### jsonToString jsonVariable
Returns a string containing the JSON representing the value of jsonVariable.

| Parameter | Description |
| --- | --- |
| jsonVariable | A variable (any Go type) to be marshalled to a JSON encoded string. |

For example:
    $jsonString := jsonToString .

would return the JSON representation of the template "." variable.

### jsonValid jsonData
Reports whether a string is valid JSON via wrapper around
[json.Valid](https://godoc.org/encoding/json#Valid)

| Parameter | Description |
| --- | --- |
| jsonData | A string value containing the JSON to validate. |

For example:
    $myJSON := "{\\"c\\": \\"This is <b>bold</b> and this is not\\"}"
    $isValid := jsonValid $myJSON
