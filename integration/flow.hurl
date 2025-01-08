#? Create Link (Invalid Data)
#* Req
PUT {{base_url}}
Content-Type: application/json
{ "url": "{{wrong_data}}" }
#* Res
HTTP 400

#? Access Link (Not Found)
#* Req
GET {{base_url}}/{{wrong_data}}
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
PATCH {{base_url}}/{{wrong_data}}/disable
{ "secret": "{{secret}}" }
#* Res
HTTP 404

#? Disable Link (Wrong Secret)
#* Req
PATCH {{base_url}}/{{path}}/disable
{ "secret": "{{wrong_data}}" }
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

#? Access Link (Disabled)
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

#? Access Link (Disabled Idempotency Check)
#* Req
GET {{base_url}}/{{path}}
#* Res
HTTP 404

#? Edit Link (Disabled)
#* Req
PATCH {{base_url}}/{{path}}/edit
{ "secret": "{{secret}}", "url": "{{edit_url}}" }
#* Res
HTTP 404

#? Enable Link (Wrong Secret)
#* Req
PATCH {{base_url}}/{{path}}/enable
{ "secret": "{{wrong_data}}" }
#* Res
HTTP 404

#? Enable Link (OK)
#* Req
PATCH {{base_url}}/{{path}}/enable
{ "secret": "{{secret}}" }
#* Res
HTTP 200
[Asserts]
jsonpath "$.path" isString
jsonpath "$.target" == "{{target_url_normalized}}"
jsonpath "$.active" == true
jsonpath "$.secret" isString

#? Access Link (Enabled)
#* Req
GET {{base_url}}/{{path}}
#* Res
HTTP 301
[Asserts]
header "Location" == "{{target_url_normalized}}"

#? Enable Link (Idempotency Check)
#* Req
PATCH {{base_url}}/{{path}}/enable
{ "secret": "{{secret}}" }
#* Res
HTTP 200
[Asserts]
jsonpath "$.path" isString
jsonpath "$.target" == "{{target_url_normalized}}"
jsonpath "$.active" == true
jsonpath "$.secret" isString

#? Access Link (Enabled Idempotency Check)
#* Req
GET {{base_url}}/{{path}}
#* Res
HTTP 301
[Asserts]
header "Location" == "{{target_url_normalized}}"

#? Edit Link (Wrong Secret)
#* Req
PATCH {{base_url}}/{{path}}/edit
{ "secret": "{{wrong_data}}", "url": "{{edit_url}}" }
#* Res
HTTP 404

#? Edit Link (Invalid Data)
#* Req
PATCH {{base_url}}/{{path}}/edit
{ "secret": "{{secret}}", "url": "{{wrong_data}}" }
#* Res
HTTP 400

#? Edit Link (OK)
#* Req
PATCH {{base_url}}/{{path}}/edit
{ "secret": "{{secret}}", "url": "{{edit_url}}" }
#* Res
HTTP 200
[Asserts]
jsonpath "$.path" isString
jsonpath "$.target" == "{{edit_url_normalized}}"
jsonpath "$.active" == true
jsonpath "$.secret" isString

#? Access Link (Edited)
#* Req
GET {{base_url}}/{{path}}
#* Res
HTTP 301
[Asserts]
header "Location" == "{{edit_url_normalized}}"
