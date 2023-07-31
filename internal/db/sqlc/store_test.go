package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestReturnBookTx(t *testing.T) {
	store := NewStore(testDB)
	book1 := addRandomBook(t)
	user1 := addRandomuser(t)

	r, e := store.ReturnBookTx(context.Background(), ReturnBookTxParams{
		BookID:       book1.ID,
		UserID:       user1.ID,
		LastActivity: time.Now().Local(),
	})

	require.NoError(t, e)
	require.NotEmpty(t, r)

	// results := make(chan ReturnBookTxResult)
	// errs := make(chan error)
	// fmt.Println("Out of go routines")
	// for i := 0; i < 1; i++ {
	// 	go func() {
	// 		fmt.Println("In go routines")
	// 		r, e := store.ReturnBookTx(context.Background(), ReturnBookTxParams{
	// 			BookID:       book1.ID,
	// 			UserID:       user1.ID,
	// 			LastActivity: time.Now().Local(),
	// 		})
	// 		fmt.Printf("Results: %v", r.Transaction.BookID)
	// 		fmt.Printf("Error: %v", e)
	// 		results <- r
	// 		errs <- e
	// 	}()
	// }

	// fmt.Println("Came out of go routines")
	// for i := 0; i < 1; i++ {
	// 	fmt.Println("Processing channnel err")
	// 	e := <-errs
	// 	require.NoError(t, e)
	// 	fmt.Println("Processing channnel results")
	// 	r := <-results
	// 	require.NotEmpty(t, r)
	// 	require.Equal(t, book1.ID, r.Transaction.BookID)
	// 	require.Equal(t, user1.ID, r.Transaction.UserID)
	// 	_, e = store.GetTransaction(context.Background(), r.Transaction.ID)
	// 	require.NoError(t, e)
	// }
	// fmt.Println("Processed channnel")

}
