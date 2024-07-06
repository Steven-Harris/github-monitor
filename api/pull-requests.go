package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const pullRequestsApi = "https://api.github.com/repos/%s/pulls?state=open"
const githubApiVersion = "2022-11-28"

func GetPullRequests() ([]PullRequest, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, errors.New("GITHUB_TOKEN not found")
	}

	repos := os.Getenv("GITHUB_REPOS")
	if repos == "" {
		return nil, errors.New("GITHUB_REPOS not found")
	}

	repoArr := strings.Split(repos, ",")

	pullRequests := make([]PullRequest, 0)
	for _, repo := range repoArr {
		url := fmt.Sprintf(pullRequestsApi, repo)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, errors.New("could not make request to Github")
		}

		req.Header.Add("Accept", "application/vnd.github+json")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Add("X-Github-Api-Version", githubApiVersion)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, errors.New("error making request to get pull requests")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, errors.New("could not read response body from Github Api")
		}

		var repoPulls []PullRequest
		err = json.Unmarshal(body, &repoPulls)
		if err != nil {
			return nil, errors.New("error mapping response into object")
		}

		pullRequests = append(pullRequests, repoPulls...)
	}

	return pullRequests, nil
}
