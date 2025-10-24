package b4string

import (
	"borderlands_4_serials/b4s/serial_datatypes/varint"
	"borderlands_4_serials/lib/bit"
	"borderlands_4_serials/lib/byte_mirror"
	"fmt"
)

func Read(br *bit.Reader) (string, error) {
	length, err := varint.Read(br)
	if err != nil {
		return "", fmt.Errorf("failed to read b4string length as varint: %w", err)
	}

	str := make([]byte, length, length)
	for i := uint32(0); i < length; i++ {
		rawCharBits, ok := br.ReadN(7)
		if !ok {
			return "", fmt.Errorf("unexpected end of data while reading b4string character")
		}
		str[i] = byte_mirror.Uint7Mirror[byte(rawCharBits)]
	}

	return string(str), nil
}
