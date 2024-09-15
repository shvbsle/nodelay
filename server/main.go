package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func main() {
	// Define the port to listen on
	port := ":8080"

	// Create an HTTP handler
	http.HandleFunc("/", handleRequest)

	// Get the value for NoDelay (whether to disable Nagle's algorithm)
	useNoDelay := os.Getenv("NO_DELAY") == "true"

	// Start a TCP listener
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error starting listener: %v", err)
	}

	fmt.Printf("Server running on %s with TCP_NODELAY=%v\n", port, useNoDelay)

	for {
		// Accept incoming connections
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Cast to TCPConn to set TCP_NODELAY
		if tcpConn, ok := conn.(*net.TCPConn); ok {
			if useNoDelay {
				fmt.Println("Disabling Nagle's algorithm (TCP_NODELAY)")
				tcpConn.SetNoDelay(true)
			}
			// Set keepalive and keepalive period
			tcpConn.SetKeepAlive(true)
			tcpConn.SetKeepAlivePeriod(10 * time.Second)
		}

		// Serve the connection using the HTTP server
		go http.Serve(&singleUseListener{conn.(*net.TCPConn)}, nil)
	}
}

// singleUseListener wraps *net.TCPConn to allow it to be used as a net.Listener
type singleUseListener struct {
	*net.TCPConn
}

// Accept returns the connection and implements the net.Listener interface
func (ln *singleUseListener) Accept() (net.Conn, error) {
	return ln.TCPConn, nil
}

// Addr returns the network address, required by net.Listener interface
func (ln *singleUseListener) Addr() net.Addr {
	return ln.TCPConn.RemoteAddr()
}
