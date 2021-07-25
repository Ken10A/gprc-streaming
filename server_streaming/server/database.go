package main

import (
	"log"
	pb "server_streaming/proto"
)

type DatabaseService struct{}

func (db *DatabaseService) Search(r *pb.SearchRequest, s pb.Database_SearchServer) error {
	responses := []string{
		"Highest ranked content",
		"Some ranked content",
		"Some ranked content",
		"Lowest ranked content",
	}

	for idx, resp := range responses {
		result := &pb.SearchResponse{
			MatchedTerm: r.Term,
			Rank:        int32(idx + 1),
			Content:     resp,
		}

		err := s.Send(result)
		if err != nil {
			log.Printf("Error sending message to the client: %v", err)
			return err
		}
	}

	return nil
}
