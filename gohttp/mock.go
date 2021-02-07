package gohttp

import (
	"fmt"
	"net/http"
)

// Mock struct provides a clean way to configure http mocks based on the combination
// of request method, URL, and request body
type Mock struct {
	Method string
	Url string
	RequestBody string

	ResponseBody string
	Error error
	ResponseStatusCode int
}

// GetResponse returns a Response object based on the mock configuration
func (m *Mock) GetResponse() (*Response, error){
	if m.Error != nil {
		return nil, m.Error
	}

	response := Response{
		status: fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		statusCode: m.ResponseStatusCode,
		body: []byte(m.ResponseBody),
	}

	return &response, nil
}


