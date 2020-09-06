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
