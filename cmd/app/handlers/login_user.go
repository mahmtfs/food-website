package handlers

import (
	"html/template"
	"net/http"
	"working-project/internal/app/method/user"
	"working-project/internal/app/model"
)

func LoginUser(w http.ResponseWriter, r *http.Request){
	u := model.User{}
	r.ParseForm()
	u.Email = r.Form["email"][0]
	u.EncryptedPassword = r.Form["password"][0]
	err := user.UserDBRepository{}.CheckLogin(u)
	user.HandleError(err)
	var tmplt = template.Must(template.ParseFiles("pkg/templates/index.html"))
	err = tmplt.Execute(w, nil)
	user.HandleError(err)
}
