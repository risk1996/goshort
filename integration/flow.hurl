#? Create Link
#* Req
POST {{base_url}}
Content-Type: application/json
{ "url": "https://www.google.com/" }
#* Res
HTTP 201
[Asserts]
jsonpath "$.id" > 0
jsonpath "$.path" isString
jsonpath "$.target" == "https://www.google.com"
[Captures]
path: jsonpath "$.path"

#? Access Link
#* Req
GET {{base_url}}/{{path}}
#* Res
HTTP 301
[Asserts]
header "Location" == "https://www.google.com"
