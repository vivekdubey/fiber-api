package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vivekdubey/fiber-api/internal/util"
)

func addRandomuser(t *testing.T) User {
	first_name := util.RandomFirstName()
	user, err := testQueries.AddUser(context.Background(), first_name)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.FirstName, first_name)
	return user
}

func TestAddUser(t *testing.T) {
	addRandomBook(t)
}

func TestGetuser(t *testing.T) {
	user1 := addRandomuser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.FirstName, user2.FirstName)
}

func TestDeleteUser(t *testing.T) {
	user1 := addRandomuser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
