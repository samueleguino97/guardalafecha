// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
)

const getTenantBySlug = `-- name: GetTenantBySlug :one
SELECT id, name, slug FROM tenants
WHERE slug = ? LIMIT 1
`

func (q *Queries) GetTenantBySlug(ctx context.Context, slug string) (Tenant, error) {
	row := q.db.QueryRowContext(ctx, getTenantBySlug, slug)
	var i Tenant
	err := row.Scan(&i.ID, &i.Name, &i.Slug)
	return i, err
}

const getToken = `-- name: GetToken :one
SELECT token, expires_at, user_id, token_type, tenant_id, "primary", "foreign" FROM tokens
WHERE token = ? LIMIT 1
`

func (q *Queries) GetToken(ctx context.Context, token string) (Token, error) {
	row := q.db.QueryRowContext(ctx, getToken, token)
	var i Token
	err := row.Scan(
		&i.Token,
		&i.ExpiresAt,
		&i.UserID,
		&i.TokenType,
		&i.TenantID,
		&i.Primary,
		&i.Foreign,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name, slug, reserved_spots, tenant_id FROM users
WHERE id = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.ReservedSpots,
		&i.TenantID,
	)
	return i, err
}
