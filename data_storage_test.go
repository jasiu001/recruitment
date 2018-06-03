package main

import "testing"

func TestNewAccount(t *testing.T) {
	account := NewAccount("testName")

	if account.UserName() != "testName" {
		t.Errorf("User name should be equal to 'testName' but '%s' is", account.UserName())
	}
}
