package main

import "fmt"

type errorResponse struct {
	Error *Error `json:"error"`
}

type Error struct {
	ErrorResponse *Response
	Code          int         `json:"code,omitempty"`
	Errors        interface{} `json:"errors,omitempty"`
	Message       string      `json:"message,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s API call to %s failed %v. Code: %d, Errors: %v, Message: %s", e.ErrorResponse.Response.Method, e.ErrorResponse.Response.Request.URL.String(), e.ErrorResponse.Response.Status, e.Code, e.Errors, e.Message)
}
