-- name: CreatePost :one
INSERT INTO posts (id, title, body, user_id, is_published, is_draft) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetAllPosts :many
SELECT * FROM posts;

-- name: GetPostByID :one
SELECT * FROM posts WHERE id = $1;