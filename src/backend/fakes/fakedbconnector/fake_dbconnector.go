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
