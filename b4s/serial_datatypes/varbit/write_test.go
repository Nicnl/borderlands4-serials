package varbit

import (
	"borderlands_4_serials/lib/bit"
	"borderlands_4_serials/lib/helpers"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite2VarBit(t *testing.T) {
	tests := []struct {
		bin string
		int uint32
		pos int
	}{
		// Special case: the game parser seems to allow a zero bit count
		{"00000", 0, 5},

		{"10000 1", 1, 5 + 1},
		{"01000 01", 2, 5 + 2},
		{"11000 001", 4, 5 + 3},
		{"00100 0001", 8, 5 + 4},
		{"10100 00001", 16, 5 + 5},
		{"01100 000001", 32, 5 + 6},
		{"11100 0000001", 64, 5 + 7},
		{"00010 00000001", 128, 5 + 8},
		{"10010 000000001", 256, 5 + 9},
		{"01010 0000000001", 512, 5 + 10},
		{"11010 00000000001", 1024, 5 + 11},
		{"00110 000000000001", 2048, 5 + 12},
		{"10110 0000000000001", 4096, 5 + 13},
		{"01110 00000000000001", 8192, 5 + 14},
		{"11110 000000000000001", 16384, 5 + 15},
		{"00001 0000000000000001", 32768, 5 + 16},
		{"10001 00000000000000001", 65536, 5 + 17},
		{"01001 000000000000000001", 131072, 5 + 18},
		{"11001 0000000000000000001", 262144, 5 + 19},
		{"00101 00000000000000000001", 524288, 5 + 20},
		{"10101 000000000000000000001", 1048576, 5 + 21},
		{"01101 0000000000000000000001", 2097152, 5 + 22},
		{"11101 00000000000000000000001", 4194304, 5 + 23},
		{"00011 000000000000000000000001", 8388608, 5 + 24},
		{"10011 0000000000000000000000001", 16777216, 5 + 25},
		{"01011 00000000000000000000000001", 33554432, 5 + 26},
		{"11011 000000000000000000000000001", 67108864, 5 + 27},
		{"00111 0000000000000000000000000001", 134217728, 5 + 28},
		{"10111 00000000000000000000000000001", 268435456, 5 + 29},
		{"01111 000000000000000000000000000001", 536870912, 5 + 30},
		{"11111 0000000000000000000000000000001", 1073741824, 5 + 31},
		{"11111 1111111111111111111111111111111", uint32(helpers.IntPow(2, 31)) - 1, 5 + 31},
	}

	for _, test := range tests {
		t.Run(test.bin, func(t *testing.T) {
			// Write the int
			bw := bit.NewWriter()
			fmt.Println(bw)
			Write(bw, test.int)
			assert.Equal(t, test.pos, bw.Pos())
			assert.Equal(t, strings.ReplaceAll(test.bin, " ", ""), bw.String())
		})
	}
}

// Do a full roundtrip, see if writing and reading gives the same result
func TestWriteReadRoundtrip(t *testing.T) {
	for i := 0; i < helpers.IntPow(2, 18); i++ {
		// Write
		bw := bit.NewWriter()
		Write(bw, uint32(i))
		data := bw.Data()

		// Read
		br2 := bit.NewReader(data)
		val, err := Read(br2)
		assert.NoError(t, err)

		// Compare
		assert.Equal(t, uint32(i), val)
		assert.Equal(t, bw.Pos(), br2.Pos())
	}
}
