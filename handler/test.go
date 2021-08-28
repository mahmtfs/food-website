package handler

import "net/http"

func TestHandler(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session")
	untyped, ok := session.Values["email"]
	if !ok {
		return
	}
	email, ok := untyped.(string)
	if !ok {
		return
	}
	w.Write([]byte(email))
}
