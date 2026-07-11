package user

import (
	user1 "booking-service/userpb/v1"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type UserGateway struct {
	conn   *grpc.ClientConn
	client user1.UserServiceClient
}

func New(address string) (*UserGateway, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not connect %w", err)
	}

	return &UserGateway{
		conn:   conn,
		client: user1.NewUserServiceClient(conn),
	}, nil
}

func (ug *UserGateway) Add(ctx context.Context, req *user1.CreateUserRequest) (*user1.CreateUserResponse, error) {
	return ug.client.Add(ctx, req)
}

func (ug *UserGateway) GetByID(ctx context.Context, req *user1.GetUserRequest) (*user1.GetUserResponse, error) {
	return ug.client.GetByID(ctx, req)
}

func (ug *UserGateway) Update(ctx context.Context, req *user1.UpdateUserRequest) (*user1.UpdateUserResponse, error) {
	return ug.client.Update(ctx, req)
}
