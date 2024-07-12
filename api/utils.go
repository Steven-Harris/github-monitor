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

func GetActionRepos() ([]string, error) {
	r := os.Getenv("ACTION_REPOS")
	if r == "" {
		return nil, errors.New("ACTION_REPOS not found")
	}
	r = strings.ReplaceAll(r, "\n", "")
	return strings.Split(r, ","), nil
}

func GetActionFilter() string {
	f := os.Getenv("ACTION_FILTER")
	return f
}

func GetPRRepos() ([]PullRequestConfig, error) {
	r := os.Getenv("PR_REPOS")
	if r == "" {
		return nil, errors.New("PR_REPOS not found")
	}
	r = strings.ReplaceAll(r, "\n", "")
	repos := strings.Split(r, ",")
	var prConfigs []PullRequestConfig
	for _, repo := range repos {
		parts := strings.Split(repo, "?label=")
		label := ""
		if len(parts) > 1 {
			label = parts[1]
		}
		prConfig := PullRequestConfig{
			Repo:  parts[0],
			Label: label,
		}
		prConfigs = append(prConfigs, prConfig)
	}
	return prConfigs, nil
}
