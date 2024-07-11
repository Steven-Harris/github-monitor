package api

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (gh *ghHttpClient) GetPullRequests() (interface{}, error) {

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

func (gh *ghHttpClient) GetActions() (interface{}, error) {

	repos, err := GetRepos()
	if err != nil {
		return nil, err
	}

	runs := make([]Runs, 0)
	for _, repo := range repos {

		body, err := gh.Runs(repo)
		if err != nil {
			return nil, fmt.Errorf("error making request to get runs: %s", err)
		}

		var repoRun Runs
		err = json.Unmarshal(body, &repoRun)
		if err != nil {
			return nil, fmt.Errorf("error mapping response into object: %s", err)
		}
		repoRun.RepositoryName = strings.Split(repo, "/")[1]
		runs = append(runs, repoRun)
	}

	return runs, nil
}

func (gh *ghHttpClient) GetJobs(repo string, runId string) (interface{}, error) {

	body, err := gh.Jobs(repo, runId)
	if err != nil {
		return Jobs{}, fmt.Errorf("error making request to get runs: %s", err)
	}

	var jobs Jobs
	err = json.Unmarshal(body, &jobs)
	if err != nil {
		return Jobs{}, fmt.Errorf("error mapping response into object: %s", err)
	}

	return jobs, nil
}
