-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password_hash,
    first_name,
    last_name,
    bio
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;