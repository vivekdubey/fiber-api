package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type ReturnBookTxParams struct {
	BookID       int64     `json:"book_id"`
	UserID       int64     `json:"user_id"`
	LastActivity time.Time `json:"last_activity"`
}

type ReturnBookTxResult struct {
	Transaction  Transaction `json:"transaction"`
	LastActivity User        `json:"last_activity"`
}

func (store *Store) ReturnBookTx(ctx context.Context, arg ReturnBookTxParams) (ReturnBookTxResult, error) {
	var result ReturnBookTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transaction, err = q.AddTransaction(ctx, AddTransactionParams{
			UserID: arg.UserID,
			BookID: arg.BookID,
			Count:  -1,
		})
		if err != nil {
			return err
		}

		// Update Last activity column
		// Update book count

		// result.LastActivity, err = q.UpdateLastActivity(ctx, UpdateLastActivityParams{
		// 	ID:           arg.UserID,
		// 	LastActivity: time.Now().Local(),
		// })

		// if err != nil {
		// 	return err
		// }
		return nil
	})
	return result, err
}
