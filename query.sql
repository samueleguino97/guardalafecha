-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;
-- name: GetToken :one
SELECT * FROM tokens
WHERE token = ? LIMIT 1;
