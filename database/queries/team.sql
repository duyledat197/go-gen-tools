-- name: FindTeamByID :one
SELECT * FROM teams
WHERE id = $1 LIMIT 1;

-- name: CreateTeam :one
INSERT INTO teams(id,name)
values($1,$2)
Returning *;

-- name: UpdateTeam :one
UPDATE teams set name = $1
where id = $2
Returning *;

-- name: DeleteTeam :one
DELETE from teams where id = $1
Returning *;

-- name: GetListTeam :many
SELECT * FROM teams
offset $1 limit $2;