package deliveries

import (
	"context"
	"fmt"

	"github.com/duyledat197/go-gen-tools/idl/pb"
	"github.com/duyledat197/go-gen-tools/internal/services"
	"github.com/duyledat197/go-gen-tools/transform"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type searchDelivery struct {
	searchService services.SearchService
	pb.UnimplementedSearchServiceServer
}

func NewSearchDelivery(searchService services.SearchService) pb.SearchServiceServer {
	return &searchDelivery{
		searchService: searchService,
	}
}

func (d *searchDelivery) GetTeamHub(ctx context.Context, req *pb.SearchTeamHubRequest) (*pb.SearchTeamHubResponse, error) {
	teams, hubs, err := d.searchService.TeamHub(ctx, req.GetQ())
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Errorf("d.searchService.TeamHub: %v", err).Error())
	}

	return &pb.SearchTeamHubResponse{
		Teams: transform.TeamToPbPtrList(teams),
		Hubs:  transform.HubToPbPtrList(hubs),
	}, nil
}
