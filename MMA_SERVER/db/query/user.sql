

-- name: CreateUser :one
INSERT into users (
    username, 
    hashed_password, 
    email
) VALUES (@username, @hashed_password, @email) RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET 
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    email = COALESCE(sqlc.narg(email), email),
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
    is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
WHERE
    username = @username
RETURNING *;

-- name: SearchUserByEmail :one
SELECT EXISTS (SELECT 1 FROM users WHERE email = @email) AS email_exists;
-- name: SearchUserByUsername :one
SELECT EXISTS (SELECT * FROM users WHERE username = @username) AS username_exists;

-- name: CreateVerifyEmailsRecord :one
INSERT into verify_emails (
    username,
    email, 
    secret_code
) VALUES (@username, @email, @secret_code) RETURNING *;

-- name: SearchRecordToVerify :one
SELECT * FROM verify_emails WHERE (username = @username AND secret_code = @secret_code); 

-- name: UpdateVerifyEmail :one
UPDATE verify_emails
SET
    is_used = TRUE
WHERE
    username = @username
    AND secret_code = @secret_code
    AND is_used = FALSE
RETURNING *;