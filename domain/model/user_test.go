package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/giziellerodrigues/imersao/codepix-go/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewUser(t *testing.T) {

	email := "joao@provedor.com"
	name := "Joao"

	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	_, err = model.NewUser("", "")
	require.NotNil(t, err)
}