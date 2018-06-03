package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"html/template"
	"net/http"
	"path/filepath"
)

func init() {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleMain).Methods("GET")
	muxRouter.HandleFunc("/{userName}", handleUserInfoGet).Methods("GET")
	muxRouter.HandleFunc("/info", handleUserInfoPost).Methods("POST")

	http.Handle("/", muxRouter)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "home.html")))

	tmpl.Execute(w, nil)
}

func handleUserInfoGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctx := appengine.NewContext(r)

	userInfo(w, ctx, vars["userName"])
}

func handleUserInfoPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ctx := appengine.NewContext(r)

	userInfo(w, ctx, r.Form.Get("userName"))
}

func userInfo(w http.ResponseWriter, ctx context.Context, userName string) {
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
