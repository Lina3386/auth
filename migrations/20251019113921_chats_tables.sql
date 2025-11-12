-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    role INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL default now(),
    updated_at TIMESTAMP NOT NULL default now()
);


-- +goose Down
DROP TABLE users;
