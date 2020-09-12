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
package connector

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/rmbarron/SnackInventory/src/backend/server/testutils"
	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
)

// Implementation note: Spinning up a full mariadb / mysqld instance is slow.
// Adding a new test that sets up and tears down a full instance adds ~10s.
// However, we still want to make sure database modification from one test
// don't spill over into other tests.
// So, the testcases here follow a specific pattern of spinning up the instance
// as setup, then each subtest creates tables -> performs test -> drops tables.

// TestSuccess is a parent test to create a mariadb instance for subtests.
func TestSuccess(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	db, close := testutils.StartMysqldT(ctx, t)
	defer close()

	testutils.CreateDatabaseT(ctx, t, db)

	t.Run("CreateSnack", func(t *testing.T) {
		testutils.CreateTablesT(ctx, t, db)
		defer testutils.DropTablesT(ctx, t, db)

		si := &SQLImpl{db: db}
		if err := si.CreateSnack(ctx, "1", "testsnack"); err != nil {
			t.Fatalf("si.CreateSnack(ctx, %q, %q) = got err %v, want err nil", "1", "testsnack", err)
		}

		want := []*sipb.Snack{
			{
				Barcode: "1",
				Name:    "testsnack",
			},
		}
		got, err := si.ListSnacks(ctx)
		if err != nil {
			t.Fatalf("si.ListSnacks(ctx) = got err %v, want err nil", err)
		}
		if diff := cmp.Diff(got, want, cmpopts.IgnoreUnexported(sipb.Snack{})); diff != "" {
			t.Fatalf("si.ListSnacks(ctx) = got diff (-got +want): %s", diff)
		}
	})

	t.Run("ListSnacks", func(t *testing.T) {
		testutils.CreateTablesT(ctx, t, db)
		defer testutils.DropTablesT(ctx, t, db)

		testutils.AddSnackT(ctx, t, db, &sipb.Snack{Barcode: "123", Name: "testsnack"})

		si := &SQLImpl{db: db}
		got, err := si.ListSnacks(ctx)
		if err != nil {
			t.Fatalf("si.ListSnacks(ctx) = got err %v, want err nil", err)
		}

		want := []*sipb.Snack{
			{
				Barcode: "123",
				Name:    "testsnack",
			},
		}

		if diff := cmp.Diff(got, want, cmpopts.IgnoreUnexported(sipb.Snack{})); diff != "" {
			t.Fatalf("si.ListSnacks(ctx) = got diff (-got +want): %s", diff)
		}
	})

	t.Run("UpdateSnack", func(t *testing.T) {
		testutils.CreateTablesT(ctx, t, db)
		defer testutils.DropTablesT(ctx, t, db)

		testutils.AddSnackT(ctx, t, db, &sipb.Snack{Barcode: "123", Name: "testsnack"})

		si := &SQLImpl{db: db}
		if err := si.UpdateSnack(ctx, "123", "realsnack"); err != nil {
			t.Fatalf("si.UpdateSnack(ctx, %q, %q) = got err %v, want err nil", "123", "realsnack", err)
		}

		want := []*sipb.Snack{
			{
				Barcode: "123",
				Name:    "realsnack",
			},
		}
		got, err := si.ListSnacks(ctx)
		if err != nil {
			t.Fatalf("si.ListSnacks(ctx) = got err %v, want err nil", err)
		}
		if diff := cmp.Diff(got, want, cmpopts.IgnoreUnexported(sipb.Snack{})); diff != "" {
			t.Fatalf("si.ListSnacks(ctx) = got diff (-got +want): %s\n", diff)
		}
	})

	t.Run("DeleteSnack", func(t *testing.T) {
		testutils.CreateTablesT(ctx, t, db)
		defer testutils.DropTablesT(ctx, t, db)

		testutils.AddSnackT(ctx, t, db, &sipb.Snack{Barcode: "123", Name: "testsnack"})

		si := &SQLImpl{db: db}
		if err := si.DeleteSnack(ctx, "123"); err != nil {
			t.Fatalf("si.DeleteSnack(ctx, %q) = got err %v, want err nil", "123", err)
		}

		got, err := si.ListSnacks(ctx)
		if err != nil {
			t.Fatalf("si.ListSnacks(ctx) = got err %v, want err nil", err)
		}
		if len(got) != 0 {
			t.Fatalf("si.ListSnacks(ctx) = got %v, want []*sipb.Snack{}", got)
		}
	})

	t.Run("CreateLocation", func(t *testing.T) {
		testutils.CreateTablesT(ctx, t, db)
		defer testutils.DropTablesT(ctx, t, db)

		si := &SQLImpl{db: db}
		if err := si.CreateLocation(ctx, "fridge"); err != nil {
			t.Fatalf("si.CreateLocation(ctx, %q) = got err %v, want err nil", "fridge", err)
		}

		want := []*sipb.Location{
			{
				Name: "fridge",
			},
		}
		got, err := si.ListLocations(ctx)
		if err != nil {
			t.Fatalf("si.ListSnacks(ctx) = got err %v, want err nil", err)
		}
		if diff := cmp.Diff(got, want, cmpopts.IgnoreUnexported(sipb.Location{})); diff != "" {
			t.Fatalf("si.ListLocations(ctx) = got diff (-got +want): %s", diff)
		}
	})

	t.Run("ListLocations", func(t *testing.T) {
		testutils.CreateTablesT(ctx, t, db)
		defer testutils.DropTablesT(ctx, t, db)

		testutils.AddLocationT(ctx, t, db, &sipb.Location{Name: "fridge"})

		si := &SQLImpl{db: db}
		got, err := si.ListLocations(ctx)
		if err != nil {
			t.Fatalf("si.ListLocations(ctx) = got err %v, want err nil", err)
		}

		want := []*sipb.Location{
			{
				Name: "fridge",
			},
		}

		if diff := cmp.Diff(got, want, cmpopts.IgnoreUnexported(sipb.Location{})); diff != "" {
			t.Fatalf("si.ListLocations(ctx) = got diff (-got +want): %s", diff)
		}
	})

	t.Run("DeleteLocation", func(t *testing.T) {
		testutils.CreateTablesT(ctx, t, db)
		defer testutils.DropTablesT(ctx, t, db)

		testutils.AddLocationT(ctx, t, db, &sipb.Location{Name: "fridge"})

		si := &SQLImpl{db: db}
		if err := si.DeleteLocation(ctx, "fridge"); err != nil {
			t.Fatalf("si.DeleteLocation(ctx, %q) = got err %v, want err nil", "fridge", err)
		}

		got, err := si.ListLocations(ctx)
		if err != nil {
			t.Fatalf("si.ListLocations(ctx) = got err %v, want err nil", err)
		}
		if len(got) != 0 {
			t.Fatalf("si.ListLocations(ctx) = got %v, want []*sipb.Location{}", got)
		}
	})
}

// TestError is a parent test to create a mariadb instance for subtests.
func TestError(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	db, close := testutils.StartMysqldT(ctx, t)
	defer close()

	testutils.CreateDatabaseT(ctx, t, db)

	// Try to read from a database with no tables, causing SELECT to fail.
	t.Run("CreateSnack_SelectError", func(t *testing.T) {
		si := &SQLImpl{db: db}
		if err := si.CreateSnack(ctx, "1", "testsnack"); err == nil {
			t.Fatalf("si.CreateSnack(ctx, %q, %q) = got err nil, want err", "1", "testsnack")
		}
	})

	t.Run("CreateSnack_AlreadyExists", func(t *testing.T) {
		testutils.CreateTablesT(ctx, t, db)
		defer testutils.DropTablesT(ctx, t, db)

		testutils.AddSnackT(ctx, t, db, &sipb.Snack{Barcode: "1", Name: "testsnack"})

		si := &SQLImpl{db: db}
		if err := si.CreateSnack(ctx, "1", "testsnack"); err == nil {
			t.Fatalf("si.CreateSnack(ctx, %q, %q) = got err nil, want err", "1", "testsnack")
		}
	})

	t.Run("ListSnacks_SelectError", func(t *testing.T) {
		si := &SQLImpl{db: db}
		if _, err := si.ListSnacks(ctx); err == nil {
			t.Fatal("si.ListSnacks(ctx) = got err nil, want err")
		}
	})

	t.Run("UpdateSnack_Error", func(t *testing.T) {
		si := &SQLImpl{db: db}
		if err := si.UpdateSnack(ctx, "123", "realsnack"); err == nil {
			t.Fatalf("si.UpdateSnack(ctx, %q, %q) = got err nil, want err", "123", "realsnack")
		}
	})

	t.Run("CreateLocation_SelectError", func(t *testing.T) {
		si := &SQLImpl{db: db}
		if err := si.CreateLocation(ctx, "fridge"); err == nil {
			t.Fatalf("si.CreateLocation(ctx, %q) = got err nil, want err", "fridge")
		}
	})

	t.Run("CreateLocation_AlreadyExists", func(t *testing.T) {
		testutils.CreateTablesT(ctx, t, db)
		defer testutils.DropTablesT(ctx, t, db)

		testutils.AddLocationT(ctx, t, db, &sipb.Location{Name: "fridge"})

		si := &SQLImpl{db: db}
		if err := si.CreateLocation(ctx, "fridge"); err == nil {
			t.Fatalf("si.CreateLocation(ctx, %q) = got err nil, want err", "fridge")
		}
	})

	t.Run("ListLocations_SelectError", func(t *testing.T) {
		si := &SQLImpl{db: db}
		if _, err := si.ListLocations(ctx); err == nil {
			t.Fatalf("si.ListLocations(ctx) = got err nil, want err")
		}
	})
}
