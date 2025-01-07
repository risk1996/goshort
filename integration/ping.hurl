GET {{baseUrl}}/ping

HTTP 200
[Asserts]
body == "pong"
