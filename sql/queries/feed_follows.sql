-- name: CreateFeedFollow :one
insert into feed_follows(id, user_id, feed_id, created_at, updated_at)
values ($1, $2, $3, (now() at time zone 'utc'), (now() at time zone 'utc'))
returning *;

