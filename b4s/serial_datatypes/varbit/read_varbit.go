package varbit

import (
	"borderlands_4_serials/lib/bit"
	"borderlands_4_serials/lib/byte_mirror"
	"fmt"
)

func Read(br *bit.Reader) (uint32, error) {
	length, ok := br.ReadN(5)
	if !ok {
		return 0, fmt.Errorf("unexpected end of data while reading varbit length")
	}
	length = uint32(byte_mirror.Uint5Mirror[byte(length)])

	if length == 0 {
		// TODO: check if really 0, or a special case meaning 32
		return 0, nil
		_, ok := br.Read()
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading varbit length")
		}
	}

	var v uint32
	for i := uint32(0); i < length; i++ {
		bit, ok := br.Read()
		if !ok {
			return 0, fmt.Errorf("unexpected end of data while reading varbit value")
		}

		v |= uint32(bit) << i
	}

	return v, nil
}
