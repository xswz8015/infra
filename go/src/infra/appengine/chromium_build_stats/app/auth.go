package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine/v2/log"
	"google.golang.org/appengine/v2/user"
)

func authPage(w http.ResponseWriter, req *http.Request, status int, tmpl *template.Template, u *user.User, reqpath string) {
	ctx := req.Context()
	login, err := user.LoginURL(ctx, reqpath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logout, err := user.LogoutURL(ctx, reqpath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(status)
	data := struct {
		User   *user.User
		Login  string
		Logout string
	}{
		User:   u,
		Login:  login,
		Logout: logout,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Errorf(ctx, "tmpl: %v", err)
	}
}
