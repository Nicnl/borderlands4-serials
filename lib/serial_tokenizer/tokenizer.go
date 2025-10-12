package serial_tokenizer

import (
	"fmt"
	"io"
)

func (t *Tokenizer) nextToken() (Token, error) {
	// Record the split position for a better debug  output
	t.splitPositions = append(t.splitPositions, t.bs.Pos())

	// Read the first two bits, as zero-tokens are two bits long
	bit1, ok := t.bs.Read()
	if !ok {
		return 0, io.EOF
	}

	bit2, ok := t.bs.Read()
	if !ok {
		return 0, io.EOF
	}

	// Token 00 (terminator) + 01 (separator)
	{
		if bit1 == 0 && bit2 == 0 {
			return TOK_SEP1, nil
		}

		if bit1 == 0 && bit2 == 1 {
			return TOK_SEP2, nil
		}
	}

	// If we are here, the first bit was a 1 => tokens are now three-bits long
	bit3, ok := t.bs.Read()
	if !ok {
		return 0, io.EOF
	}

	// 100 => Varint
	if bit1 == 1 && bit2 == 0 && bit3 == 0 {
		return TOK_VARINT, nil
	}

	// 110 => Varbit
	if bit1 == 1 && bit2 == 1 && bit3 == 0 {
		return TOK_VARBIT, nil
	}

	// 101 => complex weapon part
	if bit1 == 1 && bit2 == 0 && bit3 == 1 {
		return TOK_PART, nil
	}

	// 111 => ??? looks like part, not sure
	if bit1 == 1 && bit2 == 1 && bit3 == 1 {
		return TOK_PART_111, nil
	}

	// Rewind so the caller can read the invalid token if needed
	t.bs.Rewind(3)
	return 0, fmt.Errorf("invalid token, got %d%d%d at position %d", bit1, bit2, bit3, t.bs.Pos())
}
