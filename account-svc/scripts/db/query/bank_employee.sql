-- name: CreateEmployee :one
INSERT INTO bank_employee (
    email, 
    password,
    full_name,
    role
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetEmployee :one
SELECT * FROM bank_employee
WHERE id = $1 LIMIT 1;

-- name: GetEmployeeByEmail :one
SELECT * FROM bank_employee
WHERE email = $1 LIMIT 1;

-- name: ListEmployee :many
SELECT * FROM bank_employee
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteEmployee :exec
DELETE FROM bank_employee WHERE id = $1;