package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleMain).Methods("GET")
	muxRouter.HandleFunc("/{userName}", handleUserInfoGet).Methods("GET")
	muxRouter.HandleFunc("/info", handleUserInfoPost).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", muxRouter))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "home.html")))

	tmpl.Execute(w, nil)
}

func handleUserInfoGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userInfo(w, vars["userName"])
}

func handleUserInfoPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userInfo(w, r.Form.Get("userName"))
}

func userInfo(w http.ResponseWriter, userName string) {
	ctx := context.Background()
	client := NewClient(NewToken(), ctx)

	account := NewAccount(userName)
	err := FillAccount(ctx, client, account)

	if err != nil {
		tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "error.html")))
		tmpl.Execute(w, NewErrorPage(fmt.Sprintf("Problem in getting repository information %v\n", err)))
		return
	}

	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "info.html")))
	tmpl.Execute(w, NewInfoPage(account))
}
