-- name: AddNumber :one
INSERT INTO numbers (
    number
) VALUES (
    $1
)
RETURNING *;
-- name: GetNumbers :many
SELECT * FROM numbers;