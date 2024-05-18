CREATE TABLE IF NOT EXISTS subscriptions (
    id bigserial,
    email text NOT NULL UNIQUE,
    name text NOT NULL,
    PRIMARY KEY(id)
)