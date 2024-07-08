-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;
-- name: GetToken :one
SELECT * FROM tokens
WHERE token = ? LIMIT 1;
-- name: GetTenantBySlug :one
SELECT * FROM tenants
WHERE slug = ? LIMIT 1;
