package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type UserInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetRole() string
	GetStatus() bool
}

type User struct {
	ID       string `valid:"uuidv4"`
	Name     string `valid:"required,stringlength(5|20)"`
	Email    string `valid:"email,required"`
	Password string `valid:"required"`
	Role     string `valid:"optional"`
	Status   bool   `valid:"optional"`
}

func NewUser() *User {
	user := User{
		ID:     uuid.NewV4().String(),
		Role:   "Guest",
		Status: true,
	}

	return &user
}

func (u *User) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *User) Enable() error {
	u.Status = true
	return nil
}

func (u *User) Disable() error {
	u.Status = false
	return nil
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) GetStatus() bool {
	return u.Status
}
