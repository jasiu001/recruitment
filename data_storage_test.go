package main

import (
	"reflect"
	"testing"
)

func TestNewAccount(t *testing.T) {
	account := NewAccount("testName")

	if account.UserName() != "testName" {
		t.Errorf("User name should be equal to 'testName' but '%s' is", account.UserName())
	}
}

type providerTest struct {
	name     string
	fullName string
}

func (p providerTest) GetName() string {
	return p.name
}

func (p providerTest) GetFullName() string {
	return p.fullName
}

func TestAccount_AddRepository(t *testing.T) {
	provider := providerTest{name: "test", fullName: "test/testRepo"}
	account := NewAccount("testName")

	account.AddRepository(provider, map[string]int{"langA": 150, "langB": 50})

	repositories := account.GetRepositories()
	for _, repo := range repositories {
		if repo != "test" {
			t.Errorf("Account should include repository with name 'test' but it does not")
		}
	}

	languages, _ := account.GetRepositoryLanguage("test")
	if !reflect.DeepEqual(languages, map[string]string{"langA": "75.00", "langB": "25.00"}) {
		t.Errorf("Languages repository statistic are incorrect")
	}
}
