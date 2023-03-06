package model_test

import (
	"log"
	"testing"
	"wwchacalww/go-psyc/domain/model"

	"github.com/stretchr/testify/require"
)

func TestRefreshToken_Create(t *testing.T) {
	rt := model.NewRefreshToken()
	acc := model.Account{
		Name:   "FakeName",
		UserID: "fake-user-id",
		Email:  "fake@email.tst",
		Role:   "fake-role",
	}
	rt.Account = acc

	require.NotEmpty(t, rt.Token)
	require.Equal(t, rt.Account.Email, "fake@email.tst")
	log.Println(rt)
}
