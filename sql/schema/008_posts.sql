-- +goose Up
create table posts (
    id UUID primary key,
    title text not null,
    description text,
    published_at timestamp not null,
    url text not null unique,
    feed_id UUID not null references feeds(id) on delete cascade,
    created_at timestamp not null,
    updated_at timestamp not null
);

-- +goose Down
drop table posts;