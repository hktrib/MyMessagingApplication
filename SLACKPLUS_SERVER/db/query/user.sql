

-- name: CreateUser :one
INSERT into users (
    username, 
    hashed_password, 
    email
) VALUES ( $1, $2, $3
) RETURNING *;

-- name: SearchUserByEmail :one
SELECT EXISTS (SELECT 1 FROM users WHERE email = $1) AS email_exists;
-- name: SearchUserByUsername :one
SELECT EXISTS (SELECT * FROM users WHERE username = $1) AS username_exists;