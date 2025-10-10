package byte_mirror

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestMirrorLookup_Uint3(t *testing.T) {
	assert.Equal(t, byte(0b000), Uint3Mirror[0b000])
	assert.Equal(t, byte(0b111), Uint3Mirror[0b111])

	assert.Equal(t, byte(0b100), Uint3Mirror[0b001])
	assert.Equal(t, byte(0b010), Uint3Mirror[0b010])
	assert.Equal(t, byte(0b001), Uint3Mirror[0b100])

	assert.Equal(t, byte(0b101), Uint3Mirror[0b101])
	assert.Equal(t, byte(0b010), Uint3Mirror[0b010])
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
