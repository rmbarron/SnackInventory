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
	// Snack Registry Operations
	CreateSnack(ctx context.Context, barcode, name string) error
	ListSnacks(ctx context.Context) ([]*sipb.Snack, error)
	UpdateSnack(ctx context.Context, barcode, name string) error
	DeleteSnack(ctx context.Context, barcode string) error

	// Location Registry Operations
	CreateLocation(ctx context.Context, name string) error
	ListLocations(ctx context.Context) ([]*sipb.Location, error)
	DeleteLocation(ctx context.Context, name string) error

	// Location Contents Operations
	AddSnack(ctx context.Context, snackBarcode, locationName string) (createdSnack, createdLocation bool, err error)
}

type snackInventoryServer struct {
	c dbConnector
}

func (s *snackInventoryServer) CreateSnack(ctx context.Context, req *sipb.CreateSnackRequest) (*sipb.CreateSnackResponse, error) {
	if err := s.c.CreateSnack(ctx, req.GetSnack().GetBarcode(), req.GetSnack().GetName()); err != nil {
		if c := status.Code(err); c == codes.AlreadyExists {
			return nil, err
		}
		return nil, status.Errorf(codes.Internal, "could not create snack: %v", err)
	}
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

func (s *snackInventoryServer) UpdateSnack(ctx context.Context, req *sipb.UpdateSnackRequest) (*sipb.UpdateSnackResponse, error) {
	if err := s.c.UpdateSnack(ctx, req.GetSnack().GetBarcode(), req.GetSnack().GetName()); err != nil {
		return nil, status.Errorf(codes.Internal, "could not update snack: %v", err)
	}
	return &sipb.UpdateSnackResponse{}, nil
}

func (s *snackInventoryServer) DeleteSnack(ctx context.Context, req *sipb.DeleteSnackRequest) (*sipb.DeleteSnackResponse, error) {
	if err := s.c.DeleteSnack(ctx, req.GetBarcode()); err != nil {
		return nil, status.Errorf(codes.Internal, "could not delete snack: %v", err)
	}
	return &sipb.DeleteSnackResponse{}, nil
}

func (s *snackInventoryServer) CreateLocation(ctx context.Context, req *sipb.CreateLocationRequest) (*sipb.CreateLocationResponse, error) {
	if err := s.c.CreateLocation(ctx, req.GetLocation().GetName()); err != nil {
		if c := status.Code(err); c == codes.AlreadyExists {
			return nil, err
		}
		return nil, status.Errorf(codes.Internal, "could not create location: %v", err)
	}
	return &sipb.CreateLocationResponse{}, nil
}

func (s *snackInventoryServer) ListLocations(ctx context.Context, req *sipb.ListLocationsRequest) (*sipb.ListLocationsResponse, error) {
	locations, err := s.c.ListLocations(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not list locations: %v", err)
	}
	return &sipb.ListLocationsResponse{Locations: locations}, nil
}

func (s *snackInventoryServer) DeleteLocation(ctx context.Context, req *sipb.DeleteLocationRequest) (*sipb.DeleteLocationResponse, error) {
	if err := s.c.DeleteLocation(ctx, req.GetName()); err != nil {
		return nil, status.Errorf(codes.Internal, "could not delete location: %v", err)
	}
	return &sipb.DeleteLocationResponse{}, nil
}

func (s *snackInventoryServer) AddSnack(ctx context.Context, req *sipb.AddSnackRequest) (*sipb.AddSnackResponse, error) {
	createdSnack, createdLocation, err := s.c.AddSnack(ctx, req.GetSnackBarcode(), req.GetLocationName())
	// Even if we have an error, we still need to fill in if foreign values
	// were created.
	return &sipb.AddSnackResponse{
		SnackCreated:    createdSnack,
		LocationCreated: createdLocation,
	}, err
}

func (s *snackInventoryServer) ListContents(ctx context.Context, req *sipb.ListContentsRequest) (*sipb.ListContentsResponse, error) {
	return &sipb.ListContentsResponse{}, nil
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
