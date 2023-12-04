-- name: CreateUser :one
INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users 
WHERE id = $1;

-- name: GetUserKPIs :one
SELECT sqlc.embed(users), 
       (
           SELECT json_agg(posts)
           FROM (
               SELECT * FROM posts
               WHERE posts.user_id = users.id
               ORDER BY posts.created_at DESC
               LIMIT 3
           ) AS posts
       ) AS last_three_posts,
       (
           SELECT json_agg(comments)
           FROM (
               SELECT * FROM comments
               WHERE comments.user_id = users.id
               ORDER BY comments.created_at DESC
               LIMIT 3
           ) AS comments
       ) AS last_three_comments
FROM users
WHERE users.id = $1;


-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: UpdateUserPassword :one
UPDATE users SET password = $1 WHERE id = $2 RETURNING *;
