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

func GetRepos() ([]string, error) {
	r := os.Getenv("GITHUB_REPOS")
	if r == "" {
		return nil, errors.New("GITHUB_REPOS not found")
	}
	return strings.Split(r, ","), nil
}
