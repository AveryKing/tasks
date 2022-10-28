-- name: CreateTask :one
INSERT INTO tasks ("user", title, priority)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetTasks :many
SELECT * FROM tasks WHERE "user" = $1;