#? Create Link (Invalid Data)
#* Req
PUT {{base_url}}
Content-Type: application/json
{ "url": "123" }
#* Res
HTTP 400

#? Access Link (Not Found)
#* Req
GET {{base_url}}/123
#* Res
HTTP 404

#? Create Link (OK)
#* Req
PUT {{base_url}}
Content-Type: application/json
{ "url": "{{target_url}}" }
#* Res
HTTP 200
[Asserts]
jsonpath "$.path" isString
jsonpath "$.target" == "{{target_url_normalized}}"
jsonpath "$.active" == true
jsonpath "$.secret" isString
[Captures]
path: jsonpath "$.path"
secret: jsonpath "$.secret"

#? Access Link (OK)
#* Req
GET {{base_url}}/{{path}}
#* Res
HTTP 301
[Asserts]
header "Location" == "{{target_url_normalized}}"

#? Disable Link (Not Found)
#* Req
PATCH {{base_url}}/123/disable
{ "secret": "{{secret}}" }
#* Res
HTTP 404

#? Disable Link (Wrong Secret)
#* Req
PATCH {{base_url}}/{{path}}/disable
{ "secret": "123" }
#* Res
HTTP 404

#? Disable Link (OK)
#* Req
PATCH {{base_url}}/{{path}}/disable
{ "secret": "{{secret}}" }
#* Res
HTTP 200
[Asserts]
jsonpath "$.path" isString
jsonpath "$.target" == "{{target_url_normalized}}"
jsonpath "$.active" == false
jsonpath "$.secret" isString

#? Access Link (OK)
#* Req
GET {{base_url}}/{{path}}
#* Res
HTTP 404

#? Disable Link (Idempotency Check)
#* Req
PATCH {{base_url}}/{{path}}/disable
{ "secret": "{{secret}}" }
#* Res
HTTP 200
[Asserts]
jsonpath "$.path" isString
jsonpath "$.target" == "{{target_url_normalized}}"
jsonpath "$.active" == false
jsonpath "$.secret" isString

#? Access Link (Idempotency Check)
#* Req
GET {{base_url}}/{{path}}
#* Res
HTTP 404
