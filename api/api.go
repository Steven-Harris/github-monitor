package api

import (
	"encoding/json"
	"errors"
	"fmt"
)

const pullRequestsApi = "https://api.github.com/repos/%s/pulls?state=open"

func GetPullRequests() ([]PullRequest, error) {

	token, repos, err := GetConfig()
	if err != nil {
		return nil, err
	}

	pullRequests := make([]PullRequest, 0)
	for _, repo := range repos {
		url := fmt.Sprintf(pullRequestsApi, repo)

		body, err := GithubGet(url, token)
		if err != nil {
			return nil, fmt.Errorf("error making request to get pull requests: %s\n", err)
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
