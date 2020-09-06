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
// This file implements a call to the `ListSnacks` RPC.
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
)

var listSnacksCmd = &cobra.Command{
	Use:   "listsnacks",
	Short: "List all snacks currently registered to SnackInventory.",
	Long:  "List all snacks currently registered to SnackInventory.",
}

func listSnacks(_ *cobra.Command, _ []string) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(connTimeout))
	if err != nil {
		return fmt.Errorf("could not dial %s: %w", address, err)
	}
	defer conn.Close()

	req := &sipb.ListSnacksRequest{}
	client := sipb.NewSnackInventoryClient(conn)

	res, err := client.ListSnacks(context.Background(), req)
	if err != nil {
		return fmt.Errorf("could not list snacks: %v", err)
	}
	fmt.Println("Found snacks:")
	for _, snack := range res.GetSnacks() {
		fmt.Println(snack)
	}
	return nil
}
