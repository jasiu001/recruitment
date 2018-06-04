package main

import (
	"errors"
	"fmt"
)

type RepositoryProvider interface {
	GetName() string
	GetFullName() string
}

type Account struct {
	name         string
	email        string
	repositories []*Repository
}

type Repository struct {
	name         string
	fullName     string
	url          string
	languageList []*Language
}

type Language struct {
	name            string
	nuberOfByteCode int
	percentage      string
}

// Create new account user struct from the passed name
func NewAccount(userName string) *Account {
	return &Account{
		name: userName,
	}
}

// Return account user name
func (a Account) UserName() string {
	return a.name
}

// Return account email, it could be also empty string
func (a Account) GetEmail() string {
	return a.email
}

func (a *Account) SetEmail(value string) {
	a.email = value
}

func (r Repository) GetName() string {
	return r.name
}

// Add user repository with language statistic
func (a *Account) AddRepository(repo RepositoryProvider, languages map[string]int) {
	RepoLang := []*Language{}
	for lang, byteAmount := range languages {
		RepoLang = append(RepoLang, &Language{
			name:            lang,
			nuberOfByteCode: byteAmount,
		})
	}

	a.repositories = append(a.repositories, &Repository{
		name:         repo.GetName(),
		fullName:     repo.GetFullName(),
		languageList: RepoLang,
	})
}

// Return all repositories names
func (a Account) GetRepositories() []string {
	var repositories []string

	for _, repo := range a.repositories {
		repositories = append(repositories, repo.GetName())
	}

	return repositories
}

// Return languages statistic of repository based on repository name
func (a Account) GetRepositoryLanguage(name string) (map[string]string, error) {
	repo, err := a.findRepositoryByName(name)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	data := make(map[string]string)
	repo.countLanguagesPercentage()
	for _, lang := range repo.languageList {
		data[lang.name] = lang.percentage
	}

	return data, nil
}

func (a Account) findRepositoryByName(repoName string) (*Repository, error) {
	for _, repo := range a.repositories {
		if repo.GetName() == repoName {
			return repo, nil
		}
	}

	return &Repository{}, errors.New(fmt.Sprintf("There is no repository with name: %s", repoName))
}

func (r *Repository) countLanguagesPercentage() {
	ld := NewLanguageData()
	for _, lang := range r.languageList {
		ld.AddLanguageItem(lang.name, lang.nuberOfByteCode)
	}

	ld.CountPercentage()

	for _, lang := range r.languageList {
		lang.percentage = ld.GetPercentege(lang.name)
	}
}
