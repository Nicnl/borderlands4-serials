package serial_tokenizer

import (
	"fmt"
)

var BITS = 13

func (t *Tokenizer) readPart() (uint32, error) {
	// Read fixed 13bits
	block, ok := t.bs.ReadN(BITS)
	if !ok {
		return 0, fmt.Errorf("unexpected end of data while reading part")
	}

	return block, nil
}
