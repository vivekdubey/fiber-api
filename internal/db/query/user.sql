-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: AddUser :one
INSERT INTO users (
  first_name
) VALUES (
  $1
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateLastActivity :one
UPDATE users
  set last_activity = $2
WHERE id = $1
RETURNING *;