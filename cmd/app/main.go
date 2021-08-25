package main

import (
	"net/http"
	"working-project/cmd/app/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Login)
	http.HandleFunc("/sign_up", handlers.SignUp)
	http.HandleFunc("/login_user", handlers.LoginUser)
	http.HandleFunc("/sign_up_user", handlers.SignUpUser)
	http.ListenAndServe(":8000", nil)
}
