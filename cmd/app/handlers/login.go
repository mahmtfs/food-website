package handlers

import (
	"html/template"
	"net/http"
	"working-project/internal/app/method/user"
)

func Login(w http.ResponseWriter, r *http.Request){
	var tmplt = template.Must(template.ParseFiles("pkg/templates/login.html"))
	err := tmplt.Execute(w, nil)
	user.HandleError(err)

}
