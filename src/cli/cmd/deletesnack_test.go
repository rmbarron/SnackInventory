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
package cmd

import (
	"net"
	"testing"

	"github.com/rmbarron/SnackInventory/src/backend/fakeserver"
	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeleteSnack(t *testing.T) {
	fsi := fakeserver.FakeSnackInventoryServer{
		DeleteSnackRes: &sipb.DeleteSnackResponse{},
	}
	// Use port 0 for the OS to choose an open port.
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("net.Listen(%q, %q) = got err %v, want nil", "tcp", ":0", err)
	}
	defer lis.Close()

	// Serve using our fake server in a separate goroutine.
	grpcServer := grpc.NewServer()
	svc := sipb.NewSnackInventoryService(&fsi)
	sipb.RegisterSnackInventoryService(grpcServer, svc)
	go grpcServer.Serve(lis)
	defer grpcServer.Stop()

	// Inject the address of our fake server to the address flag variable.
	tmpAddr := address
	address = lis.Addr().String()
	defer func() { address = tmpAddr }()

	if err = deleteSnack(nil, nil); err != nil {
		t.Fatalf("deleteSnack(nil, nil) = got err %v, want nil", err)
	}
}

func TestDeleteSnack_ServerError(t *testing.T) {
	fsi := fakeserver.FakeSnackInventoryServer{
		DeleteSnackErr: status.Error(codes.ResourceExhausted, "server overloaded"),
	}
	// Use port 0 for the OS to choose an open port.
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("net.Listen(%q, %q) = got err %v, want nil", "tcp", ":0", err)
	}
	defer lis.Close()

	// Serve using our fake server in a separate goroutine.
	grpcServer := grpc.NewServer()
	svc := sipb.NewSnackInventoryService(&fsi)
	sipb.RegisterSnackInventoryService(grpcServer, svc)
	go grpcServer.Serve(lis)
	defer grpcServer.Stop()

	// Inject the address of our fake server to the address flag variable.
	tmpAddr := address
	address = lis.Addr().String()
	defer func() { address = tmpAddr }()

	if err = deleteSnack(nil, nil); err == nil {
		t.Fatal("deleteSnack(nil, nil) = got err nil, want err")
	}
}
