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
	teststrings := []string{"Hello, world🏏", " ", "!12345"}
	for _, ts := range teststrings {
		f.Add([]byte(ts))
	}
	f.Fuzz(func(t *testing.T, b []byte) {
		res, err := testutils.RunClientServer(HandleFunc, [][]byte{b}, len(b))

		if err != nil {
			if err != io.EOF {
				t.Fatal(err)
			}
		}
		if !testutils.Equals(b, res) {
			t.Fatalf("Not equal:\n%v\n%v", b, res)
		}
	})
}

