package budgetchat

import (
	"log"
	"net"
)

func ChatServer() {
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
		go func(c net.Conn) {
			c.Write([]byte("fuck off\n"))
			c.Close()
		}(conn)
	}
}
