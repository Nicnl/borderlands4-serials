package byte_mirror

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMirrorLookup_Uint2(t *testing.T) {
	assert.Equal(t, byte(0b00), Uint2Mirror[0b00])
	assert.Equal(t, byte(0b11), Uint2Mirror[0b11])
	assert.Equal(t, byte(0b10), Uint2Mirror[0b01])
	assert.Equal(t, byte(0b01), Uint2Mirror[0b10])
}

func TestMirrorLookup_Uint3(t *testing.T) {
	assert.Equal(t, byte(0b000), Uint3Mirror[0b000])
	assert.Equal(t, byte(0b111), Uint3Mirror[0b111])

	assert.Equal(t, byte(0b100), Uint3Mirror[0b001])
	assert.Equal(t, byte(0b010), Uint3Mirror[0b010])
	assert.Equal(t, byte(0b001), Uint3Mirror[0b100])

	assert.Equal(t, byte(0b101), Uint3Mirror[0b101])
	assert.Equal(t, byte(0b010), Uint3Mirror[0b010])
}

func TestMirrorLookup_Uint4(t *testing.T) {
	assert.Equal(t, byte(0b0000), Uint4Mirror[0b0000])
	assert.Equal(t, byte(0b1111), Uint4Mirror[0b1111])

	assert.Equal(t, byte(0b1000), Uint4Mirror[0b0001])
	assert.Equal(t, byte(0b0100), Uint4Mirror[0b0010])
	assert.Equal(t, byte(0b0010), Uint4Mirror[0b0100])
	assert.Equal(t, byte(0b0001), Uint4Mirror[0b1000])

	assert.Equal(t, byte(0b1010), Uint4Mirror[0b0101])
	assert.Equal(t, byte(0b0101), Uint4Mirror[0b1010])
}

func TestMirrorLookup_Uint5(t *testing.T) {
	assert.Equal(t, byte(0b00000), Uint5Mirror[0b00000])
	assert.Equal(t, byte(0b11111), Uint5Mirror[0b11111])

	assert.Equal(t, byte(0b10000), Uint5Mirror[0b00001])
	assert.Equal(t, byte(0b01000), Uint5Mirror[0b00010])
	assert.Equal(t, byte(0b00100), Uint5Mirror[0b00100])
	assert.Equal(t, byte(0b00010), Uint5Mirror[0b01000])
	assert.Equal(t, byte(0b00001), Uint5Mirror[0b10000])

	assert.Equal(t, byte(0b10101), Uint5Mirror[0b10101])
	assert.Equal(t, byte(0b01010), Uint5Mirror[0b01010])

	assert.Equal(t, byte(0b11011), Uint5Mirror[0b11011])
}

func TestMirrorLookup_Uint7(t *testing.T) {
	assert.Equal(t, byte(0b0000000), Uint7Mirror[0b0000000])
	assert.Equal(t, byte(0b1111111), Uint7Mirror[0b1111111])

	assert.Equal(t, byte(0b1000000), Uint7Mirror[0b0000001])
	assert.Equal(t, byte(0b0100000), Uint7Mirror[0b0000010])
	assert.Equal(t, byte(0b0010000), Uint7Mirror[0b0000100])
	assert.Equal(t, byte(0b0001000), Uint7Mirror[0b0001000])
	assert.Equal(t, byte(0b0000100), Uint7Mirror[0b0010000])
	assert.Equal(t, byte(0b0000010), Uint7Mirror[0b0100000])
	assert.Equal(t, byte(0b0000001), Uint7Mirror[0b1000000])

	assert.Equal(t, byte(0b1010101), Uint7Mirror[0b1010101])
	assert.Equal(t, byte(0b0101010), Uint7Mirror[0b0101010])

	assert.Equal(t, byte(0b1110001), Uint7Mirror[0b1000111])
	assert.Equal(t, byte(0b1000111), Uint7Mirror[0b1110001])
	assert.Equal(t, byte(0b1100110), Uint7Mirror[0b0110011])
	assert.Equal(t, byte(0b0110011), Uint7Mirror[0b1100110])
	assert.Equal(t, byte(0b1010101), Uint7Mirror[0b1010101])

	assert.Equal(t, byte(0b0110101), Uint7Mirror[0b1010110])
	assert.Equal(t, byte(0b0101011), Uint7Mirror[0b1101010])
}

func TestMirrorLookup_Uint8(t *testing.T) {
	assert.Equal(t, byte(0b00000000), Uint8Mirror[0b00000000])
	assert.Equal(t, byte(0b11111111), Uint8Mirror[0b11111111])

	assert.Equal(t, byte(0b10000000), Uint8Mirror[0b00000001])
	assert.Equal(t, byte(0b01000000), Uint8Mirror[0b00000010])
	assert.Equal(t, byte(0b00100000), Uint8Mirror[0b00000100])
	assert.Equal(t, byte(0b00010000), Uint8Mirror[0b00001000])
	assert.Equal(t, byte(0b00001000), Uint8Mirror[0b00010000])
	assert.Equal(t, byte(0b00000100), Uint8Mirror[0b00100000])
	assert.Equal(t, byte(0b00000010), Uint8Mirror[0b01000000])
	assert.Equal(t, byte(0b00000001), Uint8Mirror[0b10000000])

	assert.Equal(t, byte(0b10101010), Uint8Mirror[0b01010101])
	assert.Equal(t, byte(0b01010101), Uint8Mirror[0b10101010])

	assert.Equal(t, byte(0b11110000), Uint8Mirror[0b00001111])
	assert.Equal(t, byte(0b00001111), Uint8Mirror[0b11110000])
	assert.Equal(t, byte(0b11001100), Uint8Mirror[0b00110011])
	assert.Equal(t, byte(0b00110011), Uint8Mirror[0b11001100])
	assert.Equal(t, byte(0b10100101), Uint8Mirror[0b10100101])

	assert.Equal(t, byte(0b01101010), Uint8Mirror[0b01010110])
	assert.Equal(t, byte(0b01010110), Uint8Mirror[0b01101010])

	assert.Equal(t, byte(0b10010101), Uint8Mirror[0b10101001])
	assert.Equal(t, byte(0b10101001), Uint8Mirror[0b10010101])
}

func TestMirrorLookup_Uint11(t *testing.T) {
	assert.Equal(t, uint32(0b00000000000), Uint11Mirror[0b00000000000])
	assert.Equal(t, uint32(0b11111111111), Uint11Mirror[0b11111111111])

	assert.Equal(t, uint32(0b10000000000), Uint11Mirror[0b00000000001])
	assert.Equal(t, uint32(0b01000000000), Uint11Mirror[0b00000000010])
	assert.Equal(t, uint32(0b00100000000), Uint11Mirror[0b00000000100])
	assert.Equal(t, uint32(0b00010000000), Uint11Mirror[0b00000001000])
	assert.Equal(t, uint32(0b00001000000), Uint11Mirror[0b00000010000])
	assert.Equal(t, uint32(0b00000100000), Uint11Mirror[0b00000100000])
	assert.Equal(t, uint32(0b00000010000), Uint11Mirror[0b00001000000])
	assert.Equal(t, uint32(0b00000001000), Uint11Mirror[0b00010000000])
	assert.Equal(t, uint32(0b00000000100), Uint11Mirror[0b00100000000])
	assert.Equal(t, uint32(0b00000000010), Uint11Mirror[0b01000000000])
	assert.Equal(t, uint32(0b00000000001), Uint11Mirror[0b10000000000])

	assert.Equal(t, uint32(0b10101010101), Uint11Mirror[0b10101010101])
	assert.Equal(t, uint32(0b01010101010), Uint11Mirror[0b01010101010])

	assert.Equal(t, uint32(0b11111000000), Uint11Mirror[0b00000011111])
	assert.Equal(t, uint32(0b00000111111), Uint11Mirror[0b11111100000])
	assert.Equal(t, uint32(0b11001100110), Uint11Mirror[0b01100110011])
	assert.Equal(t, uint32(0b01100110011), Uint11Mirror[0b11001100110])
	assert.Equal(t, uint32(0b10100110101), Uint11Mirror[0b10101100101])

	assert.Equal(t, uint32(0b01101010110), Uint11Mirror[0b01101010110])
	assert.Equal(t, uint32(0b01010101100), Uint11Mirror[0b00110101010])

	assert.Equal(t, uint32(0b10010101001), Uint11Mirror[0b10010101001])
	assert.Equal(t, uint32(0b10101010101), Uint11Mirror[0b10101010101])
}

func TestMirrorGeneric(t *testing.T) {
	t.Run("Size_2", func(t *testing.T) {
		assert.Equal(t, uint32(0b00), GenericMirror(0b00, 2))
		assert.Equal(t, uint32(0b11), GenericMirror(0b11, 2))
		assert.Equal(t, uint32(0b10), GenericMirror(0b01, 2))
		assert.Equal(t, uint32(0b01), GenericMirror(0b10, 2))
	})

	t.Run("Size_3", func(t *testing.T) {
		assert.Equal(t, uint32(0b000), GenericMirror(0b000, 3))
		assert.Equal(t, uint32(0b111), GenericMirror(0b111, 3))
		assert.Equal(t, uint32(0b100), GenericMirror(0b001, 3))
		assert.Equal(t, uint32(0b010), GenericMirror(0b010, 3))
		assert.Equal(t, uint32(0b001), GenericMirror(0b100, 3))
	})

	t.Run("Size_5", func(t *testing.T) {
		assert.Equal(t, uint32(0b00000), GenericMirror(0b00000, 5))
		assert.Equal(t, uint32(0b11111), GenericMirror(0b11111, 5))
		assert.Equal(t, uint32(0b10000), GenericMirror(0b00001, 5))
		assert.Equal(t, uint32(0b01000), GenericMirror(0b00010, 5))
		assert.Equal(t, uint32(0b00100), GenericMirror(0b00100, 5))
		assert.Equal(t, uint32(0b00010), GenericMirror(0b01000, 5))
		assert.Equal(t, uint32(0b00001), GenericMirror(0b10000, 5))
	})

	t.Run("Size_8", func(t *testing.T) {
		assert.Equal(t, uint32(0b00000000), GenericMirror(0b00000000, 8))
		assert.Equal(t, uint32(0b11111111), GenericMirror(0b11111111, 8))
		assert.Equal(t, uint32(0b10000000), GenericMirror(0b00000001, 8))
		assert.Equal(t, uint32(0b01000000), GenericMirror(0b00000010, 8))
		assert.Equal(t, uint32(0b00100000), GenericMirror(0b00000100, 8))
		assert.Equal(t, uint32(0b00010000), GenericMirror(0b00001000, 8))
		assert.Equal(t, uint32(0b00001000), GenericMirror(0b00010000, 8))
		assert.Equal(t, uint32(0b00000100), GenericMirror(0b00100000, 8))
		assert.Equal(t, uint32(0b00000010), GenericMirror(0b01000000, 8))
		assert.Equal(t, uint32(0b00000001), GenericMirror(0b10000000, 8))

		assert.Equal(t, uint32(0b10101010), GenericMirror(0b01010101, 8))
		assert.Equal(t, uint32(0b01010101), GenericMirror(0b10101010, 8))
	})

	t.Run("Size_11", func(t *testing.T) {
		assert.Equal(t, uint32(0b00000000000), GenericMirror(0b00000000000, 11))
		assert.Equal(t, uint32(0b11111111111), GenericMirror(0b11111111111, 11))
		assert.Equal(t, uint32(0b10000000000), GenericMirror(0b00000000001, 11))
		assert.Equal(t, uint32(0b01000000000), GenericMirror(0b00000000010, 11))
		assert.Equal(t, uint32(0b00100000000), GenericMirror(0b00000000100, 11))
		assert.Equal(t, uint32(0b00010000000), GenericMirror(0b00000001000, 11))
		assert.Equal(t, uint32(0b00001000000), GenericMirror(0b00000010000, 11))
		assert.Equal(t, uint32(0b00000100000), GenericMirror(0b00000100000, 11))
		assert.Equal(t, uint32(0b00000010000), GenericMirror(0b00001000000, 11))
		assert.Equal(t, uint32(0b00000001000), GenericMirror(0b00010000000, 11))
		assert.Equal(t, uint32(0b00000000100), GenericMirror(0b00100000000, 11))
		assert.Equal(t, uint32(0b00000000010), GenericMirror(0b01000000000, 11))
		assert.Equal(t, uint32(0b00000000001), GenericMirror(0b10000000000, 11))

		assert.Equal(t, uint32(0b10101010101), GenericMirror(0b10101010101, 11))
		assert.Equal(t, uint32(0b01010101010), GenericMirror(0b01010101010, 11))
	})

	t.Run("Size_32", func(t *testing.T) {
		assert.Equal(t, uint32(0b00000000000000000000000000000000), GenericMirror(0b00000000000000000000000000000000, 32))
		assert.Equal(t, uint32(0b11111111111111111111111111111111), GenericMirror(0b11111111111111111111111111111111, 32))
		assert.Equal(t, uint32(0b10000000000000000000000000000000), GenericMirror(0b00000000000000000000000000000001, 32))
		assert.Equal(t, uint32(0b01000000000000000000000000000000), GenericMirror(0b00000000000000000000000000000010, 32))
		assert.Equal(t, uint32(0b00100000000000000000000000000000), GenericMirror(0b00000000000000000000000000000100, 32))
		assert.Equal(t, uint32(0b00010000000000000000000000000000), GenericMirror(0b00000000000000000000000000001000, 32))
		assert.Equal(t, uint32(0b00001000000000000000000000000000), GenericMirror(0b00000000000000000000000000010000, 32))
		assert.Equal(t, uint32(0b00000100000000000000000000000000), GenericMirror(0b00000000000000000000000000100000, 32))
		assert.Equal(t, uint32(0b00000010000000000000000000000000), GenericMirror(0b00000000000000000000000001000000, 32))
		assert.Equal(t, uint32(0b00000001000000000000000000000000), GenericMirror(0b00000000000000000000000010000000, 32))
		assert.Equal(t, uint32(0b00000000100000000000000000000000), GenericMirror(0b00000000000000000000000100000000, 32))
		assert.Equal(t, uint32(0b00000000010000000000000000000000), GenericMirror(0b00000000000000000000001000000000, 32))
		assert.Equal(t, uint32(0b00000000001000000000000000000000), GenericMirror(0b00000000000000000000010000000000, 32))
		assert.Equal(t, uint32(0b00000000000100000000000000000000), GenericMirror(0b00000000000000000000100000000000, 32))
		assert.Equal(t, uint32(0b00000000000010000000000000000000), GenericMirror(0b00000000000000000001000000000000, 32))
		assert.Equal(t, uint32(0b00000000000001000000000000000000), GenericMirror(0b00000000000000000010000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000100000000000000000), GenericMirror(0b00000000000000000100000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000010000000000000000), GenericMirror(0b00000000000000001000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000001000000000000000), GenericMirror(0b00000000000000010000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000100000000000000), GenericMirror(0b00000000000000100000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000010000000000000), GenericMirror(0b00000000000001000000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000001000000000000), GenericMirror(0b00000000000010000000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000000100000000000), GenericMirror(0b00000000000100000000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000000010000000000), GenericMirror(0b00000000001000000000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000000001000000000), GenericMirror(0b00000000010000000000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000000000100000000), GenericMirror(0b00000000100000000000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000000000010000000), GenericMirror(0b00000001000000000000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000000000001000000), GenericMirror(0b00000010000000000000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000000000000100000), GenericMirror(0b00000100000000000000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000000000000010000), GenericMirror(0b00001000000000000000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000000000000001000), GenericMirror(0b00010000000000000000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000000000000000100), GenericMirror(0b00100000000000000000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000000000000000010), GenericMirror(0b01000000000000000000000000000000, 32))
		assert.Equal(t, uint32(0b00000000000000000000000000000001), GenericMirror(0b10000000000000000000000000000000, 32))

		assert.Equal(t, uint32(0b10101010101010101010101010101010), GenericMirror(0b01010101010101010101010101010101, 32))
		assert.Equal(t, uint32(0b01010101010101010101010101010101), GenericMirror(0b10101010101010101010101010101010, 32))
	})
}
