package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
)

// Fill Account Reposnse by data from Github api
func FillAccount(ctx context.Context, client *github.Client, account *Account) error {
	userResponse, _, err := client.Users.Get(ctx, account.UserName())
	if err != nil {
		return fmt.Errorf("Email error: %s", err)
	}

	account.SetEmail(userResponse.GetEmail())

	if userResponse.GetPublicRepos() == 0 {
		return nil
	}

	repositoryResponse, _, err := client.Repositories.List(ctx, account.UserName(), nil)
	if err != nil {
		return fmt.Errorf("Repositories error: %s", err)
	}

	for _, repository := range repositoryResponse {
		languageResponse, _, err := client.Repositories.ListLanguages(ctx, account.UserName(), repository.GetName())
		if err != nil {
			return fmt.Errorf("Language error: %s", err)
		}
		account.AddRepository(repository, languageResponse)
	}

	return nil
}
