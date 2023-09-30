-- name: CreatePost :one
INSERT INTO posts (id, title, body, image, user_id, is_published, is_draft) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: GetAllPosts :many
SELECT * FROM posts;

-- name: GetPostByID :one
SELECT * FROM posts WHERE id = $1;

-- name: GetPostsByUserID :many
SELECT * FROM posts WHERE user_id = $1;

-- name: UpdatePost :one
UPDATE posts 
SET
  title = COALESCE(sqlc.narg('title'), title),
  body = COALESCE(sqlc.narg('body'), body),
  image = COALESCE(sqlc.narg('image'), image),
  is_published = COALESCE(sqlc.narg('is_published'), is_published),
  is_draft = COALESCE(sqlc.narg('is_draft'), is_draft)
WHERE 
  id = sqlc.arg('id')
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;