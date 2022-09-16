package smoketest

import (
	"io"
	"testing"

	"github.com/danwhitford/protohackers/testutils"
)

func TestEcho(t *testing.T) {
	s := "hello, world 🏏"
	b := []byte(s)
	testutils.TestClientServer(
		t,
		HandleFunc,
		b,
		b,
	)
}

func FuzzEcho(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		b := []byte(orig)
		res, err := testutils.RunClientServer(HandleFunc, b, len(b))

		if err != nil {
			if err != io.EOF {
				t.Fatal(err)
			}
		}
		if !Compare(b, res) {
			t.Fatalf("Not equal:\n%v\n%v", b, res)
		}
	})
}

func Compare(a, b []byte) bool {
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
