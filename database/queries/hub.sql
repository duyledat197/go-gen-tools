-- name: FindHubByID :one
SELECT * FROM hubs
WHERE id = $1 LIMIT 1;

-- name: CreateHub :one
INSERT INTO hubs(id,name,location_id)
values($1,$2,$3)
Returning *;

-- name: UpdateHub :one
UPDATE hubs set name = $1
where id = $2
Returning *;

-- name: DeleteHub :one
DELETE from hubs where id = $1
Returning *;

-- name: GetListHub :many
SELECT * FROM hubs
offset $1 limit $2;

-- name: SearchHub :many 
SELECT * from hubs
where name like ('%' || $1 || '%');