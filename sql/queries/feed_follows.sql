-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (
    id,
    created_at,
    updated_at,
    user_id,
    feed_id
  ) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
  ) RETURNING *
)
SELECT 
  inserted_feed_follow.*,
  feeds.name as feed_name,
  users.name as user_name
FROM inserted_feed_follow
INNER JOIN users
  ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds
  ON inserted_feed_follow.feed_id = feeds.id;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows ff
  USING feeds f, users u
WHERE ff.feed_id = (SELECT id FROM feeds WHERE feeds.url = $1)
  AND ff.user_id = (SELECT id FROM users WHERE users.name = $2);

-- name: GetFollowedFeeds :many
SELECT feeds.name as feed_name, feeds.url as feed_url, users.name as user_name
FROM feed_follows
INNER JOIN users
  ON user_id = users.id
INNER JOIN feeds
  ON feed_id = feeds.id
WHERE users.name = $1;
