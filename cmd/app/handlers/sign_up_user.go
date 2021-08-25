package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"working-project/internal/app/method/user"
	"working-project/internal/app/model"
)

func SignUpUser(w http.ResponseWriter, r *http.Request){
	u := model.User{}
	err := r.ParseForm()
	user.HandleError(err)
	u.FirstName = r.Form["first_name"][0]
	u.LastName = r.Form["last_name"][0]
	u.Email = r.Form["email"][0]
	u.EncryptedPassword = r.Form["password"][0]
	fmt.Println(u.FirstName, u.LastName, u.Email, u.EncryptedPassword)
	err = user.UserDBRepository{}.CheckSignUp(u)
	user.HandleError(err)
	err = user.UserDBRepository{}.Create(u)
	user.HandleError(err)
	var tmplt = template.Must(template.ParseFiles("pkg/templates/index.html"))
	err = tmplt.Execute(w, nil)
	user.HandleError(err)
}
