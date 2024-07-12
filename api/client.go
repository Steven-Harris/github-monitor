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
const ghUri = "https://api.github.com/"

type ghHttpClient struct {
	baseUri *url.URL
	client  *http.Client
	token   string
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

func (c *ghHttpClient) search(query string) ([]byte, error) {
	rel := &url.URL{Path: "search/issues"}
	uri := c.baseUri.ResolveReference(rel)
	// Manually append the query string without encoding
	queryString := "?q=" + query // This is where you directly use the query without encoding
	fullURL := uri.String() + queryString
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, errors.New("could not build request")
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Add("X-Github-Api-Version", githubApiVersion)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not make request: %s", err)
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusNotFound, http.StatusInternalServerError:
		return nil, fmt.Errorf("url not found: %s", uri.String())
	case http.StatusOK:
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, errors.New("could not read response body from Github Api")
		}
		return body, nil
	default:
		return nil, fmt.Errorf("something happened: %s", res.Status)
	}
}

func (c *ghHttpClient) request(path string, query url.Values) ([]byte, error) {
	rel := &url.URL{Path: "repos/" + path}
	if query != nil {
		rel.RawQuery = query.Encode()
	}
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
		return nil, fmt.Errorf("could not make request: %s", err)
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusNotFound, http.StatusInternalServerError:
		return nil, fmt.Errorf("url not found: %s", uri.String())
	case http.StatusOK:
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, errors.New("could not read response body from Github Api")
		}
		return body, nil
	default:
		return nil, fmt.Errorf("something happened: %s", res.Status)
	}
}
