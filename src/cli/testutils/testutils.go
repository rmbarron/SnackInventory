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

// Package testutils provides helper functions around common test
// setup tasks.
package testutils

import (
	"net"
	"testing"

	"github.com/rmbarron/SnackInventory/src/backend/fakeserver"
	"google.golang.org/grpc"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
)

// StartTestServer launches a grpc Server on a dynamically chosen port on
// localhost. Returns the address to the new server and a close function.
// Always defer the close function after obtaining:
//
// addr, close := testutils.StartTestServer(t, fsi)
// defer close()
func StartTestServer(t *testing.T, fsi *fakeserver.FakeSnackInventoryServer) (addr string, close func()) {
	t.Helper()
	// Use port 0 for the OS to choose an open port.
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("net.Listen(%q, %q) = got err %v, want nil", "tcp", ":0", err)
	}

	// Serve using our fake server in a separate goroutine.
	grpcServer := grpc.NewServer()
	svc := sipb.NewSnackInventoryService(fsi)
	sipb.RegisterSnackInventoryService(grpcServer, svc)
	go grpcServer.Serve(lis)

	return lis.Addr().String(), func() { lis.Close(); grpcServer.Stop() }
}
