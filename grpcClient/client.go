package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/renix-codex/gRPC/grpcClient/protobufs"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrgServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter Org ID: ")
		id, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("could not read input: %v", err)
		}
		// Remove the newline character
		id = id[:len(id)-1]

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.GetOrgByID(ctx, &pb.OrgIDRequest{Id: id})
		if err != nil {
			log.Printf("could not get org: %v", err)
		} else {
			log.Printf("Org: %s, Name: %s, Address: %s, Pin: %s", r.GetId(), r.GetName(), r.GetMeta().GetAddress(), r.GetMeta().GetPin())
		}
	}
}
