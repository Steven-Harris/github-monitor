package api

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetToken() (string, error) {
	t := os.Getenv("GITHUB_TOKEN")
	if t == "" {
		return "", errors.New("GITHUB_TOKEN not found")
	}
	fmt.Println("Token: ", t)
	return t, nil
}

func GetActionRepos() ([]string, error) {
	r := os.Getenv("MONITOR_ACTION_REPOS")
	if r == "" {
		return nil, errors.New("MONITOR_ACTION_REPOS not found")
	}
	return strings.Split(r, ","), nil
}

func GetPRRepos() ([]string, error) {
	r := os.Getenv("MONITOR_PR_REPOS")
	if r == "" {
		return nil, errors.New("MONITOR_PR_REPOS not found")
	}
	return strings.Split(r, ","), nil
}
