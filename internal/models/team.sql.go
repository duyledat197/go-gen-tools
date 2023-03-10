// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: team.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTeam = `-- name: CreateTeam :one
INSERT INTO teams(id,name,type,location_id,hub_id)
values($1,$2,$3,$4,$5)
Returning id, name, type, hub_id, location_id, created_at, updated_at, deleted_at
`

type CreateTeamParams struct {
	ID         string `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Type       string `db:"type" json:"type"`
	LocationID string `db:"location_id" json:"location_id"`
	HubID      string `db:"hub_id" json:"hub_id"`
}

func (q *Queries) CreateTeam(ctx context.Context, arg CreateTeamParams) (*Team, error) {
	row := q.db.QueryRow(ctx, createTeam,
		arg.ID,
		arg.Name,
		arg.Type,
		arg.LocationID,
		arg.HubID,
	)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.HubID,
		&i.LocationID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const deleteTeam = `-- name: DeleteTeam :one
DELETE from teams where id = $1
Returning id, name, type, hub_id, location_id, created_at, updated_at, deleted_at
`

func (q *Queries) DeleteTeam(ctx context.Context, id string) (*Team, error) {
	row := q.db.QueryRow(ctx, deleteTeam, id)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.HubID,
		&i.LocationID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const findTeamByID = `-- name: FindTeamByID :one
SELECT id, name, type, hub_id, location_id, created_at, updated_at, deleted_at FROM teams
WHERE id = $1 LIMIT 1
`

func (q *Queries) FindTeamByID(ctx context.Context, id string) (*Team, error) {
	row := q.db.QueryRow(ctx, findTeamByID, id)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.HubID,
		&i.LocationID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getListTeam = `-- name: GetListTeam :many
SELECT id, name, type, hub_id, location_id, created_at, updated_at, deleted_at FROM teams
offset $1 limit $2
`

type GetListTeamParams struct {
	Offset int32 `db:"offset" json:"offset"`
	Limit  int32 `db:"limit" json:"limit"`
}

func (q *Queries) GetListTeam(ctx context.Context, arg GetListTeamParams) ([]*Team, error) {
	rows, err := q.db.Query(ctx, getListTeam, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.HubID,
			&i.LocationID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchTeam = `-- name: SearchTeam :many
SELECT id, name, type, hub_id, location_id, created_at, updated_at, deleted_at from teams 
where name LIKE ('%' || $1 || '%') or type like ('%' || $1 || '%')
`

func (q *Queries) SearchTeam(ctx context.Context, dollar_1 pgtype.Text) ([]*Team, error) {
	rows, err := q.db.Query(ctx, searchTeam, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.HubID,
			&i.LocationID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTeam = `-- name: UpdateTeam :one
UPDATE teams set name = $1
where id = $2
Returning id, name, type, hub_id, location_id, created_at, updated_at, deleted_at
`

type UpdateTeamParams struct {
	Name string `db:"name" json:"name"`
	ID   string `db:"id" json:"id"`
}

func (q *Queries) UpdateTeam(ctx context.Context, arg UpdateTeamParams) (*Team, error) {
	row := q.db.QueryRow(ctx, updateTeam, arg.Name, arg.ID)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.HubID,
		&i.LocationID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
