package main

import (
	"io"
	"log"

	pb "github.com/fatihkc/go-practice/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		res := "Hello " + req.GetFirstName() + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending response: %v\n", err)
		}
	}
}
