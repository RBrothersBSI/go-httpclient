package examples

import (
	"errors"
	"fmt"
	"github.com/RBrothersBSI/go-httpclient/gohttp"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start test suite for package examples")
	//Tell the http library to mock any further requests from here
	gohttp.StartMockServer()
	os.Exit(m.Run())
}

func TestGet(t *testing.T){
	//Init
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url: "https://api.github.com",
			Error: errors.New("Timeout getting response"),
		})

		//Execution:
		endpoints, err := GetEndpoints()

		//Validation
		if endpoints != nil {
			t.Error("No endpoints expected at this point in test")
		}

		if err == nil {
			t.Error("Expected Error, got ", err)
		}

		if err.Error() != "Timeout getting response" {
			t.Error("Expected 'Timeout Getting Response' got ", err.Error())
		}

	})
	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		gohttp.FlushMocks()
		//Init
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url: "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody: `{"current_user_url": 123}`,
		})

		//Execution:
		endpoints, err := GetEndpoints()

		//Validation
		if endpoints != nil {
			t.Error("No endpoints expected at this point")
		}

		if err == nil {
			t.Error("Expected Error, got ", err)
		}

		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct") {
			t.Error("Expected 'Timeout Getting Response' got ", err.Error())
		}

	})
	t.Run("TestNoError", func(t *testing.T) {
		//Init
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url: "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody: `{"current_user_url": "https://api.github.com/user"}`,
		})

		//Execution:
		endpoints, err := GetEndpoints()

		//Validation

		if err != nil {
			t.Error(fmt.Sprintf("no error was expected and we got %s", err.Error()))
		}

		if endpoints == nil {
			t.Error("endpoints were expected at this point")
		}

		if endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("invalid current user url")
		}
	})
}