POST {{baseUrl}}/admin
Content-Type: application/json
[Options]
variable: value={{newUuid}}
[BasicAuth]
foo:bar
{ "value": "{{value}}" }

HTTP 200
[Asserts]
jsonpath "$.status" == "ok"

###

GET {{baseUrl}}/user/foo

HTTP 200
[Asserts]
jsonpath "$.user" == "foo"
jsonpath "$.value" == "{{value}}"
