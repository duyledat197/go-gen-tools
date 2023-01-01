package deliveries

import (
	"context"
	"fmt"
	"time"

	"github.com/lalaland/backend/internal/models"
	"github.com/lalaland/backend/internal/services"
	"github.com/lalaland/backend/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userDelivery struct {
	userService services.UserService
	pb.UnimplementedUserServiceServer
}

// NewUserDelivery ...
func NewUserDelivery(userService services.UserService) pb.UserServiceServer {
	return &userDelivery{
		userService: userService,
	}
}

func (d *userDelivery) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.User, error) {
	id, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return nil, err
	}
	u, err := d.userService.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toUserPb(u), nil
}

func (d *userDelivery) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("validate failed: %w", err).Error())
	}
	hashedPassword, err := models.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	if err := d.userService.CreateUser(ctx, &models.User{
		Role:           req.Role.String(),
		Name:           req.Name,
		Phone:          req.Phone,
		Age:            int(req.Age),
		CardID:         req.CardId,
		Gender:         req.Gender,
		BirthDay:       req.BirthDay,
		HashedPassword: hashedPassword,
		Email:          req.Email,
	}); err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Success: true}, nil
}

func (d *userDelivery) GetList(ctx context.Context, req *pb.GetListRequest) (*pb.GetListResponse, error) {

	users, total, err := d.userService.GetListUser(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	result := make([]*pb.User, 0, len(users))
	for _, user := range users {
		result = append(result, toUserPb(user))
	}
	return &pb.GetListResponse{
		Data:  result,
		Total: int32(total),
	}, nil
}

func (d *userDelivery) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func toUserPb(u *models.User) *pb.User {
	deletedAt := ""
	if !u.DeletedAt.IsZero() {
		deletedAt = u.DeletedAt.Format(time.RFC3339)
	}
	return &pb.User{
		UserId:    u.ID.Hex(),
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
		DeletedAt: deletedAt,
	}
}
