package handler

import (
	"food-website/model"
	"food-website/service"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))
var user model.User

func LoginPageHandler(w http.ResponseWriter, r *http.Request){
	err := service.CheckTokens(w, r)
	switch{
	case err.Error() == "access token expired":
		errUpdate := service.UpdateAccessToken(w, r, user)
		if errUpdate != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errUpdate.Error()))
			return
		}
		break
	case err == nil:
		http.Redirect(w, r, "/", 302)
		return
	case err.Error() == "all tokens expired":
		break
	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = templates.ExecuteTemplate(w, "login.html", nil)
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
	user.Email = r.PostForm.Get("email")
	user.HashedPassword = r.PostForm.Get("password")
	err = service.CheckLogin(&user)
	if  err != nil{
		templates.ExecuteTemplate(w, "login.html", err.Error())
		return
	}
	err = service.UpdateAccessToken(w, r, user)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = service.UpdateRefreshToken(w, r, user)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	http.Redirect(w, r, "/", 302)
	return
}
