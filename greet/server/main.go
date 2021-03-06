package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	greetPb "github.com/ekkinox/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type server struct {
	greetPb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetPb.GreetRequest) (*greetPb.GreetResponse, error) {

	title := req.GetGreeting().GetTitle()
	name := req.GetGreeting().GetName()

	fmt.Printf("Greet: received title %s and name %s\n", title, name)

	return &greetPb.GreetResponse{
		Result: fmt.Sprintf("Greetings %s %s!", title, name),
	}, nil
}

func (*server) GreetManyTimes(req *greetPb.GreetManyTimesRequest, stream greetPb.GreetService_GreetManyTimesServer) error {

	title := req.GetGreeting().GetTitle()
	name := req.GetGreeting().GetName()

	fmt.Printf("GreetManyTimes: received title %s and name %s\n", title, name)

	for i := 0; i < 10; i++ {

		resp := &greetPb.GreetManyTimesResponse{
			Result: fmt.Sprintf("Greetings %s %s (#%d)", title, name, i),
		}

		stream.Send(resp)

		time.Sleep(1 * time.Second)
	}

	return nil
}

func (*server) LongGreet(stream greetPb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet: received client stream")

	result := ""
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&greetPb.LongGreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Failed to receive: %v", err)
		}

		title := req.GetGreeting().GetTitle()
		name := req.GetGreeting().GetName()

		result += fmt.Sprintf("Greetings %s %s\n", title, name)
	}
}

func (*server) GreetAll(stream greetPb.GreetService_GreetAllServer) error {

	fmt.Printf("GreetAll: received RPC ... ")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			fmt.Printf("client is done sending\n")
			return nil
		}

		if err != nil {
			return err
		}

		err = stream.Send(&greetPb.GreetAllResponse{
			Result: fmt.Sprintf("Greetings %s %s !\n", req.Greeting.Title, req.Greeting.Name),
		})

		if err != nil {
			return err
		}
	}
}

func main() {
	fmt.Println("Starting gRPC server on :50051 ...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	tls := true
	if tls {
		creds, err := credentials.NewServerTLSFromFile("../../ssl/server.crt", "../../ssl/server.pem")
		if err != nil {
			log.Fatalf("error loading certs: %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	greetPb.RegisterGreetServiceServer(s, &server{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
