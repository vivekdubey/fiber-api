// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: transaction.sql

package db

import (
	"context"
)

const addTransaction = `-- name: AddTransaction :one
INSERT INTO transactions (
  user_id, book_id, count
) VALUES (
  $1, $2, $3
)
RETURNING id, user_id, book_id, created_at, count
`

type AddTransactionParams struct {
	UserID int64 `json:"user_id"`
	BookID int64 `json:"book_id"`
	Count  int64 `json:"count"`
}

func (q *Queries) AddTransaction(ctx context.Context, arg AddTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, addTransaction, arg.UserID, arg.BookID, arg.Count)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.BookID,
		&i.CreatedAt,
		&i.Count,
	)
	return i, err
}
