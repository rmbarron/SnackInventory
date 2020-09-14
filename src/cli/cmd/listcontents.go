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
// This file implements a call to the `ListContents` RPC.
package cmd

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
)

var (
	listContentsLocationName string

	listContentsCmd = &cobra.Command{
		Use:   "listcontents [--flags]",
		Short: "List snacks & count [for locations] in SnackInventory",
		Long: `List snacks & count in SnackInventory.
    If --location is provided, only counts for that location will be found.
    Otherwise, returns counts in all locations.`,
		RunE: listContents,
	}
)

func init() {
	listContentsCmd.Flags().StringVar(&listContentsLocationName, "location", "", "Name of location to list contents. Empty value will query entire contents.")
}

func listContents(_ *cobra.Command, _ []string) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(connTimeout))
	if err != nil {
		return fmt.Errorf("could not dial %s: %w", address, err)
	}
	defer conn.Close()

	client := sipb.NewSnackInventoryClient(conn)
	req := &sipb.ListContentsRequest{
		LocationName: listContentsLocationName,
	}

	res, err := client.ListContents(context.Background(), req)
	if err != nil {
		return fmt.Errorf("could not list contents: %w", err)
	}
	fmt.Println("Found contents:")
	fmt.Println(proto.MarshalTextString(res))
	return nil
}
