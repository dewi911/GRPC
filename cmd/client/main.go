package main

import (
	"GRPC/proto/notification"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := notification.NewNotificationServiceClient(conn)

	response, err := c.Notify(context.Background(), &notification.NotificationRequest{
		Message: "Hello World",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println("Status:", response.Status)

}
