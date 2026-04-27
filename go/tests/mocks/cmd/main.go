package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/saichler/l8spring/go/tests/mocks"
)

func main() {
	address := flag.String("address", "https://localhost:2773", "Server address")
	user := flag.String("user", "admin", "Username")
	password := flag.String("password", "admin", "Password")
	insecure := flag.Bool("insecure", false, "Skip TLS verification")
	flag.Parse()

	if err := mocks.RunMockGenerator(*address, *user, *password, *insecure); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
