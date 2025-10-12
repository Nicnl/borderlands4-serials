package serial_tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadVarintNormal(t *testing.T) {
	var tests = []struct {
		bin      string
		expected uint32
		pos      int
	}{
		{"10000", 1, 5},
		{"01000", 2, 5},
		{"00100", 4, 5},
		{"00010", 8, 5},

		{"11110", 15, 5},
		{"01110", 14, 5},
		{"00110", 12, 5},

		{"00001 10000", 16, 10},
		{"00001 01000", 32, 10},
		{"00001 00100", 64, 10},
		{"00001 00010", 128, 10},

		{"00001 00001 10000", 256, 15},
		{"00001 00001 01000", 512, 15},
		{"00001 00001 00100", 1024, 15},
		{"00001 00001 00010", 2048, 15},

		{"00001 00001 00001 10000", 4096, 20},
		{"00001 00001 00001 01000", 8192, 20},
		{"00001 00001 00001 00100", 16384, 20},
		{"00001 00001 00001 00010", 32768, 20},

		// All zero
		{"00000", 0, 5},
		{"00001 00000", 0, 10},
		{"00001 00001 00000", 0, 15},
		{"00001 00001 00001 00000", 0, 20},
	}

	for _, tt := range tests {
		t.Run(tt.bin, func(t *testing.T) {
			data := binToBytes(tt.bin)
			tok := NewTokenizer(data)
			val, err := tok.readVarint()
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, val)
			assert.Equal(t, tt.pos, tok.bs.Pos())
		})
	}
}
