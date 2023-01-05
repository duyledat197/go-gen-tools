package services

import (
	"context"
	"testing"

	"github.com/duyledat197/go-gen-tools/internal/models"
	"github.com/duyledat197/go-gen-tools/internal/repositories"
	"github.com/duyledat197/go-gen-tools/mocks"
	"github.com/jackc/pgtype"
	"github.com/stretchr/testify/mock"
)

func Test_hubService_Create(t *testing.T) {
	type fields struct {
		hubRepo repositories.HubRepository
	}
	type args struct {
		ctx context.Context
		hub *models.Hub
	}

	mockHubRepo := mocks.HubRepository{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setup   func(ctx context.Context)
		wantErr bool
	}{
		{
			name: "happy case",
			fields: fields{
				hubRepo: &mockHubRepo,
			},
			args: args{
				ctx: context.Background(),
				hub: &models.Hub{
					ID: pgtype.Text{
						String: "id",
					},
					Name: pgtype.Text{
						String: "name",
					},
				},
			},
			setup: func(ctx context.Context) {
				mockHubRepo.On("Create", mock.Anything, mock.Anything).Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &hubService{
				hubRepo: tt.fields.hubRepo,
			}
			tt.setup(tt.args.ctx)
			if err := s.Create(tt.args.ctx, tt.args.hub); (err != nil) != tt.wantErr {
				t.Errorf("hubService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
