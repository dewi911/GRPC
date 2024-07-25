package main

import (
	"GRPC/proto/notification"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	notification.UnimplementedNotificationServiceServer
}

func (s *server) Notify(ctx context.Context, n *notification.NotificationRequest) (*notification.NotificationResponse, error) {
	fmt.Println("RECEIVED NITIFICATION", n.Message)
	return &notification.NotificationResponse{
		Status: "OK",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	notification.RegisterNotificationServiceServer(s, &server{})

	go func() {
		log.Println("Starting gRPC Server on :9000")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	log.Println("Shutting down gRPC Server")
	s.GracefulStop()
	log.Println("Server gracefully stopped")

}
