package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type ProfessionalInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetEmail() string
	GetBirthDate() time.Time
	GetRegister() string
	GetGender() string
	GetSpecialty() string
	GetDescription() string
	GetFone() string
	GetStatus() bool
	GetUser() UserInterface
}

type Professional struct {
	ID          string    `valid:"uuidv4"`
	Name        string    `valid:"required"`
	Email       string    `valid:"email,required"`
	BirthDate   time.Time `valid:"required"`
	Register    string    `valid:"optional"`
	Gender      string    `valid:"required"`
	Specialty   string    `valid:"required"`
	Description string    `valid:"optional"`
	Fone        string    `valid:"required"`
	Status      bool      `valid:"optional"`
	User        UserInterface
}

func NewProfessional() *Professional {
	professional := Professional{
		ID:     uuid.NewV4().String(),
		Status: true,
	}

	return &professional
}

func (p *Professional) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Professional) Enable() error {
	p.Status = true
	return nil
}

func (p *Professional) Disable() error {
	p.Status = false
	return nil
}

func (p *Professional) GetID() string {
	return p.ID
}

func (p *Professional) GetName() string {
	return p.Name
}

func (p *Professional) GetEmail() string {
	return p.Email
}

func (p *Professional) GetBirthDate() time.Time {
	return p.BirthDate
}

func (p *Professional) GetRegister() string {
	return p.Register
}

func (p *Professional) GetGender() string {
	return p.Gender
}

func (p *Professional) GetSpecialty() string {
	return p.Specialty
}

func (p *Professional) GetDescription() string {
	return p.Description
}

func (p *Professional) GetFone() string {
	return p.Fone
}

func (p *Professional) GetStatus() bool {
	return p.Status
}

func (p *Professional) GetUser() UserInterface {
	return p.User
}
