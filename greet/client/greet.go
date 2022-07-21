package main

import (
	"context"
	"log"

	pb "github.com/fatihkc/go-practice/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet function was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Fatih",
	})

	if err != nil {
		log.Fatalf("Could not greet %v\n", err)
	}

	log.Printf("Greeting: %v\n", res.Result)
}
