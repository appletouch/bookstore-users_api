# bookstore-users_api
This is a go application that investigates making enterprise level Go software.


##Design:
Api is build with MVC pattern


###Endpoints
- GET     <host>:<port>/heartbeat
- POST    <host>:<port>/users
- GET     <host>:<port>/users/<user_id>

####Controllers:
The main goal of the controller is the entry point.
It handels the request and does the validation.
Only in this layer we use the http framework / server
- heartbeat =cloud required controller to check if api is alive


####Services
Contains the business logic

####Domain


####Error response
RFC 7807 provides a standard format for returning problem details from HTTP APIs. In particular, it specifies the following:

Error responses MUST use standard HTTP status codes in the 400 or 500 range to detail the general category of error.
Error responses will be of the Content-Type application/problem, appending a serialization format of either json or xml: 
application/problem+json, application/problem+xml.

Error responses will have each of the following keys:
- **detail (string)** - A human-readable description of the specific error.
- **type (string)** - a URL to a document describing the error condition (optional, and "about:blank" is assumed if none is provided; should resolve to a human-readable document).
- **title (string)** - A short, human-readable title for the general error type; the title should not change for given types.
- **status (number)** - Conveying the HTTP status code; this is so that all information is in one place, but also to correct for changes in the status code due to the usage of proxy servers. The status member, if present, is only advisory as generators MUST use the same status code in the actual HTTP response to assure that generic HTTP software that does not understand this format still behaves correctly.
- **instance (string)** - This optional key may be present, with a unique URI for the specific error; this will often point to an error log for that specific response.


##Frameworks used:
###web framework
- gin-ginonic (https://github.com/gin-gonic/gin) 40x faster the httprouter
Serving rest api in production
 a) go get -u github.com/gin-gonic/gin
 b) import "github.com/gin-gonic/gin"
 