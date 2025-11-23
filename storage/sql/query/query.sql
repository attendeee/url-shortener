-- name: GetAll :many
SELECT * FROM urls;

-- name: GetByShorthand :one
SELECT * FROM urls WHERE shorthand = ? LIMIT 1;

-- name: CreateShorthand :exec
INSERT INTO urls (longhand, shorthand) VALUES (?, ?);
