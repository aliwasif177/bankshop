package db

import (
	"context"
	"testing"
	"time"

	"github.com/aliwasif177/bankshop/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword("secret")
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		Email:          util.RandomEmail(),
		FullName:       util.RandomOwner(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)

	return user
}

func TestCreateUser(t *testing.T) {

	CreateRandomUser(t)

}

func TestGetUser(t *testing.T) {
	user := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user2.HashedPassword, user.HashedPassword)
	require.Equal(t, user2.Email, user.Email)
	require.Equal(t, user2.Username, user.Username)
	require.Equal(t, user2.Username, user.Username)
	require.Equal(t, user2.FullName, user.FullName)
	require.WithinDuration(t, user.CreatedAt.Time, user2.CreatedAt.Time, time.Second)

}
