package repository

import "wwchacalww/go-psyc/domain/model"

type UserInput struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   bool   `json:"status"`
}

type UserRepositoryInterface interface {
	Create(input UserInput) (model.UserInterface, error)
	FindById(id string) (model.UserInterface, error)
	FindByEmail(email string) (model.UserInterface, error)
	List() ([]model.UserInterface, error)
	ChangePassword(id, pwd string) error
	ChangeRole(id, role string) error
	ChangeAvatarUrl(email string) (string, error)
}

type UserPersistenceInterface interface {
	Create(user model.UserInterface) error
	FindById(id string) (model.UserInterface, error)
	FindByEmail(email string) (model.UserInterface, error)
	List() ([]model.UserInterface, error)
	ChangePassword(id, pwd string) error
	ChangeRole(id, role string) error
	ChangeAvatarUrl(email, avatar_url string) error
}
