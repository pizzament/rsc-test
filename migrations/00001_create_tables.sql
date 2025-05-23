-- +goose Up
-- +goose StatementBegin
create table logs (
    banner_id integer not null,
    time_stamp timestamp without time zone default now(),
    count integer default 1,
    PRIMARY KEY (banner_id, time_stamp)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table logs;
-- +goose StatementEnd
