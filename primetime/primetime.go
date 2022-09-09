package primetime

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
)

type Request struct {
	Method *string
	Number *float64
}

type Response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeF(f float64) bool {
	if f != float64(int(f)) {
		return false
	} else {
		return isPrime(int(f))
	}
}

func handleRequest(request Request) Response {
	return Response{Method: "isPrime", Prime: isPrimeF(*request.Number)}
}

func writeMalformed(c net.Conn) {
	malformed := Response{Method: "malformed"}
	b, err := json.Marshal(malformed)
	if err != nil {
		log.Print(err)
		return
	}
	c.Write(b)
	c.Write([]byte{'\n'})
}

func HandleFunc(c net.Conn) {
	defer c.Close()
	log.Printf("Serving %s\n", c.RemoteAddr().String())
	reader := bufio.NewReader(c)
	for {
		netData, err := reader.ReadBytes('\n')
		if err != nil {
			log.Print(err)
			return
		}

		var request Request
		err = json.Unmarshal(netData, &request)
		if err != nil {
			log.Print(err)
			writeMalformed(c)
			return
		}

		if request.Method == nil || request.Number == nil {
			log.Printf("Missing field in %+v", request)
			writeMalformed(c)
			return
		}

		if *request.Method != "isPrime" {
			log.Printf("Unrecognised method: '%s'", *request.Method)
			writeMalformed(c)
			return
		}

		response := handleRequest(request)
		b, err := json.Marshal(response)
		if err != nil {
			log.Print(err)
			return
		}
		c.Write(b)
		c.Write([]byte{'\n'})
	}
}
