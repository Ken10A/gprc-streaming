package main

import (
	pb "bidirectional_streaming/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	// pb.RegisterDatabaseServer(s, new(DatabaseService))
	// pb.RegisterArchiverServer(s, new(ArchiverService))
	pb.RegisterTokenizerServer(s, new(TokenizerService))

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
