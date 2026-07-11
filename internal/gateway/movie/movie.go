package movie

import (
	movie1 "booking-service/movie"
	"context"
	"fmt"

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

func (g *Gateway) Create(ctx context.Context, request *movie1.CreateMovieRequest) (*movie1.CreateMovieResponse, error) {
	return g.client.Create(ctx, request)
}

func (g *Gateway) GetByID(ctx context.Context, request *movie1.GetMovieRequest) (*movie1.GetMovieResponse, error) {
	return g.client.GetByID(ctx, request)
}

func (g *Gateway) List(ctx context.Context, request *movie1.ListMovieRequest) (*movie1.ListMovieResponse, error) {
	return g.client.List(ctx, request)
}

func (g *Gateway) Update(ctx context.Context, request *movie1.UpdateMovieRequest) (*movie1.UpdateMovieResponse, error) {
	return g.client.Update(ctx, request)
}

func (g *Gateway) Delete(ctx context.Context, request *movie1.DeleteMovieRequest) (*movie1.DeleteMovieResponse, error) {
	return g.client.Delete(ctx, request)
}
