package serial_tokenizer

import (
	"fmt"
	"io"
)

func (t *Tokenizer) nextToken() (Token, error) {
	bit1, ok := t.bs.Read()
	if !ok {
		return 0, io.EOF
	}

	bit2, ok := t.bs.Read()
	if !ok {
		return 0, io.EOF
	}

	if bit1 == 0 && bit2 == 0 {
		return TOK_SEP1, nil
	}

	if bit1 == 0 && bit2 == 1 {
		return TOK_SEP2, nil
	}

	bit3, ok := t.bs.Read()
	if !ok {
		return 0, io.EOF
	}

	if bit1 == 1 && bit2 == 0 && bit3 == 0 {
		return TOK_VARINT, nil
	}

	if bit1 == 1 && bit2 == 1 && bit3 == 0 {
		return TOK_VARBIT, nil
	}

	if bit1 == 1 && bit2 == 0 && bit3 == 1 {
		return TOK_PART, nil
	}

	if bit1 == 1 && bit2 == 1 && bit3 == 1 {
		//return TOK_111, nil
	}

	t.bs.Rewind(3)
	return 0, fmt.Errorf("invalid token, got %d%d%d at position %d", bit1, bit2, bit3, t.bs.Pos())
}
