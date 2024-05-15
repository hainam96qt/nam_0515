-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    password,
    address
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

/* name: GetUserByEmail :one */
SELECT * FROM users WHERE email = $1 ;