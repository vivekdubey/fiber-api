package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vivekdubey/fiber-api/internal/util"
)

func addRandomTransaction(t *testing.T) Transaction {
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

	return tx
}
func TestAddTransaction(t *testing.T) {
	addRandomTransaction(t)
}

func TestGetTransaction(t *testing.T) {
	t1 := addRandomTransaction(t)

	t2, err := testQueries.GetTransaction(context.Background(), t1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, t2)
	require.Equal(t, t2.ID, t1.ID)
	require.Equal(t, t2.BookID, t1.BookID)
	require.Equal(t, t2.UserID, t1.UserID)

}
