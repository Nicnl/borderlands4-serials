package varbit

import (
	"borderlands_4_serials/lib/bit"
	"borderlands_4_serials/lib/helpers"
)

func Write(bw *bit.Writer, value uint32) {
	// Figure out how many bits we need to represent 'value'
	nBits := helpers.IntBitsSize(value, 0, helpers.IntPow(2, VARBIT_LENGTH_BLOCK_SIZE)-1)
	//fmt.Println("nBits =", nBits, "for value =", value)

	// Write length
	lengthBits := nBits
	for range VARBIT_LENGTH_BLOCK_SIZE {
		bw.WriteBit(byte(lengthBits) & 0b1)
		lengthBits >>= 1
	}

	// Write value bits
	for i := 0; i < nBits; i++ {
		bw.WriteBit(byte(value) & 0b1)
		value >>= 1
	}
}
