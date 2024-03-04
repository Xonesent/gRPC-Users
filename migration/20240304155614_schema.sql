-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users
(
    user_id      serial        not null unique,
    user_name varchar(1024) not null unique
);

CREATE SCHEMA IF NOT EXISTS user_schema;
ALTER TABLE IF EXISTS public.users SET SCHEMA user_schema;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS users;
DROP TABLE user_schema.users;

-- +goose StatementEnd
