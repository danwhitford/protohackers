package main

import (
	"log"
	"net"

	"github.com/danwhitford/protohackers/meanstoanend"
)

func main() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", "0.0.0.0:2000")
	log.Printf("Listening on:  %v", l.Addr())
	if err != nil {
		log.Print(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()		
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go meanstoanend.HandleFunc(conn)
	}
}
