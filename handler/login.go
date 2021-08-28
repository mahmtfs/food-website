package handler

import (
	"food-website/service"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))
var store = sessions.NewCookieStore([]byte("tsc"))

func LoginPageHandler(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session")
	if session.Values["email"] != nil{
		http.Redirect(w, r, "/", 302)
		return
	}
	err := templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to execute a template"))
		return
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to parse form"))
		return
	}
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	err = service.CheckLogin(email, password)
	if  err != nil{
		templates.ExecuteTemplate(w, "login.html", err.Error())
		return
	}
	session, _ := store.Get(r, "session")
	session.Values["email"] = email
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}
