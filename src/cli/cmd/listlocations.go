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
// This file implements a call to the `ListLocations` RPC.
package cmd

import (
	"context"
	"fmt"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var listLocationsCmd = &cobra.Command{
	Use:   "listlocations",
	Short: "List all locations currently registered to SnackInventory.",
	Long:  "List all locations currently registered to SnackInventory.",
	RunE:  listLocations,
}

func listLocations(_ *cobra.Command, _ []string) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(connTimeout))
	if err != nil {
		return fmt.Errorf("could not dial %s: %w", address, err)
	}
	defer conn.Close()

	req := &sipb.ListLocationsRequest{}
	client := sipb.NewSnackInventoryClient(conn)

	res, err := client.ListLocations(context.Background(), req)
	if err != nil {
		return fmt.Errorf("could not list locations: %v", err)
	}
	fmt.Println("Found locations:")
	for _, location := range res.GetLocations() {
		fmt.Println(location)
	}
	return nil
}
