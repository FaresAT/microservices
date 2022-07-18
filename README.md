# URL Shortener created using Gin
Small demo playing around with gin's routing features, and is purely a backend service (no templating or webpage routing)
Currently no tests - plan to add soon.

### Example Usage:

``
curl --request POST --data '{
"original_url": "https://www.google.com",
"request_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"
}'   
http://localhost:5000/shorten-url
``

* Request id would ideally be passed in from the frontend application.
* Host can be changed to whatever the host/port is set to within `handler/handlers.go` and `main.go`

### Example Response:
``{"message":"successfully shortened URL","short_url":"http://localhost:5000/MjM5MDU1ND"}``

