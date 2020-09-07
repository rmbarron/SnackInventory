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
package main

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/rmbarron/SnackInventory/src/backend/fakes/fakedbconnector"
	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestListSnacks(t *testing.T) {
	snack := &sipb.Snack{
		Barcode: "123",
		Name:    "snack",
	}
	fdbc := &fakedbconnector.FakeDBConnector{
		ListSnacksRes: []*sipb.Snack{snack},
	}

	si := snackInventoryServer{c: fdbc}
	res, err := si.ListSnacks(context.Background(), &sipb.ListSnacksRequest{})
	if err != nil {
		t.Fatalf("si.ListSnacks(ctx, &sipb.ListSnacksRequest{}) = got err %v, want err nil", err)
	}

	want := &sipb.ListSnacksResponse{
		Snacks: []*sipb.Snack{snack},
	}
	if diff := cmp.Diff(
		res, want,
		cmpopts.IgnoreUnexported(sipb.ListSnacksResponse{}),
		cmpopts.IgnoreUnexported(sipb.Snack{})); diff != "" {
		t.Fatalf("si.ListSnacks(ctx, &sipb.ListSnacksRequest{}) = got diff (-got +want): %s", diff)
	}
}

func TestListSnacks_StorageError(t *testing.T) {
	fdbc := &fakedbconnector.FakeDBConnector{
		ListSnacksErr: status.Error(codes.Internal, "encountered error"),
	}

	si := snackInventoryServer{c: fdbc}
	if _, err := si.ListSnacks(context.Background(), &sipb.ListSnacksRequest{}); err == nil {
		t.Fatal("si.ListSnacks(ctx, &sipb.ListSnacksRequest{}) = got err nil, want err")
	}
}
