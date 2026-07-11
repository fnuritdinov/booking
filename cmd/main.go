package main

import (
	"booking-service/booking"
	"booking-service/internal/config"
	"booking-service/internal/gateway/movie"
	"booking-service/internal/gateway/user"
	"booking-service/internal/repository"
	"booking-service/internal/server"
	"booking-service/internal/service"
	"booking-service/pkg/db"
	"booking-service/pkg/logger"
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	cfg, err := config.New("/config/config.env")
	if err != nil {
		log.Fatal("config.New", err)
	}

	conn, err := db.New(db.Option{
		Host:     cfg.DBHost,
		Port:     cfg.Port,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.BDName,
	})
	if err != nil {
		log.Fatal("failed to connect to db: %w", err)
	}

	lg, err := logger.New(true)
	if err != nil {
		log.Fatal("failed to create logger: $w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis, err := net.Listen(cfg.NETWORK, cfg.ADDRESS)
	if err != nil {
		log.Fatal("failed to listen: %w", err)
	}

	movieGateway, err := movie.New(cfg.ADDRESS)
	if err != nil {
		log.Fatal("failed to connect: %w", err)
	}

	userGateway, err := user.New(cfg.ADDRESS)
	if err != nil {
		log.Fatal("failed to connect: %w", err)
	}

	grpcServer := grpc.NewServer()

	bookingRepo := repository.New(conn)

	bookingService := service.New(bookingRepo, *movieGateway, *userGateway)

	bookingServer := server.New(bookingService, *lg)

	booking.RegisterBookingServiceServer(grpcServer, bookingServer)

	reflection.Register(grpcServer)

	go func() {
		lg.Info("server listening at %v", zap.String("addr", lis.Addr().String()))
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal("failed to serve: %w", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	lg.Info("shutting down server...")
	grpcServer.GracefulStop()
	lg.Info("server stopped")
}
