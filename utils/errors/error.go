package errors

import (
	"fmt"
	"net/http"
)

//{
//"errors": {
//		"code": "NotFoundError",
//		"message": "Either there is no API method associated with the URL path of the request, or the request refers to one or more resources that were not found.",
//		"severity": "ERROR"
//		}
//}

type RestErr struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// standard error is made based on the error code or a custom error is made based on the message provided.
func New(code int, message ...string) *RestErr {
	var restErr *RestErr

	if len(message) == 0 {
		fmt.Printf("Error code:%d was thrown by applicaton...", code)
		switch code {
		case 400:
			restErr = &RestErr{
				Status: http.StatusBadRequest,
				Title:  "Bad Request",
				Detail: "Not a valid request or not correctly formatted",
			}
		case 401:
			fmt.Println("Unauthorized")
			restErr = &RestErr{
				Status: http.StatusUnauthorized,
				Title:  "Invalid token",
				Detail: "No valid token found",
			}
		case 404:
			restErr = &RestErr{
				Status: http.StatusNotFound,
				Title:  "Invalid token",
				Detail: "No valid token found",
			}
		case 500:
			restErr = &RestErr{
				Status: http.StatusInternalServerError,
				Title:  "Invalid token",
				Detail: "No valid token found",
			}
		default:
			restErr = &RestErr{
				Status: http.StatusNotAcceptable,
				Title:  "Invalid token",
				Detail: "No valid token found",
			}
		}
	} else {
		fmt.Println("Custom Bad request error")
		restErr = &RestErr{
			Status: code,
			Title:  "User error",
			Detail: message[0],
		}
	}

	return restErr
}
