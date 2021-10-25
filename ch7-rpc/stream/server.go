package main

import (
	"flag"
	"log"
	pb "micro-go-book/ch7-rpc/stream/stream-pb"
	string_service "micro-go-book/ch7-rpc/stream/string-service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	stringService := new(string_service.StringService)
	pb.RegisterStringServiceServer(grpcServer, stringService)
	grpcServer.Serve(lis)
}
