package varint

import (
	"borderlands_4_serials/lib/bit"
	"borderlands_4_serials/lib/byte_mirror"
	"fmt"
)

const (
	VARINT_NB_BLOCKS       = 4
	VARINT_BITS_PER_BLOCK  = 4
	VARINT_MAX_USABLE_BITS = VARINT_NB_BLOCKS * VARINT_BITS_PER_BLOCK
)

func Read(br *bit.Reader) (uint32, error) {
	var (
		dataRead = 0
		output   uint32
	)
	for range VARINT_NB_BLOCKS {
		// Read standard block
		{
			block32, ok := br.ReadN(VARINT_BITS_PER_BLOCK)
			if !ok {
				return 0, fmt.Errorf("unexpected end of data while reading varint")
			}
			output |= uint32(byte_mirror.Uint4Mirror[byte(block32)]) << dataRead
			dataRead += VARINT_BITS_PER_BLOCK
		}

		// Continuation bit
		{
			cont, ok := br.Read()
			if !ok {
				return 0, fmt.Errorf("unexpected end of data while reading varint")
			}

			if cont == 0 {
				break
			}
		}
	}

	return output, nil
}
