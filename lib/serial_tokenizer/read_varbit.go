package serial_tokenizer

import (
	"borderlands_4_serials/lib/byte_mirror"
	"fmt"
)

func (t *Tokenizer) readVarBit() (uint32, error) {
	length, ok := t.bs.ReadN(5)
	if !ok {
		return 0, fmt.Errorf("unexpected end of data while reading varbit length")
	}
	length = uint32(byte_mirror.Uint5Mirror[byte(length)])

	if length == 0 {
		_, ok := t.bs.Read()
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading varbit length")
		}
	}

	var v uint32
	for i := uint32(0); i < length; i++ {
		bit, ok := t.bs.Read()
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading varbit value")
		}

		v |= uint32(bit) << i
	}

	return v, nil
}
