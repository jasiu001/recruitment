package main

import (
	"context"
	"fmt"
	"os"
)

const userName = "jasiu001"

func main() {
	ctx := context.Background()
	client := NewClient(NewToken(), ctx)

	account := NewAccount(userName)
	err := FillAccount(ctx, client, account)

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	fmt.Println(account)
}
