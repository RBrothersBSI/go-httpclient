package gohttp

import "net/http"

type HttpClient interface {
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
	SetHeaders(headers http.Header)
}

type httpClient struct{
	Headers http.Header
}

func New() HttpClient {
	client:= &httpClient{}
	return client
}

func (c *httpClient) SetHeaders(headers http.Header){
	c.Headers = headers
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error){
	return c.do(http.MethodGet,url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodGet,url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error){
	return c.do(http.MethodGet,url, headers,  body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error){
	return c.do(http.MethodGet,url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet,url, headers, nil)
}






