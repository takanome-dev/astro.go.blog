-- name: CreatePost :one
INSERT INTO posts (id, title, body, image, user_id, is_published, is_draft) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: GetAllPosts :many
SELECT sqlc.embed(posts), sqlc.embed(users) FROM posts
JOIN users ON posts.user_id = users.id
ORDER BY posts.created_at DESC;

-- name: GetPostByID :one
SELECT sqlc.embed(posts), sqlc.embed(users), 
COALESCE(
  (
    SELECT json_agg(json_build_object('comment', comments, 'user', users))::text
    FROM comments 
    JOIN users ON comments.user_id = users.id
    WHERE posts.id = comments.post_id
  ), 
  NULL
  ) as comments
FROM posts
JOIN users ON posts.user_id = users.id
WHERE posts.id = $1;

-- name: GetPostsByUserID :many
SELECT sqlc.embed(posts), sqlc.embed(users) FROM posts
JOIN users ON posts.user_id = users.id
WHERE user_id = $1;

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