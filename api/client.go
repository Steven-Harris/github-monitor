package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const githubApiVersion = "2022-11-28"
const ghUri = "https://api.github.com/repos/"

type ghHttpClient struct {
	baseUri *url.URL
	client  *http.Client
	token   string
}

type GHClient interface {
	Pulls(repo string) ([]byte, error)
}

func NewGitHubHttpClient() (*ghHttpClient, error) {
	token, err := GetToken()
	if err != nil {
		return nil, err
	}

	base, err := url.Parse(ghUri)
	if err != nil {
		return nil, err
	}

	return &ghHttpClient{
		baseUri: base,
		client: &http.Client{
			Timeout:   30 * time.Second,
			Transport: http.DefaultTransport,
		},
		token: token,
	}, nil
}

func (c *ghHttpClient) Pulls(repo string) ([]byte, error) {
	rel := &url.URL{Path: fmt.Sprintf("%s/pulls", repo)}
	query := url.Values{}
	query.Add("state", "open")
	rel.RawQuery = query.Encode()
	uri := c.baseUri.ResolveReference(rel)
	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, errors.New("could not build request")
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Add("X-Github-Api-Version", githubApiVersion)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, errors.New("could not connect to Github")
	}

	switch res.StatusCode {
	case http.StatusNotFound, http.StatusInternalServerError:
		{
			return nil, fmt.Errorf("url not found: %s", uri.String())
		}
	case http.StatusOK:
		{
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return nil, errors.New("could not read response body from Github Api")
			}

			return body, nil
		}
	default:
		return nil, fmt.Errorf("something happened: %s", res.Status)
	}
}
