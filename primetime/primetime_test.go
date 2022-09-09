package primetime

import "testing"

func makeRequest(s string, n float64) Request {
	return Request{Method: &s, Number: &n}
}

func TestHandleRequest(t *testing.T) {
	req := makeRequest("isPrime", 123)
	res := handleRequest(req)
	if res.Method != "isPrime" {
		t.Fatalf("Method wrong")
	}
	if res.Prime != false {
		t.Fatalf("Result wrong")
	}
}

func TestSomePrimes(t *testing.T) {
	primes := []float64{2, 3, 5, 7, 11, 13, 17, 19}
	for _, p := range primes {
		req := makeRequest("isPrime", p)
		res := handleRequest(req)
		if res.Method != "isPrime" {
			t.Fatalf("Method wrong")
		}
		if res.Prime != true {
			t.Fatalf("Result wrong for %+v", req)
		}
	}
}

func TestSomeNonePrimes(t *testing.T) {
	nonprimes := []float64{1, 4, 6, 8, 9, 10, 15, 16, 18}
	for _, p := range nonprimes {
		req := makeRequest("isPrime", p)
		res := handleRequest(req)
		if res.Method != "isPrime" {
			t.Fatalf("Method wrong")
		}
		if res.Prime != false {
			t.Fatalf("Result wrong for %+v", req)
		}
	}
}

func TestSomeNoneFloats(t *testing.T) {
	nonprimes := []float64{1.4, 4.2, 6.4, 8.7, 2.1, 7.3, 19.5}
	for _, p := range nonprimes {
		req := makeRequest("isPrime", p)
		res := handleRequest(req)
		if res.Method != "isPrime" {
			t.Fatalf("Method wrong")
		}
		if res.Prime != false {
			t.Fatalf("Result wrong for %+v", req)
		}
	}
}
