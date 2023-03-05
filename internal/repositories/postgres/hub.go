package postgres

import (
	"context"
	"encoding/json"

	"github.com/duyledat197/go-gen-tools/internal/models"
	"github.com/duyledat197/go-gen-tools/internal/repositories"
)

type hubRepository struct {
	db models.DBTX
}

func NewHubRepository(db models.DBTX) repositories.HubRepository {
	return &hubRepository{
		db: db,
	}
}

func (r *hubRepository) Create(ctx context.Context, hub *models.Hub, opts ...repositories.Options) error {
	q := models.New(r.db)
	if len(opts) > 0 && opts[0].Tx != nil {
		q = q.WithTx(opts[0].Tx)
	}
	b, err := json.Marshal(hub)
	if err != nil {
		return err
	}
	var params models.CreateHubParams
	if err := json.Unmarshal(b, &params); err != nil {
		return err
	}
	if _, err := q.CreateHub(ctx, params); err != nil {
		return err
	}

	return nil
}

func (r *hubRepository) Update(ctx context.Context, filter, hub *models.Hub, opts ...repositories.Options) error {
	q := models.New(r.db)
	if len(opts) > 0 && opts[0].Tx != nil {
		q = q.WithTx(opts[0].Tx)
	}

	q.UpdateHub(ctx, models.UpdateHubParams{})

	return nil
}

func (r *hubRepository) Delete(ctx context.Context, id string, opts ...repositories.Options) error {
	q := models.New(r.db)
	if len(opts) > 0 && opts[0].Tx != nil {
		q = q.WithTx(opts[0].Tx)
	}
	return nil
}

func (r *hubRepository) GetList(ctx context.Context, offset, limit int, opts ...repositories.Options) ([]*models.Hub, error) {
	q := models.New(r.db)
	if len(opts) > 0 && opts[0].Tx != nil {
		q = q.WithTx(opts[0].Tx)
	}
	return nil, nil
}

func (r *hubRepository) GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.Hub, error) {
	q := models.New(r.db)
	if len(opts) > 0 && opts[0].Tx != nil {
		q = q.WithTx(opts[0].Tx)
	}
	return nil, nil
}
