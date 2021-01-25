package gohttp

import "net/http"

type Response struct {
	status string
	statusCode int
	headers http.Header
	body []byte
}
