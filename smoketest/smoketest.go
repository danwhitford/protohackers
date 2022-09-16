package smoketest

import (
	"io"
	"log"
	"net"
	"time"
)

func HandleFunc(c net.Conn) {
	defer c.Close()

	log.Printf("Got connection from %v", c.RemoteAddr())
	// Echo all incoming data.
	c.SetReadDeadline(time.Now().Add(time.Second))
	written, err := io.Copy(c, c)
	if written == 0 {
		log.Println("Didn't write anything")		
	}
	if err != nil {
		log.Println(err)
	}
	log.Printf("Wrote %d bytes to %v\n", written, c.RemoteAddr())	
}
