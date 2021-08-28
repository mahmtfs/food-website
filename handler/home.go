package handler

import (
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request){
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to execute template"))
		return
	}
}
