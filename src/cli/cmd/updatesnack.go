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
// This file implements a call to the `UpdateSnack` RPC.
package cmd

import (
	"context"
	"fmt"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	updateSnackBarcode string
	updateSnackName    string

	updateSnackCmd = &cobra.Command{
		Use:   "updatesnack [--flags]",
		Short: "Update a snack in SnackInventory",
		Long: `Update a snack in SnackInventory.
    All field values are written as given. To avoid accidentally
    overriding fields with empty values, read the snack you intend to
    update first.
    --barcode is required to find the snack to be updated.`,
		RunE: updateSnack,
	}
)

func init() {
	updateSnackCmd.Flags().StringVar(
		&updateSnackBarcode, "barcode", "", "barcode of snack to update in SnackInventory.")
	updateSnackCmd.Flags().StringVar(
		&updateSnackName, "name", "", "name of snack to update in SnackInventory.")
	updateSnackCmd.MarkFlagRequired("barcode")
}

func updateSnack(_ *cobra.Command, _ []string) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(connTimeout))
	if err != nil {
		return fmt.Errorf("could not dial %s: %w", address, err)
	}
	defer conn.Close()

	client := sipb.NewSnackInventoryClient(conn)
	req := &sipb.UpdateSnackRequest{
		Snack: &sipb.Snack{
			Barcode: updateSnackBarcode,
			Name:    updateSnackName,
		},
	}

	if _, err = client.UpdateSnack(context.Background(), req); err != nil {
		return fmt.Errorf("could not update snack: %w", err)
	}
	fmt.Println("Successfully updated snack!")
	return nil
}
