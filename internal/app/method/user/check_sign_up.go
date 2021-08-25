package user

import (
	"database/sql"
	"fmt"
	"log"
	"working-project/internal/app/model"
)

func (ur UserDBRepository) CheckSignUp(u model.User) error{
	ur.db = OpenDB(&ur.db)
	defer ur.db.Close()
	if err := ur.db.Ping(); err != nil{
		return err
	}
	fmt.Println(u.Email)
	err := ur.db.QueryRow("SELECT * FROM users WHERE email = ?;",
		u.Email).Scan()
	switch err {
	case sql.ErrNoRows:
		break
	default:
		log.Fatal("The user with such email is already registered.")
	}
	return nil
}
