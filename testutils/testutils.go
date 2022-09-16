package testutils

import (
	"net"
	"testing"
)

func TestClientServer(t *testing.T, handler func(c net.Conn), input []byte, expected []byte) {
	TestClientServerBursty(t, handler, [][]byte{input}, expected)
}

func TestClientServerBursty(t *testing.T, handler func(c net.Conn), input [][]byte, expected []byte) {
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

func RunClientServer(handler func(c net.Conn), input []byte, l int) ([]byte, error) {
	server, client := net.Pipe()

	go handler(server)
	_, err := client.Write(input)
	if err != nil {
		return nil, err
	}
	b := make([]byte, l)
	_, err = client.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
