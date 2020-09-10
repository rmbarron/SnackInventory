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
	"testing"

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

// CreateDatabaseAndTablesT creates database and tables to satisfy
// SnackInventory storage model. Assumes DB cursor is not on a database.
func CreateDatabaseAndTablesT(ctx context.Context, t *testing.T, db *sql.DB) {
	t.Helper()

	if _, err := db.ExecContext(ctx, "CREATE DATABASE SnackInventory"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", "CREATE DATABASE SnackInventory", err)
	}
	if _, err := db.ExecContext(ctx, "USE SnackInventory"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil", "USE SnackInventory", err)
	}
	if _, err := db.ExecContext(ctx, "CREATE TABLE SnackRegistry ( barcode VARCHAR(20) PRIMARY KEY, name VARCHAR(255))"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil",
			"CREATE TABLE SnackRegistry ( barcode VARCHAR(20) PRIMARY KEY, name VARCHAR(255))", err)
	}
	if _, err := db.ExecContext(ctx, "CREATE TABLE LocationRegistry ( name VARCHAR(30) PRIMARY KEY)"); err != nil {
		t.Fatalf("db.ExecContext(ctx, %q) = got err %v, want err nil",
			"CREATE TABLE LocationRegistry ( name VARCHAR(30) PRIMARY KEY)", err)
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
