package repository

import (
	"errors"
	"food-website/model"
	"food-website/service"
)

type UserRepositoryI interface {
	Create(u model.User) error
	Edit(u model.User)
	Delete(u model.User)
}

type UserRepository struct {
}

func (ur UserRepository) Create(u model.User) error{
	db, err := service.GetDB()
	if err != nil{
		return err
	}
	q, err := db.Query("INSERT INTO users (id, first_name, last_name, email, hashed_password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?);", u.ID, u.FirstName, u.LastName, u.Email, u.HashedPassword, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return errors.New("failed to create a user")
	}
	q.Close()
	db.Close()
	return nil
}
