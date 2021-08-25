package user

import (
	"fmt"
	"working-project/internal/app/model"
)

func (ur UserDBRepository) Create (u model.User) error{
	ur.db = OpenDB(&ur.db)
	defer ur.db.Close()
	if err := ur.db.Ping(); err != nil{
		return err
	}
	encryptedPassword, err := HashPassword(u.EncryptedPassword)
	if err != nil{
		return err
	}
	u.EncryptedPassword = encryptedPassword
	u.ID = GetLastID(ur.db) + 1
	stmt, err := ur.db.Query("INSERT INTO users (user_id, first_name, last_name, email, encrypted_password) VALUES (?, ?, ?, ?, ?);",
		u.ID, u.FirstName, u.LastName, u.Email, u.EncryptedPassword)
	if err != nil{
		return err
	}
	defer stmt.Close()
	fmt.Println("Added a user.")
	return nil
}
