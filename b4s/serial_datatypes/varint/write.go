package varint

import "borderlands_4_serials/lib/bit"

func Write(bw *bit.Writer, value uint32) {
	// Figure out how many bits we need to represent 'value'
	nBits := 0
	{
		v2 := value
		for v2 > 0 {
			nBits++
			v2 >>= 1
		}

		// Special case: zero needs at least one block
		if nBits == 0 {
			nBits = 1
		}

		// If too many bits, the rest is discarded silently
		// TODO: handle this as an error maybe? Optional
		if nBits > VARINT_MAX_USABLE_BITS {
			nBits = VARINT_MAX_USABLE_BITS
		}
	}

	// Write complete blocks
	for nBits > VARINT_BITS_PER_BLOCK {
		// Data bits
		for range VARINT_BITS_PER_BLOCK {
			bw.WriteBit(byte(value) & 0b1)
			value >>= 1
			nBits--
		}

		// Continuation bit
		bw.WriteBit(1)
	}

	// Write partial last block
	if nBits > 0 {
		for i := 0; i < VARINT_BITS_PER_BLOCK; i++ {
			if nBits > 0 {
				// Data bits
				bw.WriteBit(byte(value) & 0b1)
				value >>= 1
				nBits--
			} else {
				// Padding bit
				bw.WriteBit(0)
			}
		}

		// No continuation (last block)
		bw.WriteBit(0)
	}
}
