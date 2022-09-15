package smoketest

import (
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
