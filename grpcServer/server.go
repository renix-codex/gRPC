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
	orgs map[string]*pb.OrgResponse
}

func (s *server) GetOrgByID(ctx context.Context, req *pb.OrgIDRequest) (*pb.OrgResponse, error) {
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
		orgs: map[string]*pb.OrgResponse{
			"1": {Id: "1", Name: "Org One", Meta: &pb.Meta{Address: "123 Main St", Pin: "12345"}},
			"2": {Id: "2", Name: "Org Two", Meta: &pb.Meta{Address: "456 Elm St", Pin: "67890"}},
		},
	}
	pb.RegisterOrgServiceServer(s, srv)
	fmt.Println("Server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
