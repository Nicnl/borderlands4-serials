package bit_reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func bit(br *BitReader, expected byte, t *testing.T) {
	bit, ok := br.Read()
	if !ok || bit != expected {
		t.Errorf("Expected %d, got %d", expected, bit)
	}
}

func TestBitReader1(t *testing.T) {
	br := NewBitReader([]byte{0b10101010, 0b11001100})

	bit(br, 1, t)
	bit(br, 0, t)
	bit(br, 1, t)
	bit(br, 0, t)
	bit(br, 1, t)
	bit(br, 0, t)
	bit(br, 1, t)
	bit(br, 0, t)

	bit(br, 1, t)
	bit(br, 1, t)
	bit(br, 0, t)
	bit(br, 0, t)
	bit(br, 1, t)
	bit(br, 1, t)
	bit(br, 0, t)
	bit(br, 0, t)

	// Test out of bounds
	_, ok := br.Read()
	if ok {
		t.Errorf("Expected false when popping out of bounds")
	}
}

func TestBitReaderRead4(t *testing.T) {
	br := NewBitReader([]byte{0b10101010, 0b11001100, 0b11110000})

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
