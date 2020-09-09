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
