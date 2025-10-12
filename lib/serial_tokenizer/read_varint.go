package serial_tokenizer

import (
	"borderlands_4_serials/lib/byte_mirror"
	"fmt"
)

func (t *Tokenizer) readVarint() (uint32, error) {
	var (
		dataRead = 0
		output   uint32
	)
	for range 4 {
		// Read standard block
		{
			block32, ok := t.bs.ReadN(4)
			if !ok {
				return 0, fmt.Errorf("unexpected end of data while reading varint")
			}
			output |= uint32(byte_mirror.Uint4Mirror[byte(block32)]) << dataRead
			dataRead += 4
		}

		// Continuation bit
		{
			cont, ok := t.bs.Read()
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

func (t *Tokenizer) readVarintLong() (uint32, error) {
	var (
		dataRead = 0
		output   uint32
	)
	for range 4 {
		// Read standard block
		{
			block32, ok := t.bs.ReadN(4)
			if !ok {
				return 0, fmt.Errorf("unexpected end of data while reading varint")
			}
			output |= uint32(byte_mirror.Uint4Mirror[byte(block32)]) << dataRead
			dataRead += 4
		}

		// Continuation bit
		{
			cont, ok := t.bs.Read()
			if !ok {
				return 0, fmt.Errorf("unexpected end of data while reading varint")
			}

			if cont == 0 {
				break
			}
		}
	}

	// read 3 bit (long form)
	{
		block3, ok := t.bs.ReadN(3)
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading varint long form")
		}
		output |= uint32(byte_mirror.Uint3Mirror[byte(block3)]) << dataRead
		dataRead += 3
	}

	return output, nil
}
