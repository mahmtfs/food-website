package user

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"working-project/internal/app/model"
)

func (ur UserDBRepository) CheckLogin(u model.User) error{
	ur.db = OpenDB(&ur.db)
	defer ur.db.Close()
	if err := ur.db.Ping(); err != nil{
		return err
	}
	var hashedpassword string
	err := ur.db.QueryRow("SELECT encrypted_password FROM users WHERE email = ?;",
		u.Email).Scan(&hashedpassword)
	switch err{
	case sql.ErrNoRows:
		log.Fatal("There is no user with such email or password.")
	default:
		err = bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(u.EncryptedPassword))
		if err != nil{
			log.Fatal("There is no user with such email or password.")
		}
	}
	fmt.Println("Logged in successfully.")
	return nil
}
