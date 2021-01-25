package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct{
	disableTimeouts bool
	maxIdleConnections int
	connectionTimeout time.Duration
	responseTimeout time.Duration
	headers http.Header
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(i int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	Build() Client
}

func(c *clientBuilder) Build() Client {
	client := &httpClient{
		builder: c,
	}
	return client
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder{
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder{
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder{
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(i int) ClientBuilder {
	c.maxIdleConnections = i
	return c
}

func (c *clientBuilder) DisableTimeouts(b bool) ClientBuilder{
	c.disableTimeouts = b
	return c
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{
	}
	return builder
}
