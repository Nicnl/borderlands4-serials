package serial_tokenizer

import (
	"borderlands_4_serials/lib/byte_mirror"
	"fmt"
)

func (t *Tokenizer) readVarInt() (uint16, error) {
	var (
		dataLeft = 4
		value    uint32
		block32  uint32
		ok       bool
	)
	for dataLeft > 0 {
		block32, ok = t.bs.ReadN(4)
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading varint")
		}
		dataLeft--

		block := uint32(byte_mirror.Uint4Mirror[byte(block32)])
		value |= block << (12 - 4*dataLeft)

		cont, ok := t.bs.Read()
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading varint")
		}
		if cont == 0 {
			break
		}
	}

	return uint16(value), nil
}
