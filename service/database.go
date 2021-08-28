package service

import (
	"database/sql"
	"errors"
	"food-website/model"
	"golang.org/x/crypto/bcrypt"
	"html/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func GetDB()(sql.DB, error){
	db, err := sql.Open("mysql", "rinat:jojo1337@tcp(127.0.0.1:3306)/website")
	if err != nil {
		return *db, errors.New("failed to open the database")
	}
	err = db.Ping()
	if err != nil {
		return *db, errors.New("failed to connect to the database")
	}
	return *db, nil
}

func GetLastID(db sql.DB) (int, error){
	var id int
	rows, err := db.Query("SELECT id FROM users;")
	if err == sql.ErrNoRows{
		rows.Close()
		return 1, nil
	}
	if err != nil {
		return 0, err
	}
	for rows.Next(){
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	rows.Close()
	return id + 1, nil
}

func IfValidEmail(email string) error{
	var db sql.DB
	db, err := GetDB()
	if err != nil{
		return errors.New("failed to open the database")
	}
	err = db.QueryRow("SELECT * FROM users WHERE email = ?;", email).Scan()
	switch err {
	case sql.ErrNoRows:
		break
	default:
		if err != nil {
			return errors.New("there is no user with such email")
		}
	}
	db.Close()
	return nil
}

func IfEmpty(str string) error{
	if str == ""{
		return errors.New("empty data is not allowed")
	}
	return nil
}

func CheckLogin (user *model.User) error{
	var db sql.DB
	var hashedPassword string
	db, err := GetDB()
	if err != nil{
		return errors.New("faled to open the database")
	}
	err = db.QueryRow("SELECT hashed_password FROM users WHERE email = ?;", user.Email).Scan(&hashedPassword)
	switch err {
	case sql.ErrNoRows:
		return errors.New("invalid email or password")
	default:
		break
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.HashedPassword))
	if err != nil {
		return errors.New("invalid email or password")
	}
	user.HashedPassword = hashedPassword
	db.Close()
	return nil
}
