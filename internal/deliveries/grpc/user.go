
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

type userDelivery struct {
	userService services.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserDelivery(userService services.UserService) pb.UserServiceServer {
	return &userDelivery{
		userService: userService,
	}
}

func (d *userDelivery) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := d.userService.Create(ctx, transform.PbToUserPtr(req.GetUser())); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Errorf("Create: %v", err).Error())
	}
	return &pb.CreateUserResponse{}, nil
}

func (d *userDelivery) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if err := d.userService.Update(ctx, req.User.Id, transform.PbToUserPtr(req.GetUser())); err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{
		Success: true,
	}, nil
}

func (d *userDelivery) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	if err := d.userService.Delete(ctx, req.User.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}

func (d *userDelivery) GetList(ctx context.Context, req *pb.GetListUserRequest) (*pb.GetListUserResponse, error) {
	users, err := d.userService.GetList(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	return &pb.GetListUserResponse{
		Data: transform.UserToPbPtrList(users),
	}, nil
}

func (d *userDelivery) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	user, err := d.userService.GetByID(ctx, req.GetUserID())
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByIDResponse{
		Data: transform.UserToPbPtr(user),
	}, nil
}
