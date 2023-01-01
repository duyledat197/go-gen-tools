-- name: GetHub :one
SELECT * FROM hubs
WHERE id = $1 LIMIT 1;

-- name: GetUserHub :many
Select * from hubs join users using(id);