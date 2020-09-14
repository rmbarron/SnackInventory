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

	driver "github.com/go-sql-driver/mysql" // MySQL driver.
	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	sqlAlreadyExistsError      = 1062
	sqlChildRowForeignKeyError = 1452
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
	if _, err := s.db.ExecContext(ctx, "INSERT INTO SnackRegistry (barcode, name) VALUES(?, ?)", barcode, name); err != nil {
		mysqlerr, ok := err.(*driver.MySQLError)
		if ok && mysqlerr.Number == sqlAlreadyExistsError {
			return status.Error(codes.AlreadyExists, err.Error())
		}
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
	if _, err := s.db.ExecContext(ctx, "INSERT INTO LocationRegistry (name) VALUES(?)", name); err != nil {
		mysqlerr, ok := err.(*driver.MySQLError)
		if ok && mysqlerr.Number == sqlAlreadyExistsError {
			return status.Error(codes.AlreadyExists, err.Error())
		}
		return err
	}
	return nil
}

// ListLocations reads all locations currently associated with SnackInventory.
func (s *SQLImpl) ListLocations(ctx context.Context) ([]*sipb.Location, error) {
	var retVal []*sipb.Location
	rows, err := s.db.QueryContext(ctx, "SELECT name FROM LocationRegistry")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			return nil, err
		}
		retVal = append(retVal, &sipb.Location{Name: name})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return retVal, nil
}

// DeleteLocation removes a location with the given name from SnackInventory.
func (s *SQLImpl) DeleteLocation(ctx context.Context, name string) error {
	if _, err := s.db.ExecContext(ctx, "DELETE FROM LocationRegistry WHERE name IN (?)", name); err != nil {
		return err
	}
	return nil
}

// I wonder if AddSnack could be handled more gracefully from a DBA standpoint?
// Currently, it operates well from a golang perspective. However, there are
// race conditions between each of the queries. AFAICT, there is no method of
// locking feasible for SQL databases. We could use a mutex lock, but that ofc
// won't lock between replicated tasks. We could probably use a lock service,
// lock file, or lock table of some sort, but it is unclear how that would look
// if we start replicating the service.
// This problem likely becomes a moot point if the conditional update v insert
// can all be handled inside a SQL query, albeit harder for me to maintain.

// AddSnack adds a snack:location mapping to SnackInventory.
// If snack or location does not exist, a minimal entry is added.
// Bool for entry being created is returned - requested string is the key.
func (s *SQLImpl) AddSnack(ctx context.Context, snackBarcode, locationName string) (createdSnack, createdLocation bool, err error) {
	// First, try to update an existing row.
	// This should only fail if the LocationContents table does not exist.
	result, err := s.db.ExecContext(ctx, "UPDATE LocationContents SET numPresent=numPresent+1 WHERE snackBarcode IN (?) and locationName IN (?)", snackBarcode, locationName)
	if err != nil {
		return createdSnack, createdLocation, err // false, false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return createdSnack, createdLocation, err // false, false, err
	}
	// Since we're updating by all the foreign keys, this operation can only
	// affect 0 or 1 rows. If it affected 1 row, we know it succeeded.
	if rowsAffected == 1 {
		return createdSnack, createdLocation, nil
	}
	// No rows were affected, which means we need to add a new row.
	if _, err := s.db.ExecContext(ctx, "INSERT INTO LocationContents (snackBarcode, locationName, numPresent) VALUES(?, ?, 1)", snackBarcode, locationName); err != nil {
		// TODO: We can likely rephrase this to remove an indentation block.
		mysqlerr, ok := err.(*driver.MySQLError)
		if !ok {
			return createdSnack, createdLocation, err
		}
		// Add child rows if exec fails because they are missing.
		if mysqlerr.Number != sqlChildRowForeignKeyError {
			return createdSnack, createdLocation, err
		}
		// Remediate the sqlChildRowForeignKeyError.
		createdSnack = true
		if err := s.CreateSnack(ctx, snackBarcode, ""); err != nil {
			createdSnack = false
			// Move on if it already exists.
			if c := status.Code(err); c != codes.AlreadyExists {
				return createdSnack, createdLocation, err
			}
		}
		createdLocation = true
		if err := s.CreateLocation(ctx, locationName); err != nil {
			createdLocation = false
			// Move on if it already exists
			if c := status.Code(err); c != codes.AlreadyExists {
				return createdSnack, createdLocation, err
			}
		}
		// Now our child rows should exist, so we'll try once more to add the mapping.
		if _, err := s.db.ExecContext(ctx, "INSERT INTO LocationContents (snackBarcode, locationName, numPresent) VALUES(?, ?, 1)", snackBarcode, locationName); err != nil {
			return createdSnack, createdLocation, err
		}
	}
	return createdSnack, createdLocation, nil // false, false, nil
}

// ListContents lists contents of a location.
// If no location is given, lists contents of all locations.
// Return value is locationName:Snack:count.
func (s *SQLImpl) ListContents(ctx context.Context, locationName string) (map[string]map[*sipb.Snack]int, error) {
	retval := map[string]map[*sipb.Snack]int{}
	var query string
	if locationName != "" {
		query = fmt.Sprintf(
			`SELECT lc.snackBarcode,
		 lc.locationName,
		 lc.numPresent,
		 s.name as snackName
		 FROM LocationContents lc
		 LEFT JOIN SnackRegistry s ON lc.snackBarcode = s.barcode
		 WHERE lc.locationName IN (%q)`, locationName)
	} else {
		query = fmt.Sprint(
			`SELECT lc.snackBarcode,
			lc.locationName,
			lc.numPresent,
			s.name as snackName
			FROM LocationContents lc
			LEFT JOIN SnackRegistry s ON lc.snackBarcode = s.barcode`)
	}
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var snackBarcode string
		var locationName string
		var numPresent int
		var snackName string
		if err = rows.Scan(&snackBarcode, &locationName, &numPresent, &snackName); err != nil {
			return nil, err
		}
		locationVal, ok := retval[locationName]
		if !ok {
			locationVal = map[*sipb.Snack]int{}
			retval[locationName] = locationVal
		}

		snack := &sipb.Snack{
			Barcode: snackBarcode,
			Name:    snackName,
		}
		locationVal[snack] = numPresent
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return retval, nil
}
