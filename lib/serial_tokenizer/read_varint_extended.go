package serial_tokenizer

import (
	"fmt"
)

func (t *Tokenizer) readPart() (uint32, byte, error) {
	index, err := t.readVarint()
	if err != nil {
		return 0, 0, err
	}

	flag, ok := t.bs.ReadN(3)
	if !ok {
		return 0, 0, fmt.Errorf("unexpected end of data while reading part 101 flag")
	}

	switch flag {
	case 0b000, 0b010, 0b001:
		return index, byte(flag), nil
	case 0b101, 0b110, 0b111:
		return 0, 0, fmt.Errorf("unsupported part 101 flag <:%03b> at position %d", flag, t.bs.Pos()-3)
	default:
		return 0, 0, fmt.Errorf("unknown part 101 flag <:%03b> at position %d", flag, t.bs.Pos()-3)
	}
}
