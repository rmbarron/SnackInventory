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

// Package testutils provides helper functions around common test
// setup tasks for src/backend/*.
package testutils

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	mysqltest "github.com/lestrrat-go/test-mysqld"
	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
)

// StartMysqldT starts a local instance of mysqld. Returns the DB and a close
// function. Returned DB has no databases or tables written.
//
// db, close := testutils.StartMysqldT(ctx, t)
// defer close()
func StartMysqldT(ctx context.Context, t *testing.T) (*sql.DB, func()) {
	t.Helper()

	mysqld, err := mysqltest.NewMysqld(nil)
	if err != nil {
		t.Fatalf("mysqltest.NewMysqld(nil) = got err %v, want err nil", err)
	}

	db, err := sql.Open("mysql", mysqld.DSN())
	if err != nil {
		t.Fatalf("sql.Open(%q, %q) = got err %v, want err nil", "mysql", mysqld.DSN(), err)
	}

	if err = db.PingContext(ctx); err != nil {
		t.Fatalf("db.PingContext(ctx) = got err %v, want err nil", err)
	}

	return db, mysqld.Stop
}

// CreateDatabaseT creates a database and moves the cursor into it.
// Assumes DB cursor is not on a database.
func CreateDatabaseT(ctx context.Context, t *testing.T, db *sql.DB) {
	t.Helper()

	if _, err := db.ExecContext(ctx, "CREATE DATABASE SnackInventory"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", "CREATE DATABASE SnackInventory", err)
	}
	if _, err := db.ExecContext(ctx, "USE SnackInventory"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", "USE SnackInventory", err)
	}
}

// CreateTablesT creates tables to satisfy SnackInventory storage model.
// Assumes cursor is in database.
func CreateTablesT(ctx context.Context, t *testing.T, db *sql.DB) {
	if _, err := db.ExecContext(ctx, "CREATE TABLE SnackRegistry ( barcode VARCHAR(20) PRIMARY KEY, name VARCHAR(255))"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil",
			"CREATE TABLE SnackRegistry ( barcode VARCHAR(20) PRIMARY KEY, name VARCHAR(255))", err)
	}
	if _, err := db.ExecContext(ctx, "CREATE TABLE LocationRegistry ( name VARCHAR(30) PRIMARY KEY)"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil",
			"CREATE TABLE LocationRegistry ( name VARCHAR(30) PRIMARY KEY)", err)
	}
	createMappingTableQuery := `CREATE TABLE LocationContents (
		ContentID int NOT NULL AUTO_INCREMENT,
		snackBarcode VARCHAR(20), locationName VARCHAR(30),
		numPresent int NOT NULL,
		PRIMARY KEY (ContentID),
		FOREIGN KEY (snackBarcode) REFERENCES SnackRegistry(barcode),
		FOREIGN KEY (locationName) REFERENCES LocationRegistry(name))`
	if _, err := db.ExecContext(ctx, createMappingTableQuery); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil",
			createMappingTableQuery, err)
	}
}

// DropTablesT drops tables in the current database corresponding to
// SnackInventory's storage model. Assumes cursor is in database.
func DropTablesT(ctx context.Context, t *testing.T, db *sql.DB) {
	if _, err := db.ExecContext(ctx, "DROP TABLE LocationContents"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil",
			"DROP TABLE LocationContext", err)
	}
	if _, err := db.ExecContext(ctx, "DROP TABLE SnackRegistry, LocationRegistry"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil",
			"DROP TABLE SnackRegistry, LocationRegistry", err)
	}
}

// AddSnackT adds a given Snack to DB's SnackRegistry table.
// Assumes DB cursor is in the correct database already.
func AddSnackT(ctx context.Context, t *testing.T, db *sql.DB, snack *sipb.Snack) {
	t.Helper()

	barcode := snack.GetBarcode()
	name := snack.GetName()

	query := fmt.Sprintf("INSERT INTO SnackRegistry (barcode, name) VALUES(%q, %q)", barcode, name)

	if _, err := db.ExecContext(ctx, query); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", query, err)
	}
}

// AddLocationT adds a given Location to DB's LocationRegistry table.
// Assumes DB cursor is in the correct database already.
func AddLocationT(ctx context.Context, t *testing.T, db *sql.DB, location *sipb.Location) {
	t.Helper()

	name := location.GetName()

	query := fmt.Sprintf("INSERT INTO LocationRegistry (name) VALUES(%q)", name)

	if _, err := db.ExecContext(ctx, query); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", query, err)
	}
}

// AddSnackMappingT inserts a snack:location entry in LocationContents table.
// Assumes DB cursor is in the correct database already.
func AddSnackMappingT(ctx context.Context, t *testing.T, db *sql.DB, snackBarcode, locationName string, numPresent int) {
	t.Helper()

	query := fmt.Sprintf("INSERT INTO LocationContents (snackBarcode, locationName, numPresent) VALUES(%q, %q, %d)", snackBarcode, locationName, numPresent)

	if _, err := db.ExecContext(ctx, query); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", query, err)
	}
}

func snackKeyMapComparerFunc(x, y map[*sipb.Snack]int) bool {
	if len(x) != len(y) {
		return false
	}

	xKeys := []*sipb.Snack{}
	for k := range x {
		xKeys = append(xKeys, k)
	}
	sort.SliceStable(xKeys, func(i, j int) bool {
		return xKeys[i].GetBarcode() < xKeys[j].GetBarcode()
	})

	yKeys := []*sipb.Snack{}
	for k := range y {
		yKeys = append(yKeys, k)
	}
	sort.SliceStable(yKeys, func(i, j int) bool {
		return yKeys[i].GetBarcode() < yKeys[j].GetBarcode()
	})

	if !cmp.Equal(xKeys, yKeys, cmpopts.IgnoreUnexported(sipb.Snack{})) {
		return false
	}

	for i := range xKeys {
		if !cmp.Equal(x[xKeys[i]], y[yKeys[i]]) {
			return false
		}
	}
	return true
}

// SnackKeyMapComparer compares a single-depth map with a *sipb.Snack as the key.
// cmp.Equal really does not like comparing proto.Message structs, even with
// protocmp.Transform option. However, it does a good job in checking equality
// of a []*proto.Message. This comparer:
// - checks that given maps are same length
// - extracts & sorts keys, then compares the sorted slice.
// - iterates through the map with the sorted key, comparing values.
//
// slice keys and map values are compared using map.Equal, so maps with a
// larger depth than 1 _may_ work but are not guaranteed.
//
// Usage:
//   if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(sipb.Snack{}), testutils.SnackKeyMapComparer()) {
//     t.Error()
//   }
func SnackKeyMapComparer() cmp.Option {
	return cmp.Comparer(snackKeyMapComparerFunc)
}
