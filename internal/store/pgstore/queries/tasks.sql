-- name: CreateTask :one
INSERT INTO tasks (
    title, description, priority
) VALUES ($1, $2, $3)
RETURNING *;


-- name: GetTaskById :one
SELECT * FROM tasks WHERE id = $1;


-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY created_at DESC;

-- name: UpdateTask :one
UPDATE tasks
SET title = $1, description = $2, priority = $3, updated_at = now()
WHERE id = $4
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;


