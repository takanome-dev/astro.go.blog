-- name: CreatePost :one
INSERT INTO posts (id, title, body, user_id, is_published, is_draft) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetAllPosts :many
SELECT * FROM posts;

-- name: GetPostByID :one
SELECT * FROM posts WHERE id = $1;

-- name: UpdatePost :one
UPDATE posts 
SET
  title = COALESCE(sqlc.narg('title'), title),
  body = COALESCE(sqlc.narg('body'), body),
  is_published = COALESCE(sqlc.narg('is_published'), is_published),
  is_draft = COALESCE(sqlc.narg('is_draft'), is_draft)
WHERE 
  id = sql.arg('id')
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;