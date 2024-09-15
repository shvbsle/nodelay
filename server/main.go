package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})
	useNoDelay := os.Getenv("NO_DELAY") == "true"
	port := ":8080"

	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error starting listener: %v", err)
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Printf("Error accepting connection: %v", err)
				continue
			}

			// Cast to TCPConn to set TCP_NODELAY
			if tcpConn, ok := conn.(*net.TCPConn); ok {
				if useNoDelay {
					fmt.Println("Disabling Nagle's algorithm (TCP_NODELAY)")
					tcpConn.SetNoDelay(true) // Disable Nagle's algorithm
				} else {
					fmt.Println("Using Nagle's algorithm (TCP_NODELAY disabled)")
					tcpConn.SetNoDelay(false) // Enable Nagle's algorithm
				}
			}

			conn.Close()
		}
	}()

	fmt.Println("Server running on port", port)
	log.Fatal(app.Listener(ln))
}
