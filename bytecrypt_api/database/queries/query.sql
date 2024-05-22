-- name: GetSubscription :many
SELECT * FROM subscriptions;

-- name: GetSubscriptionEmail :one
SELECT * FROM subscriptions WHERE email = $1;

-- name: GetSubscriptionId :one
SELECT * FROM subscriptions WHERE id = $1;

-- name: AddSubscription :one
INSERT INTO
    subscriptions (email, name)
VALUES ($1, $2) RETURNING *;

-- name: DeleteSubscriptionId :exec
DELETE FROM subscriptions WHERE id = $1;

-- name: DeleteSubscriptionEmail :exec
DELETE FROM subscriptions WHERE email = $1;

-- name: AddAdmin :one
INSERT INTO
    administrator (
        id,
        email,
        name,
        username,
        password,
        role
    )
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: DeleteAdmin :exec
DELETE FROM administrator WHERE id = $1

-- name: ChangeAdminRole :one
