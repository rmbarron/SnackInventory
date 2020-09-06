// Package main provides a simple CLI for interacting with the SnackInventory
// backend.
//
// This is a cobra-powered CLI. Each different RPC to SnackInventory is
// represented in its own thick-client file within package cmd. Each file
// defines a `func (cmd *cobra.Command, args []string) error` function that
// does the actual work of connecting the backend and performing some action.
package main

import (
	"log"

	"github.com/rmbarron/SnackInventory/src/cli/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("operation failed, encountered error: %v", err)
	}
}
