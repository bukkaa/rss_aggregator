-- name: CreateFeedFollow :one
insert into feed_follows(id, user_id, feed_id, created_at, updated_at)
values ($1, $2, $3, (now() at time zone 'utc'), (now() at time zone 'utc'))
returning *;

-- name: DeleteFeedFollow :exec
delete from feed_follows where id = $1 and user_id = $2;
