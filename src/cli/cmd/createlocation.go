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
	createLocationName string

	createLocationCmd = &cobra.Command{
		Use:   "createlocation [--flags]",
		Short: "Create a new location in SnackInventory",
		Long: `Creates a new location in SnackInventory.
    --name is required, as that is the unique identifier for each location.`,
		RunE: createLocation,
	}
)

func init() {
	createLocationCmd.Flags().StringVar(&createLocationName, "name", "",
		"Name of location to add to SnackInventory")
	createLocationCmd.MarkFlagRequired("name")
}

func createLocation(_ *cobra.Command, _ []string) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(connTimeout))
	if err != nil {
		return fmt.Errorf("could not dial %s: %w", address, err)
	}
	defer conn.Close()

	client := sipb.NewSnackInventoryClient(conn)
	req := &sipb.CreateLocationRequest{
		Location: &sipb.Location{
			Name: createLocationName,
		},
	}

	if _, err = client.CreateLocation(context.Background(), req); err != nil {
		return fmt.Errorf("could not create snack: %w", err)
	}
	fmt.Println("Successfully create location!")
	return nil
}
