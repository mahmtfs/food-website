package handler

import (
	"food-website/service"
	"net/http"
)

func AuthRequired(handle http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		err := service.CheckTokens(w, r)
		if err != nil {
			switch err.Error() {
			case "all tokens expired":
				http.Redirect(w, r, "/logout", 302)
				return
			case "access token expired":
				errUpdate := service.UpdateAccessToken(w, r, user)
				if errUpdate != nil{
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(errUpdate.Error()))
					return
				}
				break
			default:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
		}
		/*_, errAccess := r.Cookie("access-token")
		if errAccess != nil{
			if errAccess != http.ErrNoCookie{
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errAccess.Error()))
				return
			}
		}
		_, errRefresh := r.Cookie("refresh-token")
		if errRefresh != nil{
			if errRefresh != http.ErrNoCookie{
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errRefresh.Error()))
				return
			}
		}
		if errAccess == http.ErrNoCookie && errRefresh == http.ErrNoCookie{
			http.Redirect(w, r, "/login", 302)
			return
		}*/
		handle.ServeHTTP(w, r)
	}
}
