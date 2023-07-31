-- name: AddTransaction :one
INSERT INTO transactions (
  user_id, book_id, count
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;
