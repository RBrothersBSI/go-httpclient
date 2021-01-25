package main

import (
	"fmt"
	"github.com/bsi/go-httpclient/gohttp"
	"net/http"
)

var(
	httpClient = getGithubClient()
)

func getGithubClient() gohttp.Client{

	client := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(5).
		Build()

	//Common headers
//	commonHeaders := make(http.Header)
//	commonHeaders.Set("Authorization", "Bearer ABC-123")

	return client
}

func main(){
	go func() {
		getUrls()
	}()
}

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func getUrls() {
	//Custom Headers
	headers := make(http.Header)
	response, err := httpClient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Status)
}

func createUser(user User) {
	//Custom Headers
	headers := make(http.Header)
	response, err := httpClient.Post("https://api.github.com", headers, user)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Status)
}