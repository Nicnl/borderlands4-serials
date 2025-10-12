package serial_tokenizer

import (
	"borderlands_4_serials/lib/byte_mirror"
	"fmt"
)

func (t *Tokenizer) readPart() (uint32, byte, uint32, error) {
	index, err := t.readVarint()
	if err != nil {
		return 0, 0, 0, err
	}

	flag, ok := t.bs.ReadN(3)
	if !ok {
		return 0, 0, 0, fmt.Errorf("unexpected end of data while reading part 101 flag")
	}
	//fmt.Printf("flag = %03b", flag)

	switch flag {
	//case 0b010, 0b001, 0b100, 0b000, 0b110, 0b101:
	case 0b010, 0b001:
		return index, byte(flag), 0, nil
	case 0b100:
		param, ok := t.bs.ReadN(6)
		if !ok {
			return 0, 0, 0, fmt.Errorf("unexpected end of data while reading part 101:100 extra 6 bits")
		}
		param = byte_mirror.GenericMirror(param, 6)
		return index, byte(flag), uint32(param), nil
	case 0b111:
		param, ok := t.bs.ReadN(8)
		if !ok {
			return 0, 0, 0, fmt.Errorf("unexpected end of data while reading part 101:111 extra 6 bits")
		}
		param = byte_mirror.GenericMirror(param, 8)
		return index, byte(flag), uint32(param), nil
	default:
		return 0, 0, 0, fmt.Errorf("unknown part 101 flag <:%03b> at position %d", flag, t.bs.Pos()-3)
	}
}
