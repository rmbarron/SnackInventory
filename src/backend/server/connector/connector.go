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

// Package connector provides a common interface to SnackInventory's storage model.
// This package allows for implementing different connectors to different
// storage systems (SQL, NoSQL, Google Sheets, etc). The implementation used
// can then be decided via flag in server.go. Functions offered via the
// interface tend to follow the API functions, making the server implementation
// very simple.
package connector

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL driver.
	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SQLImpl implements a connector a SQL DB.
// SQLImpl connects to an arbitrary address:DBName, but assumes the presence of
// "SnackRegistry" & "Location" tables.
type SQLImpl struct {
	db *sql.DB
}

// NewSQLImpl connects to SQL and creates a SQLImpl instance.
func NewSQLImpl(ctx context.Context, user, password, hostport, dbname string) (*SQLImpl, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, hostport, dbname))
	if err != nil {
		return nil, err
	}
	// Verify the connection to SQL is open.
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return &SQLImpl{db: db}, nil
}

// CreateSnack creates a snack in the sql database.
// Returns an AlreadyExists error if it does.
func (s *SQLImpl) CreateSnack(ctx context.Context, barcode, name string) error {
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM SnackRegistry WHERE barcode IN (?)", barcode)
	if err != nil {
		return err
	}
	defer rows.Close()
	// Check if the value already exists by whether there are results in the Rows.
	if rows.Next() {
		return status.Errorf(codes.AlreadyExists, "barcode %q already has an entry", barcode)
	}
	if _, err := s.db.ExecContext(ctx, "INSERT INTO SnackRegistry (barcode, name) VALUES(?, ?)", barcode, name); err != nil {
		return err
	}
	return nil
}

// ListSnacks reads all snacks currently registered to SnackInventory.
func (s *SQLImpl) ListSnacks(ctx context.Context) ([]*sipb.Snack, error) {
	var retVal []*sipb.Snack
	rows, err := s.db.QueryContext(ctx, "SELECT barcode, name FROM SnackRegistry")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var barcode string
		var name string
		if err = rows.Scan(&barcode, &name); err != nil {
			return nil, err
		}
		retVal = append(retVal, &sipb.Snack{Barcode: barcode, Name: name})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return retVal, nil
}

// UpdateSnack updates a single snack in place in SnackInventory.
func (s *SQLImpl) UpdateSnack(ctx context.Context, barcode, name string) error {
	if _, err := s.db.ExecContext(ctx, "UPDATE SnackRegistry SET name = ? WHERE barcode IN (?)",
		name, barcode); err != nil {
		return err
	}
	return nil
}

// DeleteSnack deletes a single snack from SnackInventory.
func (s *SQLImpl) DeleteSnack(ctx context.Context, barcode string) error {
	if _, err := s.db.ExecContext(ctx, "DELETE FROM SnackRegistry WHERE barcode IN (?)", barcode); err != nil {
		return err
	}
	return nil
}

// CreateLocation adds a new location to SnackInventory.
// Returns an AlreadyExists error if it does.
func (s *SQLImpl) CreateLocation(ctx context.Context, name string) error {
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM LocationRegistry WHERE name IN (?)", name)
	if err != nil {
		return err
	}
	defer rows.Close()
	// Check if the value already exists by whether there are results in the Rows.
	if rows.Next() {
		return status.Errorf(codes.AlreadyExists, "name %q already has an entry", name)
	}
	if _, err := s.db.ExecContext(ctx, "INSERT INTO LocationRegistry (name) VALUES(?)", name); err != nil {
		return err
	}
	return nil
}
