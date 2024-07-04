package main

import (
	"context"
	"crypto/rsa"
	"errors"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/go-github/v62/github"
	"golang.org/x/oauth2"
)

const (
	DefaultApplicationTokenExpiration = 10 * time.Minute
	bearerTokenType                   = "Bearer"
)

type applicationTokenSource struct {
	id         int64
	privateKey *rsa.PrivateKey
	expiration time.Duration
}

type ApplicationTokenOpt func(*applicationTokenSource)

func WithApplicationTokenExpiration(expiration time.Duration) ApplicationTokenOpt {
	return func(a *applicationTokenSource) {
		if expiration > DefaultApplicationTokenExpiration || expiration <= 0 {
			expiration = DefaultApplicationTokenExpiration
		}
		a.expiration = expiration
	}
}

func NewApplicationTokenSource(id int64, privateKey []byte, opts ...ApplicationTokenOpt) (oauth2.TokenSource, error) {
	if id == 0 {
		return nil, errors.New("application id is required")
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, err
	}

	t := &applicationTokenSource{
		id:         id,
		privateKey: privKey,
		expiration: DefaultApplicationTokenExpiration,
	}

	for _, opt := range opts {
		opt(t)
	}

	return t, nil
}

func (t *applicationTokenSource) Token() (*oauth2.Token, error) {
	// To protect against clock drift, set the issuance time 60 seconds in the past.
	now := time.Now().Add(-60 * time.Second)
	expiresAt := now.Add(t.expiration)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		Issuer:    strconv.FormatInt(t.id, 10),
	})

	tokenString, err := token.SignedString(t.privateKey)
	if err != nil {
		return nil, err
	}

	return &oauth2.Token{
		AccessToken: tokenString,
		TokenType:   bearerTokenType,
		Expiry:      expiresAt,
	}, nil
}

type InstallationTokenSourceOpt func(*installationTokenSource)

func WithInstallationTokenOptions(opts *github.InstallationTokenOptions) InstallationTokenSourceOpt {
	return func(i *installationTokenSource) {
		i.opts = opts
	}
}

func WithHTTPClient(client *http.Client) InstallationTokenSourceOpt {
	return func(i *installationTokenSource) {
		client.Transport = &oauth2.Transport{
			Source: i.src,
			Base:   client.Transport,
		}

		i.apps = github.NewClient(client).Apps
	}
}

type installationTokenSource struct {
	id   int64
	src  oauth2.TokenSource
	apps *github.AppsService
	opts *github.InstallationTokenOptions
}

func NewInstallationTokenSource(id int64, src oauth2.TokenSource, opts ...InstallationTokenSourceOpt) oauth2.TokenSource {
	client := &http.Client{
		Transport: &oauth2.Transport{
			Source: src,
		},
	}

	i := &installationTokenSource{
		id:   id,
		src:  src,
		apps: github.NewClient(client).Apps,
	}

	for _, opt := range opts {
		opt(i)
	}

	return i
}

// Token generates a new GitHub App installation token for authenticating as a GitHub App installation.
func (t *installationTokenSource) Token() (*oauth2.Token, error) {
	ctx := context.Background()

	token, _, err := t.apps.CreateInstallationToken(ctx, t.id, t.opts)
	if err != nil {
		return nil, err
	}

	return &oauth2.Token{
		AccessToken: token.GetToken(),
		TokenType:   bearerTokenType,
		Expiry:      token.GetExpiresAt().Time,
	}, nil
}
