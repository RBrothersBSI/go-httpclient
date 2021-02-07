package examples

import (
	"github.com/RBrothersBSI/go-httpclient/gohttp"
	"github.com/RBrothersBSI/go-httpclient/gomime"
	"net/http"
	"time"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {

	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.HeaderUserAgent)

	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("Ryans-pc").
		Build()

	return client
}
