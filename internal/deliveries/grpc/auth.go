package deliveries

import (
	"context"
	"fmt"

	"github.com/duyledat197/interview-hao/internal/services"
	"github.com/duyledat197/interview-hao/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authDelivery struct {
	authService services.AuthService
	pb.UnimplementedAuthServiceServer
}

// NewAuthDelivery ...
func NewAuthDelivery(authService services.AuthService) pb.AuthServiceServer {
	return &authDelivery{
		authService: authService,
	}
}

func (d *authDelivery) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("validate failed: %w", err).Error())
	}
	// u, tkn, err := d.authService.Login(ctx, req.GetEmail(), req.GetPassword())
	// if err != nil {
	// 	return nil, status.Errorf(codes.InvalidArgument, err.Error())
	// }
	res := &pb.LoginResponse{}
	return res, nil
}

// func (d *authDelivery) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
// 	err := d.authservices.ChangePassword(ctx, req.Phone, req.Password, req.NewPassword)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.ChangePasswordResponse{}, nil
// }
