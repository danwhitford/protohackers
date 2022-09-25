package main

import (
	"log"
	"net"
	"os"

	"github.com/danwhitford/protohackers/meanstoanend"
	"github.com/danwhitford/protohackers/primetime"
	"github.com/danwhitford/protohackers/smoketest"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Need to give command")
	}

	var handlerFunc func(c net.Conn)
	funcs := map[string]func(c net.Conn){
		"smoketest":    smoketest.HandleFunc,
		"primetime":    primetime.HandleFunc,
		"meanstoanend": meanstoanend.HandleFunc,
	}
	cmd := os.Args[1]
	handlerFunc, prs := funcs[cmd]
	if !prs {
		log.Fatalf("Command not recognised: '%s'\n", cmd)
	}

	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", "0.0.0.0:2000")
	log.Printf("Listening on:  %v", l.Addr())
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go handlerFunc(conn)
	}
}
