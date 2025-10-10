package serial_tokenizer

import (
	"borderlands_4_serials/lib/byte_mirror"
	"fmt"
)

func (t *Tokenizer) readPart() (uint16, error) {
	var output uint16

	// First block of 4 bits
	{
		block, ok := t.bs.ReadN(4)
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading part length")
		}
		block = uint32(byte_mirror.Uint4Mirror[byte(block)])

		output |= uint16(block)
	}

	// Figure out the length of the next block
	remainingLength := 3
	{
		continuationBit, ok := t.bs.Read()
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading part continuation bit")
		}
		if continuationBit == 1 {
			remainingLength = 8
		}
	}

	// Read the remaining block
	{
		block, ok := t.bs.ReadN(remainingLength)
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading part value")
		}

		if remainingLength == 3 {
			block = uint32(byte_mirror.Uint3Mirror[byte(block)])
		} else {
			block = uint32(byte_mirror.Uint8Mirror[byte(block)])
		}

		output |= uint16(block) << 4
	}

	return output, nil
}
