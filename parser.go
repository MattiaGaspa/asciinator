package main

import (
	"flag"
	"fmt"
	"os"
)

func parser() (string, int, int) {
	var hostFlag = flag.String("h", "localhost", "Host to run the server on")
	var portFlag = flag.Int("p", 8080, "Port to run the server on")
	var threadsFlag = flag.Int("t", 16, "Number of threads to use for processing")

	if *portFlag < 1 || *portFlag > 65535 {
		fmt.Println("Port must be between 1 and 65535")
		flag.Usage()
		os.Exit(1)
	}

	if *threadsFlag < 1 {
		fmt.Println("Threads must be at least 1")
		flag.Usage()
		os.Exit(1)
	}

	flag.Parse()
	return *hostFlag, *portFlag, *threadsFlag
}

func address(host string, port int) string {
	if host == "" && port == 0 {
		return ""
	}
	if host == "" {
		return fmt.Sprintf(":%d", port)
	}
	return fmt.Sprintf("%s:%d", host, port)
}
