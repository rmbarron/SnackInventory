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

func TestCreateSnack(t *testing.T) {
	fdbc := &fakedbconnector.FakeDBConnector{}

	req := &sipb.CreateSnackRequest{
		Snack: &sipb.Snack{Barcode: "1"},
	}

	si := snackInventoryServer{c: fdbc}
	if _, err := si.CreateSnack(context.Background(), req); err != nil {
		t.Fatalf("si.CreateSnack(ctx, %v) = got err %v, want err nil", req, err)
	}
}

func TestCreateSnack_AlreadyExists(t *testing.T) {
	fdbc := &fakedbconnector.FakeDBConnector{
		CreateSnackErr: status.Error(codes.AlreadyExists, "already exists"),
	}

	req := &sipb.CreateSnackRequest{
		Snack: &sipb.Snack{Barcode: "1"},
	}

	si := snackInventoryServer{c: fdbc}
	if _, err := si.CreateSnack(context.Background(), req); err == nil {
		t.Fatalf("si.CreateSnack(ctx, %v) = got err nil, want err", req)
	}
}

func TestCreateSnack_Internal(t *testing.T) {
	fdbc := &fakedbconnector.FakeDBConnector{
		CreateSnackErr: status.Error(codes.Internal, "internally failed"),
	}

	req := &sipb.CreateSnackRequest{
		Snack: &sipb.Snack{Barcode: "1"},
	}

	si := snackInventoryServer{c: fdbc}
	if _, err := si.CreateSnack(context.Background(), req); err == nil {
		t.Fatalf("si.CreateSnack(ctx, %v) = got err nil, want err", req)
	}
}

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

func TestUpdateSnack(t *testing.T) {
	fdbc := &fakedbconnector.FakeDBConnector{}

	si := snackInventoryServer{c: fdbc}
	req := &sipb.UpdateSnackRequest{
		Snack: &sipb.Snack{
			Barcode: "123",
			Name:    "testsnack",
		},
	}
	if _, err := si.UpdateSnack(context.Background(), req); err != nil {
		t.Fatalf("si.UpdateSnack(ctx, %v) = got err %v, want err nil", req, err)
	}
}

func TestUpdateSnack_Error(t *testing.T) {
	fdbc := &fakedbconnector.FakeDBConnector{
		UpdateSnackErr: status.Error(codes.Internal, "something failed"),
	}

	si := snackInventoryServer{c: fdbc}
	req := &sipb.UpdateSnackRequest{
		Snack: &sipb.Snack{
			Barcode: "123",
			Name:    "testsnack",
		},
	}
	if _, err := si.UpdateSnack(context.Background(), req); err == nil {
		t.Fatalf("si.UpdateSnack(ctx, %v) = got err nil, want err", req)
	}
}

func TestDeleteSnack(t *testing.T) {
	fdbc := &fakedbconnector.FakeDBConnector{}

	si := snackInventoryServer{c: fdbc}
	req := &sipb.DeleteSnackRequest{Barcode: "123"}
	if _, err := si.DeleteSnack(context.Background(), req); err != nil {
		t.Fatalf("si.DeleteSnack(ctx, %v) = got err %v, want err nil", req, err)
	}
}

func TestDeleteSnack_Error(t *testing.T) {
	fdbc := &fakedbconnector.FakeDBConnector{
		DeleteSnackErr: status.Error(codes.Internal, "something failed"),
	}

	si := snackInventoryServer{c: fdbc}
	req := &sipb.DeleteSnackRequest{Barcode: "123"}
	if _, err := si.DeleteSnack(context.Background(), req); err == nil {
		t.Fatalf("si.DeleteSnack(ctx, %v) = got err nil, want err", req)
	}
}
