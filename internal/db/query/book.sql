-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: GetBooks :many
SELECT * FROM books
ORDER BY title LIMIT $1 OFFSET $2;

-- name: AddBook :one
INSERT INTO books (
  title, author, description, count
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: UpdateBook :one
UPDATE books
  set title = $2,
  description = $3,
  author = $4,
  count = $5
WHERE id = $1
RETURNING *;

-- name: UpdateBookCount :one
UPDATE books
  set count = $2
WHERE id = $1
RETURNING *;