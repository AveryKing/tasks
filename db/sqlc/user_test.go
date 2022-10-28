package db

import (
	"context"
	"github.com/AveryKing/tasks/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createTestUser(t *testing.T) (user User) {
	arg := CreateUserParams{
		Username: util.RandomString(6),
		Password: util.RandomString(6),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)

	return
}

func TestCreateUser(t *testing.T) {
	createTestUser(t)
}

func TestGetUser(t *testing.T) {
	testUser := createTestUser(t)
	user, err := testQueries.GetUser(context.Background(), testUser.Username)

	require.NoError(t, err)
	require.Equal(t, testUser.Username, user.Username)
	require.Equal(t, testUser.Password, user.Password)
}
