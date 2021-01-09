-- name: Createusers :one
INSERT INTO users (
    email, 
    password,
    role
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: Getusers :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetusersByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: Listusers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: Deleteusers :exec
DELETE FROM users WHERE id = $1;