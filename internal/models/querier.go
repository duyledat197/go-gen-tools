// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package models

import (
	"context"
)

type Querier interface {
	GetHub(ctx context.Context, id int64) (Hub, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetUserHub(ctx context.Context) ([]GetUserHubRow, error)
}

var _ Querier = (*Queries)(nil)