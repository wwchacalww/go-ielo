package repository

import (
	"fmt"
	"log"
	"wwchacalww/go-psyc/domain/model"
	"wwchacalww/go-psyc/domain/repository"
	"wwchacalww/go-psyc/domain/utils"
)

type AuthRepository struct {
	Persistence repository.AuthPersistence
}

func NewAuthRepository(persistence repository.AuthPersistence) *AuthRepository {
	return &AuthRepository{Persistence: persistence}
}

func (repo *AuthRepository) Authenticate(email, password string) (repository.AuthenticateOutput, error) {
	user, err := repo.Persistence.CheckAccount(email)
	if err != nil {
		return repository.AuthenticateOutput{}, fmt.Errorf("email or password invalid")
	}

	check := utils.CheckPasswordHash(password, user.GetPassword())
	if check != true {
		return repository.AuthenticateOutput{}, fmt.Errorf("email or password invalid")
	}

	rt := model.NewRefreshToken()
	acc := model.Account{
		Name:   user.GetName(),
		UserID: user.GetID(),
		Email:  user.GetEmail(),
		Role:   user.GetRole(),
	}
	rt.Account = acc

	refreshToken, err := repo.Persistence.CreateRefreshToken(rt)
	if err != nil {
		return repository.AuthenticateOutput{}, err
	}

	secret_JWT := "secret_jwt" // put on .env
	token, err := utils.CreateJWToken(user.GetName(), user.GetEmail(), user.GetRole(), secret_JWT)
	log.Println(token)
	if err != nil {
		return repository.AuthenticateOutput{}, err
	}

	result := repository.AuthenticateOutput{
		Token:        token,
		RefreshToken: refreshToken.GetToken(),
	}

	return result, nil
}

func (repo *AuthRepository) RefreshToken(email, refresh_token string) (repository.AuthenticateOutput, error) {
	user, err := repo.Persistence.CheckAccount(email)
	if err != nil {
		return repository.AuthenticateOutput{}, err
	}
	err = repo.Persistence.CheckRefreshToken(user.GetID(), refresh_token)
	if err != nil {
		return repository.AuthenticateOutput{}, err
	}

	acc := model.Account{
		Name:   user.GetName(),
		UserID: user.GetID(),
		Email:  user.GetEmail(),
		Role:   user.GetRole(),
	}
	rt := model.NewRefreshToken()
	rt.Account = acc

	refreshToken, err := repo.Persistence.CreateRefreshToken(rt)
	if err != nil {
		return repository.AuthenticateOutput{}, err
	}

	secret_JWT := "secret_jwt" // put on .env
	token, err := utils.CreateJWToken(acc.Name, acc.Email, acc.Role, secret_JWT)
	log.Println(token)
	if err != nil {
		return repository.AuthenticateOutput{}, err
	}

	result := repository.AuthenticateOutput{
		Token:        token,
		RefreshToken: refreshToken.GetToken(),
	}

	return result, nil
}
