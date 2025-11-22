-- name: GetAll :many
SELECT * FROM urls;

-- name: CreateShorthand :exec
INSERT INTO urls (longhand, shorthand) VALUES (?, ?);
