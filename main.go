package main

import (
	"food-website/handler"
	"food-website/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	route := mux.NewRouter()
	route.HandleFunc("/", middleware.AuthRequired(handler.HomePageHandler)).Methods("GET")
	route.HandleFunc("/login", handler.LoginPageHandler).Methods("GET")
	route.HandleFunc("/logout", handler.LogoutPageHandler).Methods("GET")
	route.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	route.HandleFunc("/register", handler.RegisterPageHandler).Methods("GET")
	route.HandleFunc("/register", handler.RegisterHandler).Methods("POST")
	route.HandleFunc("/test", handler.TestHandler).Methods("GET")
	http.Handle("/", route)
	http.ListenAndServe(":8000", nil)
}
