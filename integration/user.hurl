GET {{baseUrl}}/user/invalid

HTTP 200
[Asserts]
jsonpath "$.user" == "invalid"
jsonpath "$.status" == "no value"
