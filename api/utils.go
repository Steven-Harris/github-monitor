package api

import (
	"errors"
	"os"
	"strings"
)

func GetToken() (string, error) {
	t := os.Getenv("GITHUB_TOKEN")
	if t == "" {
		return "", errors.New("GITHUB_TOKEN not found")
	}
	return t, nil
}

func GetOrg() (string, error) {
	o := os.Getenv("GITHUB_ORG")
	if o == "" {
		return "", errors.New("GITHUB_ORG not found")
	}
	return o, nil
}

func GetActionRepos() ([]string, error) {
	r := os.Getenv("ACTION_REPOS")
	if r == "" {
		return nil, errors.New("ACTION_REPOS not found")
	}
	r = strings.ReplaceAll(r, "\n", "")
	repos := strings.Split(r, ",")
	org, err := GetOrg()
	if err != nil {
		return nil, err
	}
	for i, repo := range repos {
		repos[i] = org + "/" + repo
	}
	return repos, nil
}

func GetActionFilter() string {
	f := os.Getenv("ACTION_FILTER")
	return f
}

func GetPRRepos() ([]string, error) {
	r := os.Getenv("PR_REPOS")
	if r == "" {
		return nil, errors.New("PR_REPOS not found")
	}
	r = strings.ReplaceAll(r, "\n", "")
	repos := strings.Split(r, ",")
	org, err := GetOrg()
	if err != nil {
		return nil, err
	}
	for i, repo := range repos {
		repos[i] = org + "/" + repo
	}
	return repos, nil
}
