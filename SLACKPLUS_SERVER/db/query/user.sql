

-- name: CreateUser :one
INSERT into users (
    username, 
    hashed_password, 
    email
) VALUES ( $1, $2, $3
) RETURNING *;