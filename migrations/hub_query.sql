-- name: GetHub :one
SELECT * FROM hubs
WHERE id = $1 LIMIT 1;