package serial_tokenizer

import (
	"borderlands_4_serials/lib/byte_mirror"
	"fmt"
)

func (t *Tokenizer) readVarintExtended() (uint32, error) {
	var (
		dataRead = 0
		output   uint32
	)

	// Read standard block
	block32, ok := t.bs.ReadN(4)
	if !ok {
		return 0, fmt.Errorf("unexpected end of data while reading varint")
	}
	block32 = uint32(byte_mirror.Uint4Mirror[byte(block32)])
	output |= block32 << dataRead
	dataRead += 4

	// Obtain continuation bit in the middle of the first block
	cont, ok := t.bs.Read()
	if !ok {
		return 0, fmt.Errorf("unexpected end of data while reading varint")
	}

	// Read the 3 extended bits
	extendedBits, ok := t.bs.ReadN(3)
	if !ok {
		return 0, fmt.Errorf("unexpected end of data while reading varint")
	}
	extendedBits = uint32(byte_mirror.Uint3Mirror[byte(extendedBits)])
	output |= extendedBits << dataRead
	dataRead += 3

	// If continuation bit is not set, we are done
	if cont == 0 {
		return output, nil
	}

	for range 3 {
		// First bit
		{
			cont, ok = t.bs.Read()
			if !ok {
				return 0, fmt.Errorf("unexpected end of data while reading varint")
			}
			output |= uint32(cont) << dataRead
			dataRead += 1
		}

		// Continuation bit in position 2
		cont, ok = t.bs.Read()
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading varint")
		}

		// Third remaining bits
		{
			block32, ok = t.bs.ReadN(3)
			if !ok {
				return 0, fmt.Errorf("unexpected end of data while reading varint")
			}
			output |= uint32(byte_mirror.Uint3Mirror[byte(block32)]) << dataRead
			dataRead += 3
		}

		if cont == 0 {
			break
		}
	}

	return output, nil
}
