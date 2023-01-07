package main

import (
	"CLI-Learn/Connections"
	"bufio"
	"flag"
	"fmt"
	"os"
)

func Connect(args []string) {
	// Define the base required flags
	connectionType := flag.String("t", "", "connection type (required)")
	flag.Parse()

	// Check if the connection type flag is set
	if *connectionType == "" {
		fmt.Println("Error: missing connection type flag (-t)")
		return
	}

	// Define a map of additional required flags for each connection type
	requiredFlags := map[string][]string{
		"tcp":    []string{"a", "p"},
		"rshell": []string{"l"},
	}

	// Check if the required flags for the specified connection type are set
	for _, flagName := range requiredFlags[*connectionType] {
		if flag.Lookup(flagName) == nil {
			// Prompt the user for the missing required flag
			fmt.Printf("Enter value for required flag (-%s): ", flagName)
			reader := bufio.NewReader(os.Stdin)
			flagValue, _ := reader.ReadString('\n')
			flag.Set(flagName, flagValue)
		}
	}

	var conn Connections.Connection
	var err error

	// Establish the connection concurrently
	done := make(chan error)
	go func() {
		switch *connectionType {
		case "tcp":
			// Create a ConnectTCP struct and establish the connection
			conn = &Connections.ConnectTCP{
				Name:    "TCP Connection",
				Address: flag.Lookup("a").Value.String(),
				Port:    flag.Lookup("p").Value.String(),
			}
			done <- conn.Connect()
		default:
			done <- fmt.Errorf("invalid connection type '%s'", *connectionType)
		}
	}()

	// Wait for the connection to be established
	err = <-done

	if err != nil {
		fmt.Println("Error establishing connection:", err)
		return
	}

	fmt.Println("Connection established successfully")
}
