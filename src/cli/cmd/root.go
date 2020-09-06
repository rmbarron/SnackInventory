// Package cmd provides the various subcommands of the SnackInventory CLI.
// This file implements the rootCmd for each other file to extend.
package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var (
	address     string
	connTimeout time.Duration

	rootCmd = &cobra.Command{
		Use:   "snackinventory [--address] subcommand [--flags]",
		Short: "A CLI for interacting with the SnackInventory backend.",
		Long: `snackinventory allows for viewing, creating, and removing current
    inventory counts within the SnackInventory backend.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&address, "address", "localhost:10000", "Address to contact SnackInventory backend.")
	rootCmd.PersistentFlags().DurationVar(
		&connTimeout, "dial_timeout", 30*time.Second, "Timeout for connecting to backend.")
	rootCmd.MarkFlagRequired("address")

	rootCmd.AddCommand(createSnackCmd)
}
