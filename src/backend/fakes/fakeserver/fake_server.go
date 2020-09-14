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

// Package fakeserver implements a fake SnackInventoryService for testing.
package fakeserver

import (
	"context"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
)

// FakeSnackInventoryServer implements sipb.SnackInventoryService.
type FakeSnackInventoryServer struct {
	// SnackRegistry Operations.
	// CreateSnackRes is the response for CreateSnack.
	CreateSnackRes *sipb.CreateSnackResponse
	// CreateSnackErr is the error returned on a call.
	// For testing behavior based on error status, use `status.Error`.
	CreateSnackErr error
	ListSnacksRes  *sipb.ListSnacksResponse
	ListSnacksErr  error
	UpdateSnackRes *sipb.UpdateSnackResponse
	UpdateSnackErr error
	DeleteSnackRes *sipb.DeleteSnackResponse
	DeleteSnackErr error

	// LocationRegistry Operations.
	CreateLocationRes *sipb.CreateLocationResponse
	CreateLocationErr error
	ListLocationsRes  *sipb.ListLocationsResponse
	ListLocationsErr  error
	DeleteLocationRes *sipb.DeleteLocationResponse
	DeleteLocationErr error

	// LocationContents Operations.
	AddSnackRes     *sipb.AddSnackResponse
	AddSnackErr     error
	ListContentsRes *sipb.ListContentsResponse
	ListContentsErr error
}

// CreateSnack creates a snack in SnackInventory.
func (f *FakeSnackInventoryServer) CreateSnack(_ context.Context, _ *sipb.CreateSnackRequest) (*sipb.CreateSnackResponse, error) {
	if f.CreateSnackErr != nil {
		return &sipb.CreateSnackResponse{}, f.CreateSnackErr
	}
	return f.CreateSnackRes, nil
}

// ListSnacks lists all snacks in SnackInventory.
func (f *FakeSnackInventoryServer) ListSnacks(_ context.Context, _ *sipb.ListSnacksRequest) (*sipb.ListSnacksResponse, error) {
	if f.ListSnacksErr != nil {
		return &sipb.ListSnacksResponse{}, f.ListSnacksErr
	}
	return f.ListSnacksRes, nil
}

// UpdateSnack updates meta-values of a snack in SnackInventory.
func (f *FakeSnackInventoryServer) UpdateSnack(_ context.Context, _ *sipb.UpdateSnackRequest) (*sipb.UpdateSnackResponse, error) {
	if f.UpdateSnackErr != nil {
		return &sipb.UpdateSnackResponse{}, f.UpdateSnackErr
	}
	return f.UpdateSnackRes, nil
}

// DeleteSnack deletes a snack from SnackInventory.
func (f *FakeSnackInventoryServer) DeleteSnack(_ context.Context, _ *sipb.DeleteSnackRequest) (*sipb.DeleteSnackResponse, error) {
	if f.DeleteSnackErr != nil {
		return &sipb.DeleteSnackResponse{}, f.DeleteSnackErr
	}
	return f.DeleteSnackRes, nil
}

// CreateLocation adds a new location to SnackInventory.
func (f *FakeSnackInventoryServer) CreateLocation(_ context.Context, _ *sipb.CreateLocationRequest) (*sipb.CreateLocationResponse, error) {
	if f.CreateLocationErr != nil {
		return &sipb.CreateLocationResponse{}, f.CreateLocationErr
	}
	return f.CreateLocationRes, nil
}

// ListLocations lists all locations in SnackInventory.
func (f *FakeSnackInventoryServer) ListLocations(_ context.Context, _ *sipb.ListLocationsRequest) (*sipb.ListLocationsResponse, error) {
	if f.ListLocationsErr != nil {
		return &sipb.ListLocationsResponse{}, f.ListLocationsErr
	}
	return f.ListLocationsRes, nil
}

// DeleteLocation removes a location from SnackInventory.
func (f *FakeSnackInventoryServer) DeleteLocation(_ context.Context, _ *sipb.DeleteLocationRequest) (*sipb.DeleteLocationResponse, error) {
	if f.DeleteLocationErr != nil {
		return &sipb.DeleteLocationResponse{}, f.DeleteLocationErr
	}
	return f.DeleteLocationRes, nil
}

// AddSnack adds a snack:location mapping in SnackInventory.
func (f *FakeSnackInventoryServer) AddSnack(_ context.Context, _ *sipb.AddSnackRequest) (*sipb.AddSnackResponse, error) {
	return f.AddSnackRes, f.AddSnackErr
}

// ListContents lists all contents in SnackInventory.
func (f *FakeSnackInventoryServer) ListContents(_ context.Context, _ *sipb.ListContentsRequest) (*sipb.ListContentsResponse, error) {
	if f.ListContentsErr != nil {
		return &sipb.ListContentsResponse{}, f.ListContentsErr
	}
	return f.ListContentsRes, nil
}
