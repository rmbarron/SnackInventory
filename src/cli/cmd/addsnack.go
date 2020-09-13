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

// Package cmd provides the various subcommands of the SnackInventory CLI.
// This file implements a call to the `AddSnack` RPC.
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
)

var (
	addSnackBarcode      string
	addSnackLocationName string

	addSnackCmd = &cobra.Command{
		Use:   "addsnack [--flags]",
		Short: "Add a snack to a location in SnackInventory",
		Long: `Add a snack to a location in SnackInventory.
    If a snack or location doesn't exist, creates a minimal entry for missing
    entry.`,
		RunE: addSnack,
	}
)

func init() {
	addSnackCmd.Flags().StringVar(&addSnackBarcode, "snack_barcode", "", "Barcode of snack to add.")
	addSnackCmd.Flags().StringVar(&addSnackLocationName, "location", "", "Name of location to add snack to.")
	addSnackCmd.MarkFlagRequired("snack_barcode")
	addSnackCmd.MarkFlagRequired("location")
}

func addSnack(_ *cobra.Command, _ []string) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(connTimeout))
	if err != nil {
		return fmt.Errorf("could not dial %s: %w", address, err)
	}
	defer conn.Close()

	client := sipb.NewSnackInventoryClient(conn)
	req := &sipb.AddSnackRequest{
		SnackBarcode: addSnackBarcode,
		LocationName: addSnackLocationName,
	}

	if _, err = client.AddSnack(context.Background(), req); err != nil {
		return fmt.Errorf("could not add snack to location: %w", err)
	}
	fmt.Println("Successfully added snack to location!")
	return nil
}
