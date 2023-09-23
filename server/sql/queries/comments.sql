-- name: CreateComment :one
INSERT INTO comments (id, body, user_id, post_id) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetAllComments :many
SELECT * FROM comments;

-- name: GetCommentByID :one
SELECT * FROM comments WHERE id = $1;

-- name: UpdateComment :one
UPDATE comments
SET body = COALESCE(sqlc.narg('body'), body)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1;