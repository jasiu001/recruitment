package main

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func NewClient(token *TokenProvider, ctx context.Context) *github.Client {

	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.Get()},
	)
	tokenClient := oauth2.NewClient(ctx, tokenService)

	return github.NewClient(tokenClient)
}
