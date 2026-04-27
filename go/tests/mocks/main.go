package mocks

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func RunMockGenerator(address, user, password string, insecure bool) error {
	fmt.Println("Spring Mock Data Generator")
	fmt.Printf("  Server: %s\n", address)
	fmt.Printf("  User:   %s\n", user)

	transport := &http.Transport{}
	if insecure {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	httpClient := &http.Client{Transport: transport}

	client := NewSpringClient(address, httpClient)

	fmt.Println("Authenticating...")
	if err := client.Authenticate(user, password); err != nil {
		return fmt.Errorf("authentication failed: %w", err)
	}
	fmt.Println("Authentication successful.")

	store := &MockDataStore{}

	fmt.Println("\nGenerating mock data...")
	if err := RunAllPhases(client, store); err != nil {
		return fmt.Errorf("mock data generation failed: %w", err)
	}

	PrintSummary(store)
	return nil
}
