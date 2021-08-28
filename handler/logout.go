package handler

import (
	"net/http"
	"time"
)

func LogoutPageHandler(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie("access-token")
	if err != nil{
		if err != http.ErrNoCookie {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	} else{
		cookie.Expires = time.Now()
		http.SetCookie(w, cookie)
	}

	cookie, err = r.Cookie("refresh-token")
	if err != nil{
		if err != http.ErrNoCookie {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	} else {
		cookie.Expires = time.Now()
		http.SetCookie(w, cookie)
	}

	http.Redirect(w, r, "/login", 302)
	return
}
