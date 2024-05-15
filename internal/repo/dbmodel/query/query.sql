-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    password,
    address
) VALUES (
    $1, $2, $3, $4
) RETURNING *;