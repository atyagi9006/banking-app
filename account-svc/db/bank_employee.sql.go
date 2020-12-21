// Code generated by sqlc. DO NOT EDIT.
// source: bank_employee.sql

package db

import (
	"context"
)

const createEmployee = `-- name: CreateEmployee :one
INSERT INTO bank_employee (
    email, 
    password,
    full_name,
    role
) VALUES (
  $1, $2, $3, $4
) RETURNING id, email, password, full_name, role, created_at
`

type CreateEmployeeParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (BankEmployee, error) {
	row := q.db.QueryRowContext(ctx, createEmployee,
		arg.Email,
		arg.Password,
		arg.FullName,
		arg.Role,
	)
	var i BankEmployee
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.FullName,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM bank_employee WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getEmployee = `-- name: GetEmployee :one
SELECT id, email, password, full_name, role, created_at FROM bank_employee
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEmployee(ctx context.Context, id int64) (BankEmployee, error) {
	row := q.db.QueryRowContext(ctx, getEmployee, id)
	var i BankEmployee
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.FullName,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const getEmployeeByEmail = `-- name: GetEmployeeByEmail :one
SELECT id, email, password, full_name, role, created_at FROM bank_employee
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetEmployeeByEmail(ctx context.Context, email string) (BankEmployee, error) {
	row := q.db.QueryRowContext(ctx, getEmployeeByEmail, email)
	var i BankEmployee
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.FullName,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const listEmployee = `-- name: ListEmployee :many
SELECT id, email, password, full_name, role, created_at FROM bank_employee
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListEmployeeParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListEmployee(ctx context.Context, arg ListEmployeeParams) ([]BankEmployee, error) {
	rows, err := q.db.QueryContext(ctx, listEmployee, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BankEmployee
	for rows.Next() {
		var i BankEmployee
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.FullName,
			&i.Role,
			&i.CreatedAt,
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
