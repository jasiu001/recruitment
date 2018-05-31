package main

import "github.com/google/go-github/github"

type Account struct {
	name         string
	email        string
	repositories []Repository
}

type Repository struct {
	name         string
	fullName     string
	url          string
	languageList []Language
}

type Language struct {
	name            string
	nuberOfByteCode int
	percentage      int
}

func NewAccount(userName string) *Account {
	return &Account{
		name: userName,
	}
}

func (a Account) UserName() string {
	return a.name
}

func (a *Account) SetEmail(value string) {
	a.email = value
}

func (a *Account) AddRepository(repo *github.Repository, languages map[string]int) {
	RepoLang := []Language{}
	for lang, byteAmount := range languages {
		RepoLang = append(RepoLang, Language{
			name:            lang,
			nuberOfByteCode: byteAmount,
		})
	}

	a.repositories = append(a.repositories, Repository{
		name:         repo.GetName(),
		fullName:     repo.GetFullName(),
		languageList: RepoLang,
	})
}
