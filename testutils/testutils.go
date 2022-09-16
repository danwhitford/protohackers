package testutils

import (
	"net"
	"testing"
)

func TestClientServer(t *testing.T, handler func(c net.Conn), input []byte, expected []byte) {
	TestClientServerBursty(t, handler, [][]byte{input}, expected)
}

func TestClientServerBursty(t *testing.T, handler func(c net.Conn), input [][]byte, expected []byte) {
	res, err := RunClientServer(handler, input, len(expected))
	if err != nil {
		t.Fatal(err)
	}

	if !Equals(res, expected) {
		t.Fatalf("Expected did not match actual\nExpected:%v\nActual  :%v\n\n", expected, res)
	}
}

func RunClientServer(handler func(c net.Conn), input [][]byte, l int) ([]byte, error) {
	server, client := net.Pipe()

	go handler(server)
	for _, b := range input {
		_, err := client.Write(b)
		if err != nil {
			return nil, err
		}
	}

	b := make([]byte, l)
	_, err := client.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Equals(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
