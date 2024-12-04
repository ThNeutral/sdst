-- name: CreateMessage :one

INSERT INTO messages (id, body, posted_at, user_id, project_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetMessagesByProject :many

SELECT * FROM messages WHERE project_id = $1;

-- name: GetMessagesByUsers :many

SELECT * FROM messages WHERE user_id = $1;
