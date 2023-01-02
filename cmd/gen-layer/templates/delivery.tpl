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

type {{.PascalCase}}Delivery struct {
	{{.CamelCase}}Service services.{{.PascalCase}}Service
	pb.Unimplemented{{.PascalCase}}ServiceServer
}

func New{{.PascalCase}}Delivery({{.CamelCase}}Service services.{{.PascalCase}}Service) pb.{{.PascalCase}}ServiceServer {
	return &{{.CamelCase}}Delivery{
		{{.CamelCase}}Service: {{.CamelCase}}Service,
	}
}

func (d *{{.CamelCase}}Delivery) Get{{.PascalCase}}ByID(ctx context.Context, req *pb.Get{{.PascalCase}}ByIDRequest) (*pb.{{.PascalCase}}, error) {
	id, err := primitive.ObjectIDFromHex(req.{{.PascalCase}}ID)
	if err != nil {
		return nil, err
	}
	u, err := d.{{.CamelCase}}Service.Get{{.PascalCase}}ByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return to{{.PascalCase}}Pb(u), nil
}

func (d *{{.CamelCase}}Delivery) Create{{.PascalCase}}(ctx context.Context, req *pb.Create{{.PascalCase}}Request) (*pb.Create{{.PascalCase}}Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("validate failed: %w", err).Error())
	}
	hashedPassword, err := models.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	if err := d.{{.CamelCase}}Service.Create{{.PascalCase}}(ctx, &models.{{.PascalCase}}{
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
	return &pb.Create{{.PascalCase}}Response{Success: true}, nil
}

func (d *{{.CamelCase}}Delivery) GetList(ctx context.Context, req *pb.GetListRequest) (*pb.GetListResponse, error) {

	{{.CamelCase}}s, total, err := d.{{.CamelCase}}Service.GetList{{.PascalCase}}(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	result := make([]*pb.{{.PascalCase}}, 0, len({{.CamelCase}}s))
	for _, {{.CamelCase}} := range {{.CamelCase}}s {
		result = append(result, to{{.PascalCase}}Pb({{.CamelCase}}))
	}
	return &pb.GetListResponse{
		Data:  result,
		Total: int32(total),
	}, nil
}

func (d *{{.CamelCase}}Delivery) Update{{.PascalCase}}(ctx context.Context, req *pb.Update{{.PascalCase}}Request) (*pb.Update{{.PascalCase}}Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update{{.PascalCase}} not implemented")
}

func to{{.PascalCase}}Pb(u *models.{{.PascalCase}}) *pb.{{.PascalCase}} {
	deletedAt := ""
	if !u.DeletedAt.IsZero() {
		deletedAt = u.DeletedAt.Format(time.RFC3339)
	}
	return &pb.{{.PascalCase}}{
		{{.PascalCase}}Id:    u.ID.Hex(),
		Role:      pb.{{.PascalCase}}Role(pb.{{.PascalCase}}Role_value[u.Role]),
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
