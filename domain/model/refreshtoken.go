package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Name   string `json:"name"`
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

type RefreshToken struct {
	Account   Account
	Token     string    `json:"refresh_token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type RefreshTokenInterface interface {
	GetAccount() Account
	GetToken() string
	GetCreatedAt() time.Time
	GetExpiredAt() time.Time
}

func NewRefreshToken() *RefreshToken {
	rt := RefreshToken{
		Token:     uuid.NewV4().String(),
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(3 * time.Hour), // 3 hours
	}

	return &rt
}

func (rt *RefreshToken) GetAccount() Account {
	return rt.Account
}

func (rt *RefreshToken) GetToken() string {
	return rt.Token
}

func (rt *RefreshToken) GetCreatedAt() time.Time {
	return rt.CreatedAt
}

func (rt *RefreshToken) GetExpiredAt() time.Time {
	return rt.ExpiredAt
}
