SnackInventory is a lightweight inventory system used to track current snack
counts in your home.

# Overview

The idea is to keep track of snacks as they enter and exit your cupboard. This
project consists of several pieces:
*  Backend gRPC service that acts as an interface above storage.
*  CLI for interacting with the backend easily.
*  A web UI for browsing the current state of the backend, and updating values as appropriate.
*  A raspberry Pi daemon, that listens to a barcode scanner and increments inventory count.

![SnackInventory architecture](https://docs.google.com/drawings/d/e/2PACX-1vSPKeEJsa81ATaFLAuXv7vw80L45y5H_UN7CoHQZ9jUj7CrBWFbGfwEz3F5Z2QnPFeh6z-bjebO-JAL/pub?w=960&h=312)

On initial scan, barcodes need their metadata (name & brand) filled in manually.
This is cumbersome, but seamless on repeat orders of the same product.

# Server Usage

The backend server acts as an interface between the user (CLI & UI) & storage.
On launch, flags specific to the used storage interface must be given. Passwords
are read from stdin to avoid having '--password' in shell history.

Ex: `go run src/backend/server/server.go --sql_user=$USER --sql_address=127.0.0.1:3306 < ~/sql_pass.txt`

# Storage Model

The primary backend for the SnackInventory server is SQL. When a SQL
implementation is used, an arbitrary database name can be given. Inside that
database, tables "SnackRegistry" and "Locations" are assumed present.

## Schema

SnackRegistry: barcode VARCHAR(20) PRIMARY KEY, name VARCHAR(255)

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


### MariaDB

SnackInventory primarily uses MariaDB as its backing DB. To set up a serving
instance:

*  `sudo apt install mariadb-server`
*  `sudo mysql_secure_installation` - Locks down core vulnerabilities (like root
   user permissions)
*  `sudo mysql` - enter interactive DB shell for setup
  *  `CREATE DATABASE SnackInventory;`
  *  `USE SnackInventory;`
  *  `CREATE TABLE SnackRegistry ( barcode VARCHAR(20) PRIMARY KEY, name VARCHAR(255));`
  *  `GRANT ALL PRIVILEGES ON SnackInventory.* TO '$USER'@'$NETWORK' IDENTIFIED BY '$PASSWORD' WITH GRANT OPTION;`
  *  `FLUSH PRIVILEGES;`

To allow remote connections (rather than just localhost connections):

*  Comment out `bind-address = 127.0.0.1` in `/etc/mysql/mariadb.conf.d/50-server.cnf`
*  Restart the mariadb service: `sudo systemctl restart mariadb.service`


### Recompiling Protos

Any changes to messages / RPCs require recompiling the generated proto code.
Currently, this requires cloning into a `github.com/$USER/SnackInventory/` dir.
Then, from the root dir (containing `githug.com`), run:
*  `protoc -I=github.com/$USER/SnackInventory/src/proto/snackinventory/ --go_out=./ --go-grpc_out=./ ./github.com/$USER/SnackInventory/src/proto/snackinventory/snackinventory.proto`
