// Package fakeserver implements a fake SnackInventoryService for testing.
package fakeserver

import (
	"context"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
)

// FakeSnackInventoryServer implements sipb.SnackInventoryService.
type FakeSnackInventoryServer struct {
	// CreateSnackRes is the response for a request.
	CreateSnackRes *sipb.CreateSnackResponse
	// CreateSnackErr is the error returned on a call.
	// For testing behavior based on error status, use `status.Error`.
	CreateSnackErr error
}

// CreateSnack creates a snack in SnackInventory.
func (f *FakeSnackInventoryServer) CreateSnack(_ context.Context, _ *sipb.CreateSnackRequest) (*sipb.CreateSnackResponse, error) {
	if f.CreateSnackErr != nil {
		return &sipb.CreateSnackResponse{}, f.CreateSnackErr
	}
	return f.CreateSnackRes, nil
}
