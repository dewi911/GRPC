package main

import (
	"GRPC/proto/notification"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "localhost:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := notification.NewNotificationServiceClient(conn)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := c.Notify(context.Background(), &notification.NotificationRequest{
		Message: "Hello World",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println("Status:", response.Status)

}
