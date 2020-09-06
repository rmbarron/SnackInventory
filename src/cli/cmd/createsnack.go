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
// This file implements a call to the `CreateSnack` RPC.
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
)

var (
	createSnackBarcode string
	createSnackName    string
	createSnackBrand   string

	createSnackCmd = &cobra.Command{
		Use:   "createsnack [--flags]",
		Short: "Create a new snack in SnackInventory.",
		Long: `Creates a new snack in SnackInventory.
    --barcode is required, as that is the unique identifier for snacks.`,
		RunE: createSnack,
	}
)

func init() {
	createSnackCmd.Flags().StringVar(&createSnackBarcode, "barcode", "", "Barcode of the snack to create.")
	createSnackCmd.Flags().StringVar(&createSnackName, "name", "", "Name of the snack to create.")
	createSnackCmd.Flags().StringVar(&createSnackBrand, "brand", "", "Brand of the snack to create.")
	createSnackCmd.MarkFlagRequired("barcode")
}

func createSnack(_ *cobra.Command, _ []string) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(connTimeout))
	if err != nil {
		return fmt.Errorf("could not dial %s: %w", address, err)
	}
	defer conn.Close()

	client := sipb.NewSnackInventoryClient(conn)
	req := &sipb.CreateSnackRequest{
		Snack: &sipb.Snack{
			Barcode: createSnackBarcode,
			Name:    createSnackName,
			Brand:   createSnackBrand,
		},
	}

	_, err = client.CreateSnack(context.Background(), req)
	if err != nil {
		return fmt.Errorf("could not create snack: %w", err)
	}
	fmt.Println("Successfully created snack!")
	return nil
}
