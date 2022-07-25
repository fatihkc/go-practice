package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/fatihkc/go-practice/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		res += fmt.Sprint("Hello %s\n", req.FirstName)
	}
}
