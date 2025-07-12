-- +goose Up

create table users
(
    id         UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL,
    name       TEXT      NOT NULL
);

-- +goose Down
DROP TABLE users;
