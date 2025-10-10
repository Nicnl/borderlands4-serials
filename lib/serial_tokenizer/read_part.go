package serial_tokenizer

import (
	"borderlands_4_serials/lib/byte_mirror"
	"fmt"
)

func (t *Tokenizer) readPart() (uint32, error) {
	var output uint32

	// First block of 4 bits
	{
		block, ok := t.bs.ReadN(4)
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading part length")
		}
		block = uint32(byte_mirror.Uint4Mirror[byte(block)])

		output |= block
	}

	// Figure out the length of the next block
	remainingLength := 3
	{
		continuationBit, ok := t.bs.Read()
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading part continuation bit")
		}
		if continuationBit == 1 {
			remainingLength = 7
		}
	}

	// Read the 2nd block
	{
		block, ok := t.bs.ReadN(remainingLength)
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading part value")
		}

		if remainingLength == 3 {
			block = uint32(byte_mirror.Uint3Mirror[byte(block)])
			output |= block << 4
			return output, nil
		} else {
			block = uint32(byte_mirror.Uint7Mirror[byte(block)])
			output |= block << 4
		}
	}

	continuationBit, ok := t.bs.Read()
	if !ok {
		return 0, fmt.Errorf("unexpected end of data while reading part continuation bit")
	}
	if continuationBit == 0 {
		return output, nil
	}

	// Next block is 11
	{
		block, ok := t.bs.ReadN(11)
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading part value")
		}
		block = byte_mirror.Uint11Mirror[block]
		output |= block << 11
	}

	return output, nil
}
