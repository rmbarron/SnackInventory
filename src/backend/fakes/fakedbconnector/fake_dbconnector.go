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

// Package fakedbconnector implements a fake connector.go impl for testing.
package fakedbconnector

import (
	"context"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
)

type FakeDBConnector struct {
	CreateSnackErr error
	ListSnacksRes  []*sipb.Snack
	ListSnacksErr  error
	UpdateSnackErr error
	DeleteSnackErr error

	CreateLocationErr error
	ListLocationsRes  []*sipb.Location
	ListLocationsErr  error
	DeleteLocationErr error

	AddSnackCreatedSnack    bool
	AddSnackCreatedLocation bool
	AddSnackErr             error
	ListContentsRes         map[string]map[*sipb.Snack]int
	ListContentsErr         error
}

func (f *FakeDBConnector) CreateSnack(_ context.Context, _, _ string) error {
	return f.CreateSnackErr
}

func (f *FakeDBConnector) ListSnacks(_ context.Context) ([]*sipb.Snack, error) {
	if f.ListSnacksErr != nil {
		return nil, f.ListSnacksErr
	}
	return f.ListSnacksRes, nil
}

func (f *FakeDBConnector) UpdateSnack(_ context.Context, _, _ string) error {
	return f.UpdateSnackErr
}

func (f *FakeDBConnector) DeleteSnack(_ context.Context, _ string) error {
	return f.DeleteSnackErr
}

func (f *FakeDBConnector) CreateLocation(_ context.Context, _ string) error {
	return f.CreateLocationErr
}

func (f *FakeDBConnector) ListLocations(_ context.Context) ([]*sipb.Location, error) {
	if f.ListLocationsErr != nil {
		return nil, f.ListLocationsErr
	}
	return f.ListLocationsRes, nil
}

func (f *FakeDBConnector) DeleteLocation(_ context.Context, _ string) error {
	return f.DeleteLocationErr
}

func (f *FakeDBConnector) AddSnack(_ context.Context, _, _ string) (bool, bool, error) {
	return f.AddSnackCreatedSnack, f.AddSnackCreatedLocation, f.AddSnackErr
}

func (f *FakeDBConnector) ListContents(_ context.Context, _ string) (map[string]map[*sipb.Snack]int, error) {
	if f.ListContentsErr != nil {
		return map[string]map[*sipb.Snack]int{}, f.ListContentsErr
	}
	return f.ListContentsRes, nil
}
