package service

import (
	"errors"
	"food-website/model"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var jwtKey = []byte("secret_cyka")

type Claims struct {
	Email string
	HashedPassword string
	jwt.StandardClaims
}

func CheckTokens(w http.ResponseWriter, r *http.Request) error{
	cookie, err := r.Cookie("access-token")
	if err != nil{
		if err != http.ErrNoCookie{
			return errors.New("internal server error")
		}
		tokenAlive := CheckRefreshToken(w, r)
		if !tokenAlive{
			return errors.New("all tokens expired")
		} else{
			return errors.New("access token expired")
		}
	}
	tokenString := cookie.Value

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error){
			return jwtKey, nil
		})
	if err != nil{
		if err == jwt.ErrSignatureInvalid{
			return errors.New("internal server error")
		}
	}

	if !token.Valid{
		return errors.New("internal server error")
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) <= 0 * time.Second{
		cookie.Expires = time.Now()
		http.SetCookie(w, cookie)
	}
	return nil
}

func CheckRefreshToken(w http.ResponseWriter, r *http.Request) bool{
	_, err := r.Cookie("refresh-token")
	if err != nil{
		return false
	}
	return true
}

func UpdateAccessToken(w http.ResponseWriter, r *http.Request, user model.User) error{
	accessExpirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Email: user.Email,
		HashedPassword: user.HashedPassword,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil{
		return errors.New("internal server error")
	}
	http.SetCookie(w, &http.Cookie{
		Name: "access-token",
		Value: tokenString,
		Expires: accessExpirationTime,
	})
	return nil
}

func UpdateRefreshToken(w http.ResponseWriter, r *http.Request, user model.User) error{
	refreshExpirationTime := time.Now().Add(time.Minute * 30)
	claims := &Claims{
		Email: user.Email,
		HashedPassword: user.HashedPassword,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil{
		return errors.New("internal server error")
	}
	http.SetCookie(w, &http.Cookie{
		Name: "refresh-token",
		Value: tokenString,
		Expires: refreshExpirationTime,
	})
	return nil
}