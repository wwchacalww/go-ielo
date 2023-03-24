package model_test

import (
	"testing"
	"time"
	"wwchacalww/go-psyc/domain/model"

	"github.com/stretchr/testify/require"
)

func TestProfessional_IsValid(t *testing.T) {
	pro := model.NewProfessional()
	pro.Name = "Fulando de Tal"
	pro.Email = "fulando@ielo.com.br"
	pro.Register = "CRP-555/2001"
	pro.Gender = "Masculino"
	pro.Fone = "(61) 5555-5555"

	require.NotNil(t, pro.GetID())
	require.Equal(t, pro.GetName(), "Fulando de Tal")

	valid, err := pro.IsValid()

	require.False(t, valid)
	require.Equal(t, err.Error(), "BirthDate: non zero value required;Specialty: non zero value required")

	pro.Specialty = "Psicologia"
	pro.BirthDate = time.Date(1981, time.July, 20, 16, 47, 37, 1, time.Local)

	valid, err = pro.IsValid()
	require.Equal(t, pro.GetBirthDate().Day(), 20)
	require.True(t, valid)
	require.Nil(t, err)
}
