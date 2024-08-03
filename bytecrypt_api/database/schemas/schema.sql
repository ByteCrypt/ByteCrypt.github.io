CREATE TABLE IF NOT EXISTS subscriptions (
    subscription_id SERIAL,
    email TEXT NOT NULL UNIQUE,
    name TEXT,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS roles (
    role_id INT UNIQUE NOT NULL,
    title text NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    user_id TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    name TEXT,
    username TEXT NOT NULL UNIQUE,
    tokens TEXT NOT NULL,
    role INT NOT NULL,
    PRIMARY KEY (user_id),
    FOREIGN KEY (role) REFERENCES roles (role_id)
);