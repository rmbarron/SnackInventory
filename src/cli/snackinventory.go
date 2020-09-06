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
