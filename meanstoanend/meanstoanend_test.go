package meanstoanend

import (
	"testing"

	"github.com/danwhitford/protohackers/testutils"
)

func TestNoPrices(t *testing.T) {
	testutils.TestClientServer(
		t,
		HandleFunc,
		[]byte{'Q', 0, 0, 0, 0, 255, 255, 255, 255},
		[]byte{0, 0, 0, 0},
	)
}

func TestSingleInsert(t *testing.T) {
	testutils.TestClientServer(
		t,
		HandleFunc,
		[]byte{
			'I', 0x00, 0x00, 0x30, 0x39, 0x00, 0x00, 0x00, 0x65,
			'Q', 0x00, 0x00, 0x30, 0x38, 0x00, 0x00, 0x30, 0x40},
		[]byte{0x00, 0x00, 0x00, 0x65},
	)
}

func TestBookExample(t *testing.T) {
	testutils.TestClientServer(
		t,
		HandleFunc,
		[]byte{
			0x49, 0x00, 0x00, 0x30, 0x39, 0x00, 0x00, 0x00, 0x65,
			0x49, 0x00, 0x00, 0x30, 0x3a, 0x00, 0x00, 0x00, 0x66,
			0x49, 0x00, 0x00, 0x30, 0x3b, 0x00, 0x00, 0x00, 0x64,
			0x49, 0x00, 0x00, 0xa0, 0x00, 0x00, 0x00, 0x00, 0x05,
			0x51, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x40, 0x00,
		},
		[]byte{0x00, 0x00, 0x00, 0x65},
	)
}

func TestBrokenSending(t *testing.T) {
	testutils.TestBrokenSending(
		t,
		HandleFunc,
		[][]byte{
			{0x49},
			{0x00, 0x00, 0x30, 0x39, 0x00, 0x00, 0x00, 0x65,
				0x49, 0x00, 0x00, 0x30, 0x3a, 0x00, 0x00, 0x00, 0x66,
				0x49, 0x00, 0x00, 0x30, 0x3b, 0x00, 0x00, 0x00, 0x64,
				0x49, 0x00, 0x00, 0xa0, 0x00, 0x00, 0x00, 0x00, 0x05,
				0x51, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x40, 0x00,
			}},
		[]byte{0x00, 0x00, 0x00, 0x65},
	)
}