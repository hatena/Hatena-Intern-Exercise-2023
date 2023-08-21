package main

import (
	"context"
	"fmt"
	"grpc_adventure/pb"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type welcomeServer struct {
	pb.UnimplementedWelcomeServer
}

func (s *welcomeServer) Greet(
	ctx context.Context,
	req *pb.GreetRequest,
) (*pb.GreetReply, error) {
	return &pb.GreetReply{
		Message: fmt.Sprintf("Welcome %s", req.Name),
	}, nil
}

func newServer() *welcomeServer {
	return &welcomeServer{}
}

func main() {
	port := 10000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterWelcomeServer(grpcServer, newServer())
	reflection.Register(grpcServer)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		grpcServer.Serve(lis)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Printf("stopping gRPC server...")
	grpcServer.GracefulStop()
}
