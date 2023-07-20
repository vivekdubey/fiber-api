package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vivekdubey/fiber-api/internal/util"
)

func TestAddTransaction(t *testing.T) {
	book1 := addRandomBook(t)
	user1 := addRandomuser(t)

	arg := AddTransactionParams{
		BookID: book1.ID,
		UserID: user1.ID,
		Count:  util.RandomInt(1, 60),
	}

	tx, err := testQueries.AddTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tx)
	require.Equal(t, book1.ID, tx.BookID)
	require.Equal(t, user1.ID, tx.UserID)
	require.Equal(t, arg.Count, tx.Count)

}
