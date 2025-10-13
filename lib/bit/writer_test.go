package bit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test data for writing bits
func TestBitWriterWriteBits(t *testing.T) {
	for _, test := range testsBits {
		t.Run(test.name, func(t *testing.T) {
			bw := NewWriter()
			fmt.Println(bw)
			bw.WriteBits(test.bits...)
			fmt.Println(bw)
			assert.Equal(t, test.bytes, bw.data)
		})
	}
}

// Check if writing bits at specific positions works correctly
func TestBitWriterWriteBitsAtPos(t *testing.T) {
	bw := NewWriter()
	fmt.Println(bw)
	assert.Equal(t, []byte{}, bw.data)

	bw.pos = 20
	bw.WriteBits(1, 0, 1)
	fmt.Println(bw)
	assert.Equal(t, []byte{0b00000000, 0b00000000, 0b00001010}, bw.data)

	bw.pos = 10
	bw.WriteBits(1, 1)
	fmt.Println(bw)
	assert.Equal(t, []byte{0b00000000, 0b00110000, 0b00001010}, bw.data)

	bw.pos = 22
	bw.WriteBits(0)
	fmt.Println(bw)
	assert.Equal(t, []byte{0b00000000, 0b00110000, 0b00001000}, bw.data)
}

// Test writing bits that do not align to byte boundaries
func TestBitWriterWriteBitsPartial(t *testing.T) {
	bw := NewWriter()
	fmt.Println(bw)
	assert.Equal(t, []byte{}, bw.data)

	bw.WriteBits(1, 0, 1)
	fmt.Println(bw)
	assert.Equal(t, []byte{0b10100000}, bw.data)

	bw.WriteBits(1, 1)
	fmt.Println(bw)
	assert.Equal(t, []byte{0b10111000}, bw.data)

	bw.WriteBits(0, 0, 1)
	fmt.Println(bw)
	assert.Equal(t, []byte{0b10111001}, bw.data)

	bw.WriteBits(0)
	fmt.Println(bw)
	assert.Equal(t, []byte{0b10111001, 0b00000000}, bw.data)

	bw.WriteBits(0, 0, 0, 0, 0, 0, 1)
	fmt.Println(bw)
	assert.Equal(t, []byte{0b10111001, 0b00000001}, bw.data)
}
