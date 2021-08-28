package handler

import "net/http"

func LogoutPageHandler(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session")
	delete(session.Values, "email")
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}
