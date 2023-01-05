package deliveries

import (
	"context"
	"fmt"

	"github.com/duyledat197/go-gen-tools/internal/services"
	"github.com/duyledat197/go-gen-tools/pb"
	"github.com/duyledat197/go-gen-tools/transform"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type hubDelivery struct {
	hubService services.HubService
	pb.UnimplementedHubServiceServer
}

func NewHubDelivery(hubService services.HubService) pb.HubServiceServer {
	return &hubDelivery{
		hubService: hubService,
	}
}

func (d *hubDelivery) CreateHub(ctx context.Context, req *pb.CreateHubRequest) (*pb.CreateHubResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("validate failed: %w", err).Error())
	}
	if err := d.hubService.Create(ctx, transform.PbToHubPtr(req.GetHub())); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Errorf("Create: %v", err).Error())
	}
	return &pb.CreateHubResponse{
		Success: true,
	}, nil
}

func (d *hubDelivery) UpdateHub(ctx context.Context, req *pb.UpdateHubRequest) (*pb.UpdateHubResponse, error) {
	if err := d.hubService.Update(ctx, req.Hub.Id, transform.PbToHubPtr(req.GetHub())); err != nil {
		return nil, err
	}
	return &pb.UpdateHubResponse{
		Success: true,
	}, nil
}

func (d *hubDelivery) GetList(ctx context.Context, req *pb.GetListHubRequest) (*pb.GetListHubResponse, error) {
	hubs, err := d.hubService.GetList(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	return &pb.GetListHubResponse{
		Data: transform.HubToPbPtrList(hubs),
	}, nil
}

func (d *hubDelivery) GetHubByID(ctx context.Context, req *pb.GetHubByIDRequest) (*pb.GetHubByIDResponse, error) {
	hub, err := d.hubService.GetByID(ctx, req.GetHubID())
	if err != nil {
		return nil, err
	}
	return &pb.GetHubByIDResponse{
		Data: transform.HubToPbPtr(hub),
	}, nil
}
