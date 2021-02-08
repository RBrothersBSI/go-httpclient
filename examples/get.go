package examples

import "fmt"

type Endpoints struct {
	CurrentUserUrl string `json:"current_user_url"`
	AuthUrl string `json:"authorizations_url"`
	RepoUrl string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	res, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		//Deal with error
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Status Code: %d", res.StatusCode))
	fmt.Println(fmt.Sprintf("Status: %s", res.Status))
	fmt.Println(fmt.Sprintf("Body: %s", res.String()))

	var endpoints Endpoints
	if err := res.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Repo Url: %s", endpoints.RepoUrl))
	return &endpoints, nil
}
