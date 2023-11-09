-- +goose Up
-- +goose StatementBegin
create table user_log (
    id int primary key generated always as identity,
    message text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user_log;
-- +goose StatementEnd
