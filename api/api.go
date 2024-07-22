package api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (gh ghHttpClient) GetPRRepos() (interface{}, error) {

	repos, err := GetPRRepos()
	if err != nil {
		return nil, err
	}

	return repos, nil
}

func (gh ghHttpClient) GetActionRepos() (interface{}, error) {

	repos, err := GetActionRepos()
	if err != nil {
		return nil, err
	}

	return repos, nil
}

func (gh *ghHttpClient) GetPullRequests(repo string) (interface{}, error) {

	query := url.Values{}
	query.Add("state", "open")
	org, err := GetOrg()
	if err != nil {
		return nil, fmt.Errorf("error getting org: %s", err)
	}

	filter := GetPRFilter()

	url := fmt.Sprintf("repo:%s/%s+is:pr+is:open+%s", org, repo, filter)
	body, err := gh.search(url)
	if err != nil {
		return nil, fmt.Errorf("error making request to get pull requests: %s", err)
	}

	var results SearchResults
	err = json.Unmarshal(body, &results)
	if err != nil {
		return nil, fmt.Errorf("error mapping response into object: %s", err)
	}

	return results, nil
}

func (gh *ghHttpClient) GetReviews(repo string, prNumber string) (interface{}, error) {

	body, err := gh.request(fmt.Sprintf("%s/pulls/%s/reviews", repo, prNumber), nil)
	if err != nil {
		return nil, fmt.Errorf("error making request to get reviews: %s", err)
	}

	var reviews []Review
	err = json.Unmarshal(body, &reviews)
	if err != nil {
		return nil, fmt.Errorf("error mapping response into object: %s", err)
	}

	return reviews, nil
}

func (gh *ghHttpClient) GetActions(repo string) (interface{}, error) {

	filter := GetActionFilter()
	if filter != "" {
		filter = fmt.Sprintf("/workflows/%s.yml", filter)
	}
	query := url.Values{}
	query.Add("per_page", "1")

	body, err := gh.request(fmt.Sprintf("%s/actions%s/runs", repo, filter), query)
	if err != nil {
		return nil, fmt.Errorf("error making request to get runs: %s", err)
	}

	var repoRun Runs
	err = json.Unmarshal(body, &repoRun)
	if err != nil {
		return nil, fmt.Errorf("error mapping response into object: %s", err)
	}
	return repoRun, nil
}

func (gh *ghHttpClient) GetJobs(repo string, runId string) (interface{}, error) {

	body, err := gh.request(fmt.Sprintf("%s/actions/runs/%s/jobs", repo, runId), nil)
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
