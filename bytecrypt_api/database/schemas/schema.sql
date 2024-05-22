CREATE TABLE IF NOT EXISTS subscriptions (
    id bigserial PRIMARY KEY,
    email text NOT NULL UNIQUE,
    name text NOT NULL,
)

CREATE TABLE IF NOT EXISTS role (
    id SERIAL PRIMARY KEY,
    title text NOT NULL UNIQUE,
)

CREATE TABLE IF NOT EXISTS administrator (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role INT NOT NULL,
    FOREIGN KEY (role) REFERENCES role(id)
)