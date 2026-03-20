package main

import (
	"log"
)

func main() {
	go func() {
		if err := StartGRPCServer(50051); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()
	if err := StartHTTPServer(18080); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
