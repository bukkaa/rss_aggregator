-- name: CreateFeed :one
insert into feeds (id, name, url, user_id, created_at, updated_at)
values ($1, $2, $3, $4, (now() at time zone 'utc'), (now() at time zone 'utc'))
returning *;

-- name: GetFeedsByUserId :many
select * from feeds 
where user_id = $1 
order by created_at desc;

-- name: GetAllFeeds :many
select * from feeds 
order by created_at desc;

-- name: GetNextFeedsToFetch :many
select * from feeds
order by last_fetched_at asc nulls first
limit $1;

-- name: MarkFeedAsFetched :one

update feeds
set last_fetched_at = now() at time zone 'utc',
    updated_at = now() at time zone 'utc'
where id = $1
returning *;