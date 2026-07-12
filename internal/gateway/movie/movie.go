package movie

import (
	"context"
	"fmt"

	movie1 "github.com/fnuritdinov/proto/moviePr"

	"google.golang.org/grpc"
)

type Gateway struct {
	conn   *grpc.ClientConn
	client movie1.MovieServiceClient
}

func New(address string) (*Gateway, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not connect %w", err)
	}

	return &Gateway{
		conn:   conn,
		client: movie1.NewMovieServiceClient(conn),
	}, nil
}

func (g *Gateway) GetByID(ctx context.Context, id int64) (*movie1.GetMovieResponse, error) {
	return g.client.GetByID(ctx, &movie1.GetMovieRequest{Id: id})
}
