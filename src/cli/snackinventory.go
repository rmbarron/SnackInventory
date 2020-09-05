// Package main provides a simple CLI for interacting with the SnackInventory
// backend.
//
// This CLI is not currently in a usable state. It exists solely to test changes
// made to the backend server, as the canonical gRPC CLI has no formal release.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
	"google.golang.org/grpc"
)

// Common flags across all subcommands.
var (
	addrFlag = flag.String("address", "localhost:10000",
		"Address to contact SnackInventory backend. Typically a host:port address.")
)

func main() {
	flag.Parse()
	if *addrFlag == "" {
		log.Fatal("--address must be supplied.")
	}

	conn, err := grpc.Dial(*addrFlag, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not dial %s: %v", *addrFlag, err)
	}
	defer conn.Close()

	client := sipb.NewSnackInventoryClient(conn)
	req := &sipb.CreateSnackRequest{
		Snack: &sipb.Snack{
			Barcode: "barcode",
		},
	}

	_, err = client.CreateSnack(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to create snack: %v", err)
	}
	fmt.Println("Successfully created snack!")
}
