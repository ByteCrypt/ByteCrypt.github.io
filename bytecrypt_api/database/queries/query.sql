-- name: GetSubscription :many
SELECT * FROM subscriptions;

-- name: GetSubscriptionEmail :one
SELECT * FROM subscriptions
WHERE email = $1;

-- name: GetSubscriptionId :one
SELECT * FROM subscriptions
WHERE id = $1;

-- name: AddSubscription :one
INSERT INTO subscriptions (
    email, name
) VALUES (
    $1, $2
)
RETURNING *;

-- name: DeleteSubscriptionId :exec
DELETE FROM subscriptions 
WHERE id = $1;

-- name: DeleteSubscriptionEmail :exec
DELETE FROM subscriptions 
WHERE email = $1;