package handler

import (
	"food-website/model"
	"food-website/repository"
	"food-website/service"
	"net/http"
	"time"
)

func RegisterPageHandler(w http.ResponseWriter, r *http.Request){
	err := service.CheckTokens(w, r)
	switch {
	case err == nil:
		http.Redirect(w, r, "/", 302)
		return
	case err.Error() == "access token expired":
		errUpdate := service.UpdateAccessToken(w, r, user)
		if errUpdate != nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	case err.Error() ==	"all tokens expired":
		break
	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = templates.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to execute a template"))
		return
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	var u model.User
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to parse form"))
		return
	}
	db, err := service.GetDB()
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	u.ID, err = service.GetLastID(db)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = db.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to close the database"))
		return
	}
	u.FirstName = r.PostForm.Get("first_name")
	err = service.IfEmpty(u.FirstName)
	if err != nil {
		templates.ExecuteTemplate(w, "register.html", err.Error())
		return
	}
	u.LastName = r.PostForm.Get("last_name")
	err = service.IfEmpty(u.LastName)
	if err != nil {
		templates.ExecuteTemplate(w, "register.html", err.Error())
		return
	}
	u.Email = r.PostForm.Get("email")
	err = service.IfEmpty(u.Email)
	if err != nil {
		templates.ExecuteTemplate(w, "register.html", err.Error())
		return
	}
	err = service.IfValidEmail(u.Email)
	if err != nil{
		templates.ExecuteTemplate(w, "register.html", "the user with such email already exists")
		return
	}
	password := r.PostForm.Get("password")
	err = service.IfEmpty(password)
	if err != nil {
		templates.ExecuteTemplate(w, "register.html", err.Error())
		return
	}
	u.HashedPassword, err = service.HashPassword(password)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to hash a password"))
		return
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = u.CreatedAt
	err = repository.UserRepository{}.Create(u)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create a user"))
		return
	}
	http.Redirect(w, r, "/login", 302)
	return
}
