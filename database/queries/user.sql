-- name: FindUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users(id,name)
values($1,$2)
Returning *;

-- name: UpdateUser :one
UPDATE users set name = $1
where id = $2
Returning *;

-- name: DeleteUser :one
DELETE from users where id = $1
Returning *;

-- name: GetListUser :many
SELECT * FROM users
offset $1 limit $2;

-- name: GetFunction :many
select * from test();