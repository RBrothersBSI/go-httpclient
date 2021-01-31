package gohttp

import (
	"testing"
)

//func TestGetRequestHeaders(t *testing.T){
//	//Initialization
//	client := httpClient{}
//	commonHeaders := make(http.Header)
//	commonHeaders.Set("Content-Type", "application/json")
//	commonHeaders.Set("User-Agent", "testing")
//	client.Headers = commonHeaders
//
//	// Execution
//	requestHeaders := make(http.Header)
//	requestHeaders.Set("X-Request-Id", "abc123")
//	finalHeaders := client.getRequestHeaders(requestHeaders)
//
//	// Validation
//	if len(finalHeaders) != 3 {
//		t.Error("Expected 3 headers, got ", len(finalHeaders))
//	}
//}

func TestGetRequestBodyNilBody(t *testing.T) {
	//Initialization
	client := httpClient{}

	t.Run("noBodyNilResponse", func(t *testing.T){
		//Execution
		body, err := client.getRequestBody("", nil)
		//Validation
		if err != nil {
			t.Error("no error expected when passing a nil body")
		}

		if body != nil {
			t.Error("Expected nil body, got ", body)
		}
	})

	t.Run("bodyWithJson", func(t *testing.T){
		//Initialization
		requestBody := []string{"a","b"}
		//Execution
		body, err := client.getRequestBody("application/json",  requestBody)

		//Validation
		if err != nil {
			t.Error("Expected json, got ", err)
		}

		if string(body) != `["a","b"]` {
			t.Error("Exepcted [\"a\",\"b\"] got ", string(body))
		}
	})
	t.Run("bodyWithXML", func(t *testing.T){})
	t.Run("defaultbody", func(t *testing.T){})

}

func TestGetRequestBodyWithJson(t *testing.T) {

}

func TestGetRequestBodyWithXml(t *testing.T) {

}

func TestGetRequestBodyDefault(t *testing.T) {

}