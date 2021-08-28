package main

import (
	"food-website/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	route := mux.NewRouter()
	route.HandleFunc("/", handler.AuthRequired(handler.HomePageHandler)).Methods("GET")
	route.HandleFunc("/login", handler.LoginPageHandler).Methods("GET")
	route.HandleFunc("/logout", handler.LogoutPageHandler).Methods("GET")
	route.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	route.HandleFunc("/register", handler.RegisterPageHandler).Methods("GET")
	route.HandleFunc("/register", handler.RegisterHandler).Methods("POST")
	http.Handle("/", route)
	http.ListenAndServe(":8000", nil)
}
