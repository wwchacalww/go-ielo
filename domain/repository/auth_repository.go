package repository

import "wwchacalww/go-psyc/domain/model"

type AuthenticateOutput struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthRepositoryInterface interface {
	Authenticate(email, password string) (AuthenticateOutput, error)
	RefreshToken(email, refresh_token string) (AuthenticateOutput, error)
}

type AuthPersistence interface {
	CreateRefreshToken(rt model.RefreshTokenInterface) (model.RefreshTokenInterface, error)
	CheckAccount(email string) (model.UserInterface, error)
	CheckRefreshToken(user_id, refresh_token string) error
}
