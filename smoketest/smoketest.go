package smoketest

import (
	"io"
	"log"
	"net"
)

func HandleFunc(c net.Conn) {
	log.Printf("Got connection from %v", c.RemoteAddr())
	// Echo all incoming data.
	written, err := io.Copy(c, c)
	if err != nil {
		log.Print(err)
	}
	log.Printf("Wrote %d bytes to %v", written, c.RemoteAddr())
	// Shut down the connection.
	c.Close()
}