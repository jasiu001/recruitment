package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const fileName = "token.json"

type TokenProvider struct {
	token string
}

type TokenFile struct {
	Token string `json:"token"`
}

// Return token struct with token necessary to connect with github API
func NewToken() *TokenProvider {
	return &TokenProvider{
		token: readFromFile(),
	}
}

func (t TokenProvider) Get() string {
	return t.token
}

func readFromFile() string {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var token TokenFile
	err = json.Unmarshal(raw, &token)
	if err != nil {
		fmt.Println("There is no token in file")
		os.Exit(1)
	}

	return token.Token
}
