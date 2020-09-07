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
// This file implements a call to the `DeleteSnack` RPC.
package cmd

import (
	"context"
	"fmt"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	deleteSnackBarcode string

	deleteSnackCmd = &cobra.Command{
		Use:   "deletesnack [--flags]",
		Short: "Delete a snack from SnackInventory.",
		Long: `Delete a snack from SnackInventory.
    --barcode is required to only delete the snack by its unique ID.`,
		RunE: deleteSnack,
	}
)

func init() {
	deleteSnackCmd.Flags().StringVar(
		&deleteSnackBarcode, "barcode", "", "barcode of snack to delete.")
	deleteSnackCmd.MarkFlagRequired("barcode")
}

func deleteSnack(_ *cobra.Command, _ []string) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(connTimeout))
	if err != nil {
		return fmt.Errorf("could not dial %s: %w", address, err)
	}
	defer conn.Close()

	client := sipb.NewSnackInventoryClient(conn)
	req := &sipb.DeleteSnackRequest{
		Barcode: deleteSnackBarcode,
	}

	if _, err = client.DeleteSnack(context.Background(), req); err != nil {
		return fmt.Errorf("could not delete snack: %w", err)
	}
	fmt.Println("Successfully deleted snack!")
	return nil
}
