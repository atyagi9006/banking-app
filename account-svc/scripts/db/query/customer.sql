-- name: CreateCustomer :one
INSERT INTO customer (
    email,
    full_name,
    address,
    kyc_type,
    kyc_id
) VALUES (
  $1, $2, $3, $4,$5
) RETURNING *;

-- name: GetCustomer :one
SELECT * FROM customer
WHERE id = $1 LIMIT 1;

-- name: GetCustomerByEmail :one
SELECT * FROM customer
WHERE email = $1 LIMIT 1;

-- name: GetCustomerByFullName :one
SELECT * FROM customer
WHERE full_name = $1 LIMIT 1;

-- name: ListCustomer :many
SELECT * FROM customer
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCustomerKYC :one
UPDATE customer 
SET kyc_type = $2,
 kyc_id = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customer WHERE id = $1;