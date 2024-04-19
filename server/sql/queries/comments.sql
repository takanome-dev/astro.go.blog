-- name: CreateComment :one
INSERT INTO comments (id, body, user_id, post_id) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetAllComments :many
SELECT * FROM comments
ORDER BY created_at DESC;

-- name: GetCommentByID :one
SELECT * FROM comments WHERE id = $1
ORDER BY created_at DESC;

-- name: UpdateComment :exec
UPDATE comments
SET body = $2, edited_at = NOW()
WHERE id = $1;

-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1;