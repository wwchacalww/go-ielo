package model_test

import (
	"log"
	"testing"
	"wwchacalww/go-psyc/domain/model"

	"github.com/stretchr/testify/require"
)

func TextIsValid_Valid(t *testing.T) {
	user := model.NewUser()
	user.Name = "Test Name"
	user.Email = "test@email.com"
	user.Password = "passwordTest"
	user.Role = "TestRole"

	log.Println(user.GetID())
	require.NotNil(t, user.GetID())
	require.Equal(t, user.GetName(), "Test Name")
	require.Equal(t, user.GetRole(), "TestRole")

	valid, err := user.IsValid()

	require.True(t, valid)
	require.Nil(t, err)
}

func TestIsValid_Invalid(t *testing.T) {
	user := model.NewUser()
	user.Name = "Tes"
	user.Email = "wrong-email"
	user.Password = "pas"

	require.NotNil(t, user.GetID())

	valid, err := user.IsValid()
	require.False(t, valid)
	require.Equal(t, err.Error(), "Email: wrong-email does not validate as email;Name: Tes does not validate as stringlength(5|20)")
	require.Equal(t, user.GetRole(), "Guest")
}
