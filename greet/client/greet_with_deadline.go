package main

import (
	"context"
	"log"
	"time"

	pb "github.com/fatihkc/go-practice/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Fatih",
	}
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
				return
			} else {
				log.Printf("Unexpected error: %v\n", e)
			}
			log.Fatalf("A non-grpc error occurred: %v", err)
		}
	}

	log.Printf("GreetWithDeadline result: %s\n", res.Result)
}
