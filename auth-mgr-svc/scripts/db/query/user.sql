-- name: CreateUser :one
INSERT INTO users (
    email, 
    password,
    role
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;