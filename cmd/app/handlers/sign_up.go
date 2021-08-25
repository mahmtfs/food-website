package handlers

import (
	"html/template"
	"net/http"
	"working-project/internal/app/method/user"
)

func SignUp(w http.ResponseWriter, r *http.Request){
	var tmplt = template.Must(template.ParseFiles("pkg/templates/sign_up.html"))
	err := tmplt.Execute(w, nil)
	user.HandleError(err)

}
