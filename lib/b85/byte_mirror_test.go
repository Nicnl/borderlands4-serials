package b85

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMirrorLookup_Simple(t *testing.T) {
	assert.Equal(t, byte(0b00000000), mirrorLookup[0b00000000])
	assert.Equal(t, byte(0b11111111), mirrorLookup[0b11111111])

	assert.Equal(t, byte(0b10000000), mirrorLookup[0b00000001])
	assert.Equal(t, byte(0b01000000), mirrorLookup[0b00000010])
	assert.Equal(t, byte(0b00100000), mirrorLookup[0b00000100])
	assert.Equal(t, byte(0b00010000), mirrorLookup[0b00001000])
	assert.Equal(t, byte(0b00001000), mirrorLookup[0b00010000])
	assert.Equal(t, byte(0b00000100), mirrorLookup[0b00100000])
	assert.Equal(t, byte(0b00000010), mirrorLookup[0b01000000])
	assert.Equal(t, byte(0b00000001), mirrorLookup[0b10000000])

	assert.Equal(t, byte(0b10101010), mirrorLookup[0b01010101])
	assert.Equal(t, byte(0b01010101), mirrorLookup[0b10101010])

	assert.Equal(t, byte(0b11110000), mirrorLookup[0b00001111])
	assert.Equal(t, byte(0b00001111), mirrorLookup[0b11110000])
	assert.Equal(t, byte(0b11001100), mirrorLookup[0b00110011])
	assert.Equal(t, byte(0b00110011), mirrorLookup[0b11001100])
	assert.Equal(t, byte(0b10100101), mirrorLookup[0b10100101])

	assert.Equal(t, byte(0b01101010), mirrorLookup[0b01010110])
	assert.Equal(t, byte(0b01010110), mirrorLookup[0b01101010])

	assert.Equal(t, byte(0b10010101), mirrorLookup[0b10101001])
	assert.Equal(t, byte(0b10101001), mirrorLookup[0b10010101])
}
