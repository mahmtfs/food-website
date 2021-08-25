package repos

import "working-project/internal/app/model"

type UserRepositoryI interface {
	Create (u model.User) error
	CheckLogin (u model.User) error
	CheckSignUp (u model.User) error
	Delete (id int) error
}
