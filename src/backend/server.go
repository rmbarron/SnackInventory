// Package main provides a gRPC service to interact with SnackInventory storage.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
	"google.golang.org/grpc"
)

var portFlag = flag.Int("port", 10000, "Port for SnackInventory to listen on.")

type snackInventoryServer struct{}

func (s *snackInventoryServer) CreateSnack(ctx context.Context, req *sipb.CreateSnackRequest) (*sipb.CreateSnackResposne, error) {
	return nil, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *portFlag))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	sipb.RegisterSnackInventoryService(grpcServer, &snackInventoryServer)
	grpcServer.Serve(lis)
}
