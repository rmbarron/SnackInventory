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
	"database/sql"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	mysqltest "github.com/lestrrat-go/test-mysqld"
	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
)

func TestListSnacks(t *testing.T) {
	ctx := context.Background()

	mysqld, err := mysqltest.NewMysqld(nil)
	if err != nil {
		t.Fatalf("mysqltest.NewMysqld(nil) = got err %v, want err nil", err)
	}
	defer mysqld.Stop()

	db, err := sql.Open("mysql", mysqld.DSN())
	if err != nil {
		t.Fatalf("sql.Open(%q, %q) = got err %v, want err nil", "mysql", mysqld.DSN(), err)
	}

	if err = db.PingContext(ctx); err != nil {
		t.Fatalf("db.PingContext(ctx) = got err %v, want err nil", err)
	}

	// Fill in initial data in the tmp instance.
	if _, err = db.ExecContext(ctx, "CREATE DATABASE SnackInventory"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", "CREATE DATABASE SnackInventory", err)
	}
	if _, err = db.ExecContext(ctx, "USE SnackInventory"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", "USE SnackInventory", err)
	}
	if _, err = db.ExecContext(ctx, "CREATE TABLE SnackRegistry ( barcode VARCHAR(20) PRIMARY KEY, name VARCHAR(255))"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil",
			"CREATE TABLE SnackRegistry ( barcode VARCHAR(20) PRIMARY KEY, name VARCHAR(255))", err)
	}
	if _, err = db.ExecContext(ctx, "INSERT INTO SnackRegistry (barcode, name) VALUES('123', 'testsnack')"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil",
			"INSERT INTO SnackRegistry (barcode, name) VALUES(123, testsnack)", err)
	}

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

	mysqld, err := mysqltest.NewMysqld(nil)
	if err != nil {
		t.Fatalf("mysqltest.NewMysqld(nil) = got err %v, want err nil", err)
	}
	defer mysqld.Stop()

	db, err := sql.Open("mysql", mysqld.DSN())
	if err != nil {
		t.Fatalf("sql.Open(%q, %q) = got err %v, want err nil", "mysql", mysqld.DSN(), err)
	}

	if err = db.PingContext(ctx); err != nil {
		t.Fatalf("db.PingContext(ctx) = got err %v, want err nil", err)
	}

	// Fill in initial data in the tmp instance.
	if _, err = db.ExecContext(ctx, "CREATE DATABASE SnackInventory"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", "CREATE DATABASE SnackInventory", err)
	}
	if _, err = db.ExecContext(ctx, "USE SnackInventory"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", "USE SnackInventory", err)
	}

	si := &SQLImpl{db: db}
	if _, err = si.ListSnacks(ctx); err == nil {
		t.Fatal("si.ListSnacks(ctx) = got err nil, want err")
	}
}
