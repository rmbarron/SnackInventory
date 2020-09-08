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

// Tests for this package spin up a local instance of mysqld and tear them down
// before and after each test. Tests here are expected to be slow.
package connector

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/rmbarron/SnackInventory/src/backend/server/testutils"
	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
)

func TestCreateSnack(t *testing.T) {
	ctx := context.Background()

	db, close := testutils.StartMysqldT(ctx, t)
	defer close()

	testutils.CreateDatabaseAndTablesT(ctx, t, db)

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
}

func TestCreateSnack_SelectError(t *testing.T) {
	ctx := context.Background()

	db, close := testutils.StartMysqldT(ctx, t)
	defer close()

	// Create database, but without any tables.
	if _, err := db.ExecContext(ctx, "CREATE DATABASE SnackInventory"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", "CREATE DATABASE SnackInventory", err)
	}
	if _, err := db.ExecContext(ctx, "USE SnackInventory"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", "USE SnackInventory", err)
	}

	si := &SQLImpl{db: db}
	if err := si.CreateSnack(ctx, "1", "testsnack"); err == nil {
		t.Fatalf("si.CreateSnack(ctx, %q, %q) = got err nil, want err", "1", "testsnack")
	}
}

func TestCreateSnack_AlreadyExists(t *testing.T) {
	ctx := context.Background()

	db, close := testutils.StartMysqldT(ctx, t)
	defer close()

	testutils.CreateDatabaseAndTablesT(ctx, t, db)
	testutils.AddSnackT(ctx, t, db, &sipb.Snack{Barcode: "1", Name: "testsnack"})

	si := &SQLImpl{db: db}
	if err := si.CreateSnack(ctx, "1", "testsnack"); err == nil {
		t.Fatalf("si.CreateSnack(ctx, %q, %q) = got err nil, want err", "1", "testsnack")
	}
}

func TestListSnacks(t *testing.T) {
	ctx := context.Background()

	db, close := testutils.StartMysqldT(ctx, t)
	defer close()

	testutils.CreateDatabaseAndTablesT(ctx, t, db)
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
}

// TestListSnacks_SelectError creates and connects to a database, but the
// "SnackRegistry" table doesn't exist, causing SELECT statements to fail.
func TestListSnacks_SelectError(t *testing.T) {
	ctx := context.Background()

	db, close := testutils.StartMysqldT(ctx, t)
	defer close()

	// Create database, but without any tables.
	if _, err := db.ExecContext(ctx, "CREATE DATABASE SnackInventory"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", "CREATE DATABASE SnackInventory", err)
	}
	if _, err := db.ExecContext(ctx, "USE SnackInventory"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", "USE SnackInventory", err)
	}

	si := &SQLImpl{db: db}
	if _, err := si.ListSnacks(ctx); err == nil {
		t.Fatal("si.ListSnacks(ctx) = got err nil, want err")
	}
}

func TestDeleteSnack(t *testing.T) {
	ctx := context.Background()

	db, close := testutils.StartMysqldT(ctx, t)
	defer close()

	testutils.CreateDatabaseAndTablesT(ctx, t, db)
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
}
