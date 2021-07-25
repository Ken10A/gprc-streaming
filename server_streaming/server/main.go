package main

import (
	"log"
	"net"
	pb "server_streaming/proto"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterDatabaseServer(s, new(DatabaseService))

	log.Print("Starting RPC server on port 8080...")
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to setup tcp listener: %v", err)
	}

	err = s.Serve(list)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
