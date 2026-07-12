package user

import (
	"context"
	"fmt"

	user1 "github.com/fnuritdinov/proto/userPr/v1"
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

func (ug *UserGateway) GetByID(ctx context.Context, id int64) (*user1.GetUserResponse, error) {
	return ug.client.GetByID(ctx, &user1.GetUserRequest{Id: id})
}
