package serial_tokenizer

import (
	"fmt"
	"io"
)

func (t *Tokenizer) NextToken() (Token, error) {
	// Record the split position for a better debug  output
	t.splitPositions = append(t.splitPositions, t.br.Pos())

	// Read the first two bits, as zero-tokens are two bits long
	b1, b2, ok := t.br.Read2()
	if !ok {
		return 0, io.EOF
	}

	tok := b1<<1 | b2
	switch tok {
	case 0b00:
		return TOK_SEP1, nil
	case 0b01:
		return TOK_SEP2, nil
	}

	// If we are here, the first bit was a 1 => one-tokens are three-bit long
	// => Read one mor bit, and combine it to form the full token
	{
		b3, ok := t.br.Read()
		if !ok {
			return 0, io.EOF
		}
		tok = tok<<1 | b3
	}

	switch tok {
	case 0b100:
		return TOK_VARINT, nil
	case 0b110:
		return TOK_VARBIT, nil
	case 0b101:
		return TOK_PART, nil
	case 0b111:
		return TOK_STRING, nil
	}

	// If we are here, it means our rules didn't match any valid token
	// => Rewind so the bitwriter, as if we never read those bits + error out
	t.br.Rewind(3)
	return 0, fmt.Errorf("invalid token, got %03b at position %d", tok, t.br.Pos())
}
