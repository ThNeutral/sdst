-- name: AddUserToProject :exec
INSERT INTO project_users (project_id, user_id, role)
VALUES ($1, $2, $3);

-- name: DeleteUserFromProject :exec
DELETE FROM project_users WHERE project_id = $1 AND user_id = $2;

-- name: GetUserById :one
SELECT * FROM project_users WHERE project_id = $1 AND user_id = $2;