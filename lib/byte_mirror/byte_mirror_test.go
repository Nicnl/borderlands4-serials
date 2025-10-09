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
