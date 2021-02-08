package examples

import (
	"errors"
	"fmt"
	"github.com/RBrothersBSI/go-httpclient/gohttp_mock"
	"net/http"
	"testing"
)



func TestCreateRepo(t *testing.T){

	t.Run("timeoutFromGithub", func(t *testing.T) {
		gohttp_mock.MockupServer.DeleteMocks()
		gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
			Method: http.MethodPost,
			Url: "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo-timeout","description":"","private":true}`,
			Error: errors.New("timeout from github"),
		})

		repository := Repository{
			Name: "test-repo-timeout",
			Private: true,
		}

		repo, err := CreateRepo(repository)

		if repo != nil {
			t.Error("Expected no repo go", repo)
		}

		if err == nil {
			t.Error("Expected error got", err)
		}

		if err.Error() != "timeout from github" {
			t.Error("Expected 'timeout from github' got", err.Error())
		}
	})

	t.Run("noError", func(t *testing.T) {
		gohttp_mock.MockupServer.DeleteMocks()

		gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
			Method: http.MethodPost,
			Url: "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo-no-error","description":"","private":true}`,
			ResponseStatusCode: http.StatusCreated,
			ResponseBody: `{"id": 123, "name":"test-repo-no-error"}`,
		})

		repository := Repository{
			Name: "test-repo-no-error",
			Private: true,
		}

		repo, err := CreateRepo(repository)

		if err != nil {
			t.Error("Expected no error got", err)
		}

		if repo == nil {
			t.Error("Expected repo got", repo)
		}

		if repo.Name != repository.Name{
			t.Error(fmt.Sprintf("expected %s got %s", repository.Name, repo.Name))
		}
	})
}

