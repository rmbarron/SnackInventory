SnackInventory is a lightweight inventory system used to track current snack
counts in your home.

# Setup

SnackInventory is a Golang gRPC service. Setup requirements are mostly that
Golang, protoc, and gRPC genrules are installed. Below instructions assume
Golang is already installed for the system.

## Linux / Ubuntu

Install Cobra, which powers the CLI for interacting with the backend:
*  `go get -u github.com/spf13/cobra/cobra`

Install protoc genrules for Go & gRPC:
*  `go get google.golang.org/protobuf/cmd/protoc-gen-go`
*  `go install google.golang.org/protobuf/cmd/protoc-gen-go`
*  `go get google.golang.org/grpc/cmd/protoc-gen-go-grpc`
*  `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`
*  `go get google.golang.org/grpc`
*  `go install google.golang.org/grpc`
*  `sudo apt-get install protobuf-compiler`

With those installed, everything should be compilable.
*  `git clone $SnackInventoryURL`
*  `cd SnackInventory`
*  `go get ./...`

Alternatively, the package should be installable via:
*  `go get github.com/rmbarron/SnackInventory`
*  `go install github.com/rmbarron/SnackInventory`

but that flow has not been tested yet.

### Recompiling Protos

Any changes to messages / RPCs require recompiling the generated proto code.
Currently, this requires cloning into a `github.com/$USER/SnackInventory/` dir.
Then, from the root dir (containing `githug.com`), run:
*  `protoc -I=github.com/$USER/SnackInventory/src/proto/snackinventory/ --go_out=./ --go-grpc_out=./ ./github.com/$USER/SnackInventory/src/proto/snackinventory/snackinventory.proto`
