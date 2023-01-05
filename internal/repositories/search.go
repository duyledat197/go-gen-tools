package repositories

import (
	"context"
	"database/sql"
	"sync"

	"github.com/duyledat197/go-gen-tools/internal/models"
)

type SearchRepository interface {
	TeamHub(ctx context.Context, q string) ([]*models.Team, []*models.Hub, error)
}

type searchRepository struct {
	db *models.Queries
}

func NewSearchRepository(q *models.Queries) SearchRepository {
	return &searchRepository{
		db: q,
	}
}

func (u *searchRepository) TeamHub(ctx context.Context, q string) ([]*models.Team, []*models.Hub, error) {
	wg := &sync.WaitGroup{}
	var errSearchTeamChan, errSearchHubChan = make(chan error, 1), make(chan error, 1)
	teams := make([]*models.Team, 0)
	hubs := make([]*models.Hub, 0)
	wg.Add(2)
	go func() {
		defer wg.Done()
		var err error
		teams, err = u.db.SearchTeam(ctx, sql.NullString{
			String: q,
			Valid:  true,
		})
		if err != nil {
			errSearchTeamChan <- err
		}

		errSearchTeamChan <- nil
	}()

	go func() {
		defer wg.Done()
		var err error
		hubs, err = u.db.SearchHub(ctx, sql.NullString{
			String: q,
			Valid:  true,
		})

		if err != nil {
			errSearchHubChan <- err
		}

		errSearchHubChan <- nil
	}()
	wg.Wait()
	if <-errSearchTeamChan != nil {
		return nil, nil, <-errSearchTeamChan
	}

	if <-errSearchHubChan != nil {
		return nil, nil, <-errSearchHubChan
	}

	return teams, hubs, nil
}
