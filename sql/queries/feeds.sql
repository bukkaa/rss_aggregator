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

