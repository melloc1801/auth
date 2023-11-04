-- +goose Up
-- +goose StatementBegin
select 'up SQL query';

create type role as enum ('USER', 'ADMIN');

create table "user" (
    id serial primary key,
    name varchar(32) not null,
    email varchar(128) not null,
    role role not null,
    password varchar(32) not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
drop table "user";
drop type role;
-- +goose StatementEnd
