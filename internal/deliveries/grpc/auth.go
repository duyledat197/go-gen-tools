package deliveries

import (
	"context"
	"fmt"
	"time"

	"github.com/lalaland/backend/internal/services"
	"github.com/lalaland/backend/pb"

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
	u, tkn, err := d.authService.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	res := &pb.LoginResponse{
		UserId:    u.ID.Hex(),
		Token:     tkn,
		Role:      pb.UserRole(pb.UserRole_value[u.Role]),
		Name:      u.Name,
		Phone:     u.Phone,
		Age:       int32(u.Age),
		CardId:    u.CardID,
		Gender:    u.Gender,
		BirthDay:  u.BirthDay,
		Email:     u.Email,
		Address:   u.Address,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
	}
	return res, nil
}

// func (d *authDelivery) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
// 	err := d.authservices.ChangePassword(ctx, req.Phone, req.Password, req.NewPassword)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.ChangePasswordResponse{}, nil
// }
