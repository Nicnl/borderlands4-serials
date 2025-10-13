package bit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	raw  []byte
	bits []byte
}{
	{
		name: "1",
		raw:  []byte{0b00000001},
		bits: []byte{
			0, 0, 0, 0, 0, 0, 0, 1,
		},
	},
	{
		name: "128",
		raw:  []byte{0b10000000},
		bits: []byte{
			1, 0, 0, 0, 0, 0, 0, 0,
		},
	},
	{
		name: "multiblocks",
		raw:  []byte{0b10101010, 0b11001100, 0b11110000},
		bits: []byte{
			1, 0, 1, 0, 1, 0, 1, 0,
			1, 1, 0, 0, 1, 1, 0, 0,
			1, 1, 1, 1, 0, 0, 0, 0,
		},
	},
}

func bit(br *Reader, expected byte, t *testing.T) {
	bit, ok := br.Read()
	if !ok || bit != expected {
		t.Errorf("Expected %d, got %d", expected, bit)
	}
}

func TestBitReaderReadBits(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			br := NewReader(test.raw)
			for _, expected := range test.bits {
				bit(br, expected, t)
			}
		})
	}
}

func TestBitReaderRead4(t *testing.T) {
	br := NewReader([]byte{0b10101010, 0b11001100, 0b11110000})

	val, ok := br.ReadN(4)
	assert.True(t, ok)
	assert.Equal(t, uint32(0b1010), val)

	val, ok = br.ReadN(4)
	assert.True(t, ok)
	assert.Equal(t, uint32(0b1010), val)

	val, ok = br.ReadN(4)
	assert.True(t, ok)
	assert.Equal(t, uint32(0b1100), val)

	val, ok = br.ReadN(4)
	assert.True(t, ok)
	assert.Equal(t, uint32(0b1100), val)

	val, ok = br.ReadN(4)
	assert.True(t, ok)
	assert.Equal(t, uint32(0b1111), val)

	val, ok = br.ReadN(4)
	assert.True(t, ok)
	assert.Equal(t, uint32(0b0000), val)

	// Test out of bounds
	_, ok = br.ReadN(4)
	if ok {
		t.Errorf("Expected false when popping out of bounds")
	}
}
