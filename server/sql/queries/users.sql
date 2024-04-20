-- name: CreateUser :one
INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users 
WHERE id = $1;

-- name: GetUserKPIs :one
SELECT sqlc.embed(users), 
       COALESCE(
           (SELECT json_agg(posts_with_comment_count)
           FROM (
                SELECT posts.*, COUNT(comments.*) AS comments_count
                FROM posts
                LEFT JOIN comments ON comments.post_id = posts.id
                WHERE posts.user_id = users.id
                GROUP BY posts.id
                ORDER BY posts.created_at DESC
                LIMIT 3
           ) AS posts_with_comment_count
           ), '[]'::json
       ) AS last_three_posts,
       COALESCE(
           (
            SELECT json_agg(comments)
           FROM (
               SELECT * FROM comments
               WHERE comments.user_id = users.id
               ORDER BY comments.created_at DESC
               LIMIT 3
           ) AS comments
           ), '[]'::json
       ) AS last_three_comments
FROM users
WHERE users.id = $1;


-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: UpdateUserPassword :one
UPDATE users SET password = $1 WHERE id = $2 RETURNING *;
