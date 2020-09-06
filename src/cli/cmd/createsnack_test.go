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
	"testing"

	"github.com/rmbarron/SnackInventory/src/backend/fakeserver"
	"github.com/rmbarron/SnackInventory/src/cli/testutils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
)

func TestCreateSnack(t *testing.T) {
	fsi := &fakeserver.FakeSnackInventoryServer{
		CreateSnackRes: &sipb.CreateSnackResponse{},
	}
	addr, close := testutils.StartTestServer(t, fsi)
	defer close()

	// Inject the address of our fake server to the address flag variable.
	tmpAddr := address
	address = addr
	defer func() { address = tmpAddr }()

	if err := createSnack(nil, nil); err != nil {
		t.Fatalf("createSnack(nil, nil) = got err %v, want nil", err)
	}
}

func TestCreateSnack_ServerError(t *testing.T) {
	fsi := &fakeserver.FakeSnackInventoryServer{
		CreateSnackErr: status.Error(codes.AlreadyExists, "could not create snack"),
	}
	addr, close := testutils.StartTestServer(t, fsi)
	defer close()

	// Inject the address of our fake server to the address flag variable.
	tmpAddr := address
	address = addr
	defer func() { address = tmpAddr }()

	if err := createSnack(nil, nil); err == nil {
		t.Fatal("createSnack(nil, nil) = got err nil, want err")
	}
}
