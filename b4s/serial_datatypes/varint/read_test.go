package varint

import (
	"borderlands_4_serials/lib/bit"
	"borderlands_4_serials/lib/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	bin              string
	int              uint32
	pos              int
	includeWriteTest bool // edge case that cannot be written, only read
}{
	{"10000", 1, 5, true},
	{"01000", 2, 5, true},
	{"00100", 4, 5, true},
	{"00010", 8, 5, true},

	{"11110", 15, 5, true},
	{"01110", 14, 5, true},
	{"00110", 12, 5, true},

	{"00001 10000", 16, 10, true},
	{"00001 01000", 32, 10, true},
	{"00001 00100", 64, 10, true},
	{"00001 00010", 128, 10, true},

	{"00001 00001 10000", 256, 15, true},
	{"00001 00001 01000", 512, 15, true},
	{"00001 00001 00100", 1024, 15, true},
	{"00001 00001 00010", 2048, 15, true},

	{"00001 00001 00001 10000", 4096, 20, true},
	{"00001 00001 00001 01000", 8192, 20, true},
	{"00001 00001 00001 00100", 16384, 20, true},
	{"00001 00001 00001 00010", 32768, 20, true},

	// All zero
	{"00000", 0, 5, true},
	{"00001 00000", 0, 10, false},
	{"00001 00001 00000", 0, 15, false},
	{"00001 00001 00001 00000", 0, 20, false},
}

func TestReadVarint(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.bin, func(t *testing.T) {
			data := helpers.BinToBytes(tt.bin)
			br := bit.NewReader(data)

			val, err := Read(br)
			assert.NoError(t, err)
			assert.Equal(t, tt.int, val)
			assert.Equal(t, tt.pos, br.Pos())
		})
	}
}
