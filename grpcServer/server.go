package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/renix-codex/gRPC/grpcServer/protobufs"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedOrgServiceServer
	orgs map[string]*pb.OrganisationResponse
}

func (s *server) GetOrganisationByID(ctx context.Context, req *pb.OrganisationIDRequest) (*pb.OrganisationResponse, error) {
	org, exists := s.orgs[req.GetId()]
	if !exists {
		return nil, fmt.Errorf("Org with ID %s not found", req.GetId())
	}
	return org, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := &server{
		orgs: map[string]*pb.OrganisationResponse{
			"1": {Id: "1", Name: "Organisation One", Meta: &pb.Meta{Address: "Bengaluru", Pin: "560075"}},
			"2": {Id: "2", Name: "Organisation Two", Meta: &pb.Meta{Address: "Kochi", Pin: "689121"}},
		},
	}
	pb.RegisterOrgServiceServer(s, srv)
	fmt.Println("Server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
