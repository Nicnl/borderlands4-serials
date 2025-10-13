package helpers

// IntBitsSize returns the minimum number of bits needed to represent the given unsigned integer value,
// Special case: if value is 0, it returns 1, because at least one bit is needed to represent zero.
// Additional bits are silently discarded if exceeding maxSize.
func IntBitsSize(v uint32, minSize int, maxSize int) int {
	nBits := 0

	// Count each bit until it's zero
	for v > 0 {
		nBits++
		v >>= 1
	}

	if nBits < minSize {
		nBits = minSize
	}

	if nBits > maxSize {
		nBits = maxSize
	}

	return nBits
}
