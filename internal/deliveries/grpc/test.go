package deliveries

import (
	"context"
	"fmt"
	"time"

	"github.com/duyledat197/interview-hao/internal/models"
	"github.com/duyledat197/interview-hao/internal/services"
	"github.com/duyledat197/interview-hao/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TestDelivery struct {
	testService services.TestService
	pb.UnimplementedTestServiceServer
}

func NewTestDelivery(testService services.TestService) pb.TestServiceServer {
	return &testDelivery{
		testService: testService,
	}
}

func (d *testDelivery) GetTestByID(ctx context.Context, req *pb.GetTestByIDRequest) (*pb.Test, error) {
	id, err := primitive.ObjectIDFromHex(req.TestID)
	if err != nil {
		return nil, err
	}
	u, err := d.testService.GetTestByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toTestPb(u), nil
}

func (d *testDelivery) CreateTest(ctx context.Context, req *pb.CreateTestRequest) (*pb.CreateTestResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("validate failed: %w", err).Error())
	}
	hashedPassword, err := models.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	if err := d.testService.CreateTest(ctx, &models.Test{
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
	return &pb.CreateTestResponse{Success: true}, nil
}

func (d *testDelivery) GetList(ctx context.Context, req *pb.GetListRequest) (*pb.GetListResponse, error) {

	tests, total, err := d.testService.GetListTest(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	result := make([]*pb.Test, 0, len(tests))
	for _, test := range tests {
		result = append(result, toTestPb(test))
	}
	return &pb.GetListResponse{
		Data:  result,
		Total: int32(total),
	}, nil
}

func (d *testDelivery) UpdateTest(ctx context.Context, req *pb.UpdateTestRequest) (*pb.UpdateTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTest not implemented")
}

func toTestPb(u *models.Test) *pb.Test {
	deletedAt := ""
	if !u.DeletedAt.IsZero() {
		deletedAt = u.DeletedAt.Format(time.RFC3339)
	}
	return &pb.Test{
		TestId:    u.ID.Hex(),
		Role:      pb.TestRole(pb.TestRole_value[u.Role]),
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
