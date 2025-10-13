package varint

import (
	"borderlands_4_serials/lib/bit"
	"borderlands_4_serials/lib/byte_mirror"
	"fmt"
)

func Read(br *bit.Reader) (uint32, error) {
	var (
		dataRead = 0
		output   uint32
	)
	for range 4 {
		// Read standard block
		{
			block32, ok := br.ReadN(4)
			if !ok {
				return 0, fmt.Errorf("unexpected end of data while reading varint")
			}
			output |= uint32(byte_mirror.Uint4Mirror[byte(block32)]) << dataRead
			dataRead += 4
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
