package main

import (
	"context"
	"log"
	"time"

	pb "github.com/fatihkc/go-practice/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Fatih"},
		{FirstName: "Kemal"},
		{FirstName: "Cemal"},
	}
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.Recv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
