package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flags with default values
	host := flag.String("host", "localhost", "Host name or IP address")
	port := flag.String("port", "8080", "Port number")
	
	// Parse the flags
	flag.Parse()

	// Check if a required argument is passed
	if len(flag.Args()) < 1 {
		fmt.Printf("Please provide message to be send.\nCheck --help to see how to specify host and port.\n")
		os.Exit(1)
	}

	// The first non-flag argument is the message
	message := flag.Args()[0]

	// Print the parsed values
	fmt.Printf("Host: %s\n", *host)
	fmt.Printf("Port: %s\n", *port)
	fmt.Printf("Message: %s\n", message)
}

