package serial_tokenizer

import "fmt"

const (
	TOK_SEP1   = iota // 00
	TOK_SEP2          // 01
	TOK_VARINT        // 100
	TOK_VARBIT        // 110
	TOK_UNK1
	TOK_UNK2

	// TODO: 101 and 111 ?
)

func (t *Tokenizer) nextToken() (byte, error) {
	bit1, ok := t.bs.Read()
	if !ok {
		return 0, fmt.Errorf("unexpected end of data while reading token")
	}

	bit2, ok := t.bs.Read()
	if !ok {
		return 0, fmt.Errorf("unexpected end of data while reading token")
	}

	if bit1 == 0 && bit2 == 0 {
		return TOK_SEP1, nil
	}

	if bit1 == 0 && bit2 == 1 {
		return TOK_SEP2, nil
	}

	bit3, ok := t.bs.Read()
	if !ok {
		return 0, fmt.Errorf("unexpected end of data while reading token")
	}

	if bit1 == 1 && bit2 == 0 && bit3 == 0 {
		return TOK_VARINT, nil
	}

	if bit1 == 1 && bit2 == 0 && bit3 == 1 {
		//return TOK_UNK1, nil
	}

	if bit1 == 1 && bit2 == 1 && bit3 == 0 {
		return TOK_VARBIT, nil
	}

	if bit1 == 1 && bit2 == 1 && bit3 == 1 {
		//return TOK_UNK2, nil
	}

	t.bs.Rewind(3)
	return 0, fmt.Errorf("invalid token, got %d%d%d at position %d", bit1, bit2, bit3, t.bs.Pos())
}
