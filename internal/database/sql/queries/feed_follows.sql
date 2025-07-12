-- name: CreateFeedFollow :one
WITH inserted AS (
INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
    RETURNING
        *)
    SELECT
        inserted.*,
        users.name AS user_name,
        feeds.name AS feed_name
    FROM
        inserted
        JOIN users ON inserted.user_id = users.id
        JOIN feeds ON inserted.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.*,
    users.name AS user_name,
    feeds.name AS feed_name
FROM
    feed_follows
    JOIN users ON feed_follows.user_id = users.id
    JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE
    feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE user_id = $1
    AND feed_id = $2;

