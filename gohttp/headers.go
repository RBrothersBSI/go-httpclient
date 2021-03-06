package gohttp

import (
	"github.com/RBrothersBSI/go-httpclient/gomime"
	"net/http"
)

func getHeaders(headers ...http.Header) http.Header{
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	//Add common headers to request
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	//Add custom headers to request
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	//Set User-Agent if it is not there
	if c.builder.userAgent != "" {
		 if result.Get(gomime.HeaderUserAgent) != "" {
		 	return result
		}
		result.Set("User-Agent", c.builder.userAgent)
	}
	return result
}