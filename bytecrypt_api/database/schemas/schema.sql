CREATE TABLE IF NOT EXISTS subscriptions (
    id bigserial,
    email text NOT NULL UNIQUE,
    name text NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL,
    title text NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS administrators (
    id TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (role) REFERENCES roles (id)
);