// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: book.sql

package db

import (
	"context"
)

const addBook = `-- name: AddBook :one
INSERT INTO books (
  title, author, description, count
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, title, author, description, created_at, count
`

type AddBookParams struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Count       int64  `json:"count"`
}

func (q *Queries) AddBook(ctx context.Context, arg AddBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, addBook,
		arg.Title,
		arg.Author,
		arg.Description,
		arg.Count,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Description,
		&i.CreatedAt,
		&i.Count,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1
`

func (q *Queries) DeleteBook(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBook, id)
	return err
}

const getBook = `-- name: GetBook :one
SELECT id, title, author, description, created_at, count FROM books
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBook(ctx context.Context, id int64) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Description,
		&i.CreatedAt,
		&i.Count,
	)
	return i, err
}

const getBooks = `-- name: GetBooks :many
SELECT id, title, author, description, created_at, count FROM books
ORDER BY title LIMIT $1 OFFSET $2
`

type GetBooksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetBooks(ctx context.Context, arg GetBooksParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, getBooks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Author,
			&i.Description,
			&i.CreatedAt,
			&i.Count,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBook = `-- name: UpdateBook :one
UPDATE books
  set title = $2,
  description = $3,
  author = $4,
  count = $5
WHERE id = $1
RETURNING id, title, author, description, created_at, count
`

type UpdateBookParams struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Count       int64  `json:"count"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, updateBook,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Author,
		arg.Count,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Description,
		&i.CreatedAt,
		&i.Count,
	)
	return i, err
}

const updateBookCount = `-- name: UpdateBookCount :one
UPDATE books
  set count = $2
WHERE id = $1
RETURNING id, title, author, description, created_at, count
`

type UpdateBookCountParams struct {
	ID    int64 `json:"id"`
	Count int64 `json:"count"`
}

func (q *Queries) UpdateBookCount(ctx context.Context, arg UpdateBookCountParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, updateBookCount, arg.ID, arg.Count)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Description,
		&i.CreatedAt,
		&i.Count,
	)
	return i, err
}
