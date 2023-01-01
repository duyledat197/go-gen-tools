// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: hub.sql

package models

import (
	"context"
	"database/sql"

	"github.com/jackc/pgtype"
)

const getHub = `-- name: GetHub :one
SELECT id, name, bio FROM hubs
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetHub(ctx context.Context, id int64) (Hub, error) {
	row := q.db.QueryRow(ctx, getHub, id)
	var i Hub
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const getUserHub = `-- name: GetUserHub :many
Select hubs.id, hubs.name, hubs.bio, users.id, users.name, users.bio from hubs join users using(id)
`

type GetUserHubRow struct {
	ID     int64          `db:"id" json:"id"`
	Name   pgtype.Text    `db:"name" json:"name"`
	Bio    sql.NullString `db:"bio" json:"bio"`
	ID_2   int64          `db:"id_2" json:"id_2"`
	Name_2 pgtype.Text    `db:"name_2" json:"name_2"`
	Bio_2  sql.NullString `db:"bio_2" json:"bio_2"`
}

func (q *Queries) GetUserHub(ctx context.Context) ([]GetUserHubRow, error) {
	rows, err := q.db.Query(ctx, getUserHub)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserHubRow
	for rows.Next() {
		var i GetUserHubRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.ID_2,
			&i.Name_2,
			&i.Bio_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
