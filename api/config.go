package api

import (
	"errors"
	"os"
	"strings"
)

func GetConfig() (token string, repos []string, err error) {
	t := os.Getenv("GITHUB_TOKEN")
	if t == "" {
		return "", nil, errors.New("GITHUB_TOKEN not found")
	}

	r := os.Getenv("GITHUB_REPOS")
	if r == "" {
		return "", nil, errors.New("GITHUB_REPOS not found")
	}
	rA := strings.Split(r, ",")

	return t, rA, nil
}
