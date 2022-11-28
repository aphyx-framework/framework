package utils

import (
	"net/url"
)

// HttpResponse ---
//
// This is the recommended way to return a response from the framework
// It is a struct that contains the status code, message and data
// For failed responses, the data field can be empty
type HttpResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func DecodeUrlParam(param string) string {
	decoded, err := url.QueryUnescape(param)

	if err != nil {
		ErrorLogger.Fatalln(err)
	}

	return decoded
}
