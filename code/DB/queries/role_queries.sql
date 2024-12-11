-- name: UpdateUserRole :exec
UPDATE project_users SET role = $1 WHERE project_id = $2 AND user_id = $3;