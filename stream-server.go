package main

import (
	"fmt"
	"net/http"
)

func main() {
	// configure stream directory name and port
	const streamDir = "stream"
	const port = 8080

	// add a handler for the stream files
	http.Handle("/", addHeaders(http.FileServer(http.Dir(streamDir))))
	fmt.Printf("Starting server on %v\n", streamDir, port)
}
