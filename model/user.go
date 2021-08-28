package model

import "time"

type User struct{
	ID int
	FirstName string
	LastName string
	Email string
	HashedPassword string
	CreatedAt time.Time
	UpdatedAt time.Time
}
