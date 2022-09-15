package testutils

import (
	"net"
	"testing"
)

func TestClientServer(t *testing.T, handler func(c net.Conn), input []byte, expected []byte) {
	server, client := net.Pipe()
	
	go handler(server)		
	client.Write(input)
	
	b := make([]byte, len(expected))
	_, err := client.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	for i := range expected {
		if expected[i] != b[i] {
			t.Fatalf("Expected did not match actual\nExepcted: %v\nActual: %v\n\n", expected, b)
		}
	}	
}

func TestBrokenSending(t *testing.T, handler func(c net.Conn), input [][]byte, expected []byte) {
	server, client := net.Pipe()
	
	go handler(server)
	for _, b := range input {
		client.Write(b)
	}
	
	b := make([]byte, len(expected))
	_, err := client.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	for i := range expected {
		if expected[i] != b[i] {
			t.Fatalf("Expected did not match actual\nExepcted: %v\nActual: %v\n\n", expected, b)
		}
	}	
}
