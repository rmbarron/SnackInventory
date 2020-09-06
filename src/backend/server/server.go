/*
Copyright 2020 Robert Barron

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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

func (s *snackInventoryServer) CreateSnack(ctx context.Context, req *sipb.CreateSnackRequest) (*sipb.CreateSnackResponse, error) {
	return &sipb.CreateSnackResponse{}, nil
}

func (s *snackInventoryServer) ListSnacks(ctx context.Context, req *sipb.ListSnacksRequest) (*sipb.ListSnacksResponse, error) {
	return &sipb.ListSnacksResponse{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *portFlag))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// TODO: Serve with TLS.
	grpcServer := grpc.NewServer()
	svc := sipb.NewSnackInventoryService(&snackInventoryServer{})
	sipb.RegisterSnackInventoryService(grpcServer, svc)
	grpcServer.Serve(lis)
}
