#? Create Link
POST {{baseUrl}}
#* Req
Content-Type: application/json
{ "url": "https://www.google.com/" }
#* Res
HTTP 201
[Asserts]
jsonpath "$.id" > 0
jsonpath "$.path" isString
jsonpath "$.target" == "https://www.google.com"
