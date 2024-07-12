package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

func (gh *ghHttpClient) GetPullRequests() (interface{}, error) {

	repos, err := GetPRRepos()
	if err != nil {
		return nil, err
	}
	query := url.Values{}
	query.Add("state", "open")
	repoPulls := make([]RepoPullRequests, len(repos))
	for i := 0; i < len(repos); i++ {

		label := ""
		if repos[i].Label != "" {
			label = fmt.Sprintf("+label:%s", repos[i].Label)
		}
		query := fmt.Sprintf("repo:%s+is:pr%s+is:open", repos[i].Repo, label)
		body, err := gh.search(query)
		if err != nil {
			return nil, fmt.Errorf("error making request to get pull requests: %s", err)
		}

		var results SearchResults
		err = json.Unmarshal(body, &results)
		if err != nil {
			return nil, fmt.Errorf("error mapping response into object: %s", err)
		}

		repoPulls[i] = RepoPullRequests{
			RepositoryName: strings.Split(repos[i].Repo, "/")[1],
			PullRequests:   results.Items,
		}
	}

	return repoPulls, nil
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

func (gh *ghHttpClient) GetActions() (interface{}, error) {

	repos, err := GetActionRepos()
	if err != nil {
		return nil, err
	}
	filter := GetActionFilter()
	if filter != "" {
		filter = fmt.Sprintf("/workflows/%s.yml", filter)
	}
	query := url.Values{}
	query.Add("per_page", "1")

	runs := make([]Runs, 0)
	for _, repo := range repos {
		body, err := gh.request(fmt.Sprintf("%s/actions%s/runs", repo, filter), query)
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
