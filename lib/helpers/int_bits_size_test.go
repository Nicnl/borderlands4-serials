package helpers

import (
	"fmt"
	"testing"
)

func TestIntsBitsSize(t *testing.T) {
	tests := []struct {
		value   uint32
		minSize int
		maxSize int
		nBits   int
	}{
		// Check zero case
		{0, 1, 6, 1},
		{0, 0, 6, 0},

		// Check normal cases under a byte
		{0b00000001, 1, 6, 1},
		{0b00000010, 1, 6, 2},
		{0b00000100, 1, 6, 3},
		{0b00001000, 1, 6, 4},
		{0b00010000, 1, 6, 5},
		{0b00100000, 1, 6, 6},
		{0b01000000, 1, 6, 6}, // discarded: > maxSize
		{0b10000000, 1, 6, 6}, // discarded: > maxSize

		// Check normal cases at a byte boundary
		{0b00000000000001, 1, 12, 1},
		{0b00000000000010, 1, 12, 2},
		{0b00000000000100, 1, 12, 3},
		{0b00000000001000, 1, 12, 4},
		{0b00000000010000, 1, 12, 5},
		{0b00000000100000, 1, 12, 6},
		{0b00000001000000, 1, 12, 7},
		{0b00000010000000, 1, 12, 8},
		{0b00000100000000, 1, 12, 9},
		{0b00001000000000, 1, 12, 10},
		{0b00010000000000, 1, 12, 11},
		{0b00100000000000, 1, 12, 12},
		{0b01000000000000, 1, 12, 12}, // discarded: > maxSize
		{0b10000000000000, 1, 12, 12}, // discarded: > maxSize

		// Check normal cases at a 32-bit boundary

	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("value=%d,maxSize=%d", test.value, test.maxSize), func(t *testing.T) {
			nBits := IntBitsSize(test.value, test.minSize, test.maxSize)
			if nBits != test.nBits {
				t.Errorf("IntBitsSize(%d, %d) = %d; want %d", test.value, test.maxSize, nBits, test.nBits)
			}
		})
	}
}
