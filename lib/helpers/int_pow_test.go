package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntPow(t *testing.T) {
	assert.Equal(t, 1, IntPow(2, 0))
	assert.Equal(t, 2, IntPow(2, 1))
	assert.Equal(t, 4, IntPow(2, 2))
	assert.Equal(t, 8, IntPow(2, 3))
	assert.Equal(t, 16, IntPow(2, 4))
	assert.Equal(t, 32, IntPow(2, 5))
	assert.Equal(t, 243, IntPow(3, 5))
	assert.Equal(t, 1, IntPow(5, 0))
	assert.Equal(t, 5, IntPow(5, 1))
	assert.Equal(t, 25, IntPow(5, 2))
	assert.Equal(t, 125, IntPow(5, 3))
	assert.Equal(t, 625, IntPow(5, 4))
	assert.Equal(t, 3125, IntPow(5, 5))
}
