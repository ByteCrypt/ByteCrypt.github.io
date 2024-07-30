-- Subscriptions
-- name: GetSubscriptions :many
SELECT * FROM subscriptions;

-- name: GetSubscriptionByEmail :one
SELECT * FROM subscriptions WHERE email = $1;

-- name: GetSubscriptionById :one
SELECT * FROM subscriptions WHERE id = $1;

-- name: AddSubscription :one
INSERT INTO
    subscriptions (email, name)
VALUES ($1, $2) RETURNING *;

-- name: DeleteSubscriptionId :exec
DELETE FROM subscriptions WHERE id = $1;

-- name: DeleteSubscriptionEmail :exec
DELETE FROM subscriptions WHERE email = $1;

-- Administrators
-- name: AddAdmin :one
INSERT INTO
    users (
        id,
        email,
        name,
        username,
        password,
        role
    )
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: DeleteAdmin :exec
DELETE FROM users WHERE id = $1;

-- name: GetAdmins :many
SELECT * FROM users;

-- name: GetAdminById :one
SELECT * FROM users WHERE id = $1;

-- name: GetAdminByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetAdminByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: UpdateAdminRole :one
UPDATE users SET role = $2 WHERE id = $1 RETURNING *;

-- name: UpdateAdminEmail :one
UPDATE users SET email = $2 WHERE id = $1 RETURNING *;

-- name: UpdateAdminName :one
UPDATE users SET name = $2 WHERE id = $1 RETURNING *;

-- name: UpdateAdminUsername :one
UPDATE users SET username = $2 WHERE id = $1 RETURNING *;

-- name: UpdateAdminPassword :one
UPDATE users SET password = $2 WHERE id = $1 RETURNING *;

-- Roles
-- name: AddRole :one
INSERT INTO roles (id, title) VALUES ($1, $2) RETURNING *;

-- name: DeleteRoleByTitle :exec
DELETE FROM roles WHERE title = $1;

-- name: DeleteRoleById :exec
DELETE FROM roles WHERE id = $1;

-- name: GetRoleById :one
SELECT * FROM roles WHERE id = $1;

-- name: GetRoleByTitle :one
SELECT * FROM roles WHERE title = $1;