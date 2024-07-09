package api

import (
	"encoding/json"
	"fmt"
)

func (gh *ghHttpClient) GetPullRequests() ([]PullRequest, error) {

	repos, err := GetRepos()
	if err != nil {
		return nil, err
	}

	pullRequests := make([]PullRequest, 0)
	for _, repo := range repos {

		body, err := gh.Pulls(repo)
		if err != nil {
			return nil, fmt.Errorf("error making request to get pull requests: %s", err)
		}

		var repoPulls []PullRequest
		err = json.Unmarshal(body, &repoPulls)
		if err != nil {
			return nil, fmt.Errorf("error mapping response into object: %s", err)
		}

		pullRequests = append(pullRequests, repoPulls...)
	}

	return pullRequests, nil
}
