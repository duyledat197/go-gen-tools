package services

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
	"github.com/duyledat197/go-gen-tools/internal/repositories"
)

type SearchService interface {
	TeamHub(ctx context.Context, q string) ([]*models.Team, []*models.Hub, error)
}

type searchService struct {
	searchRepo repositories.SearchRepository
}

func NewSearchService(searchRepo repositories.SearchRepository) SearchService {
	return &searchService{
		searchRepo: searchRepo,
	}
}

func (s *searchService) TeamHub(ctx context.Context, q string) ([]*models.Team, []*models.Hub, error) {
	teams, users, err := s.searchRepo.TeamHub(ctx, q)
	if err != nil {
		return nil, nil, err
	}

	return teams, users, nil
}
