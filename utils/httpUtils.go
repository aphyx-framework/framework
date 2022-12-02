package utils

import (
	"github.com/go-playground/validator"
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

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func DecodeUrlParam(param string) string {
	decoded, err := url.QueryUnescape(param)

	if err != nil {
		ErrorLogger.Fatalln(err)
	}

	return decoded
}

func GetErrors(err error) []*ErrorResponse {
	var errors []*ErrorResponse

	for _, err := range err.(validator.ValidationErrors) {
		var element ErrorResponse
		element.FailedField = err.StructNamespace()
		element.Tag = err.Tag()
		element.Value = err.Value().(string)
		errors = append(errors, &element)
	}

	return errors
}
