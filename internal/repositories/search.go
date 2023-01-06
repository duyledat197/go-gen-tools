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
	db models.DBTX
}

func NewSearchRepository(db models.DBTX) SearchRepository {
	return &searchRepository{
		db,
	}
}

func (u *searchRepository) TeamHub(ctx context.Context, q string) ([]*models.Team, []*models.Hub, error) {
	qu := models.New(u.db)
	wg := &sync.WaitGroup{}
	var errSearchTeamChan, errSearchHubChan = make(chan error, 1), make(chan error, 1)
	teams := make([]*models.Team, 0)
	hubs := make([]*models.Hub, 0)
	wg.Add(2)
	go func() {
		defer wg.Done()
		var err error
		teams, err = qu.SearchTeam(ctx, sql.NullString{
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
		hubs, err = qu.SearchHub(ctx, sql.NullString{
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
