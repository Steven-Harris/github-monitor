package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

const githubApiVersion = "2022-11-28"

type GithubRequest struct {
	req *http.Request
}

func (req GithubRequest) setDefaultHeaders(token string) {
	req.req.Header.Add("Accept", "application/vnd.github+json")
	req.req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.req.Header.Add("X-Github-Api-Version", githubApiVersion)
}

func GithubGet(token string, url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.New("could not build request")
	}

	ghc := &GithubRequest{req: req}

	ghc.setDefaultHeaders(token)

	res, err := http.DefaultClient.Do(ghc.req)
	if err != nil {
		return nil, errors.New("could not connect to Github")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("could not read response body from Github Api")
	}

	return body, nil
}
