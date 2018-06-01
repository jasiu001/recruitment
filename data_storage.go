package main

import (
	"errors"
	"fmt"
	"github.com/google/go-github/github"
)

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

func (a Account) GetEmail() string {
	return a.email
}

func (a *Account) SetEmail(value string) {
	a.email = value
}

func (r Repository) GetName() string {
	return r.name
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

func (a Account) GetRepositories() []string {
	var repositories []string

	for _, repo := range a.repositories {
		repositories = append(repositories, repo.GetName())
	}

	return repositories
}

func (a Account) GetRepositoryLanguage(name string) (map[string]int, error) {
	repo, err := a.findRepositoryByName(name)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	data := make(map[string]int)
	for _, lang := range repo.languageList {
		data[lang.name] = lang.percentage
	}

	return data, nil
}

func (a Account) findRepositoryByName(repoName string) (Repository, error) {
	for _, repo := range a.repositories {
		if repo.GetName() == repoName {
			return repo, nil
		}
	}

	return Repository{}, errors.New(fmt.Sprintf("There is no repository with name: %s", repoName))
}
