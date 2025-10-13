package varint

import (
	"borderlands_4_serials/lib/bit"
	"borderlands_4_serials/lib/helpers"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteVarint(t *testing.T) {
	for _, tt := range tests {
		// Skip edge cases that cannot be written, only read
		// (Example: zero with multiple blocks)
		if !tt.includeWriteTest {
			continue
		}

		t.Run(tt.bin, func(t *testing.T) {
			bw := bit.NewWriter()
			fmt.Println(bw.String())

			Write(bw, tt.int)

			expectedData := helpers.BinToBytes(tt.bin)
			assert.Equal(t, expectedData, bw.Data())
			assert.Equal(t, tt.pos, bw.Pos())
		})
	}
}

// Do a full roundtrip, see if writing and reading gives the same result
func TestWriteReadRoundtrip(t *testing.T) {
	for i := 0; i < helpers.IntPow(2, VARINT_MAX_USABLE_BITS); i++ {
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
