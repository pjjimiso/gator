-- name: CreateFeed :one
INSERT INTO feeds (
  id,
  created_at,
  updated_at,
  name,
  url,
  user_id
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name as name, feeds.url as url, users.name as username
FROM feeds
INNER JOIN users ON feeds.user_id = users.id;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET
  last_fetched_at = NOW(),
  updated_at      = NOW()
WHERE id = $1;

