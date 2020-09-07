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
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/rmbarron/SnackInventory/src/backend/server/connector"
	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	portFlag        = flag.Int("port", 10000, "Port for SnackInventory to listen on.")
	storageImplFlag = flag.String(
		"storage_architecture", "mysql",
		`Architecture to use for backing storage. Valid values include:
		 - mysql`)

	// Flags for connector.SQLImpl.
	sqlUserFlag = flag.String("sql_user", "", "Username for connecting to MySQL.")
	sqlAddrFlag = flag.String(
		"sql_address", "", "host:port address for connecting to MySQL.")
	sqlDBNameFlag = flag.String(
		"sql_database", "SnackInventory", "MySQL database name to connect to.")
)

// Interface for connecting to backing storage.
type dbConnector interface {
	ListSnacks(ctx context.Context) ([]*sipb.Snack, error)
}

type snackInventoryServer struct {
	c dbConnector
}

func (s *snackInventoryServer) CreateSnack(ctx context.Context, req *sipb.CreateSnackRequest) (*sipb.CreateSnackResponse, error) {
	return &sipb.CreateSnackResponse{}, nil
}

func (s *snackInventoryServer) ListSnacks(ctx context.Context, req *sipb.ListSnacksRequest) (*sipb.ListSnacksResponse, error) {
	snacks, err := s.c.ListSnacks(ctx)
	if err != nil {
		// TODO: Translate storage errors to corresponding canonical code.
		return nil, status.Errorf(codes.Internal, "could not list snacks: %v", err)
	}
	return &sipb.ListSnacksResponse{
		Snacks: snacks,
	}, nil
}

func (s *snackInventoryServer) DeleteSnack(ctx context.Context, req *sipb.DeleteSnackRequest) (*sipb.DeleteSnackResponse, error) {
	return &sipb.DeleteSnackResponse{}, nil
}

func main() {
	flag.Parse()

	var c dbConnector
	var err error
	switch si := *storageImplFlag; si {
	case "mysql":
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		pwd := scanner.Text()

		if *sqlUserFlag == "" {
			log.Fatal("--sql_user is required.")
		}

		if *sqlAddrFlag == "" {
			log.Fatal("--sql_address is required.")
		}

		if *sqlDBNameFlag == "" {
			log.Fatal("--sql_database is required.")
		}

		c, err = connector.NewSQLImpl(context.Background(), *sqlUserFlag, pwd, *sqlAddrFlag, *sqlDBNameFlag)
		if err != nil {
			log.Fatalf("could not connect to SQL: %v", err)
		}
	default:
		log.Fatal("unsupported storage implementation requested.")
	}

	si := &snackInventoryServer{
		c: c,
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *portFlag))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// TODO: Serve with TLS.
	grpcServer := grpc.NewServer()
	svc := sipb.NewSnackInventoryService(si)
	sipb.RegisterSnackInventoryService(grpcServer, svc)
	grpcServer.Serve(lis)
}
