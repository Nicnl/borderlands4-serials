package serial

import (
	"borderlands_4_serials/b4s/serial_datatypes/b4string"
	"borderlands_4_serials/b4s/serial_datatypes/part"
	"borderlands_4_serials/b4s/serial_datatypes/varbit"
	"borderlands_4_serials/b4s/serial_datatypes/varint"
	"borderlands_4_serials/b4s/serial_tokenizer"
	"fmt"
	"io"
)

func Deserialize(data []byte) (Serial, string, error) {
	t := serial_tokenizer.NewTokenizer(data)

	// Expect the magic header as the first bits
	if err := t.Expect("magic header", 0, 0, 1, 0, 0, 0, 0); err != nil {
		return nil, t.DoneString(), err
	}

	var (
		br     = t.BitReader()
		blocks = make([]Block, 0, 50) // Preallocate some space for performance

		// Keep track of the trailing terminators for sanitization later
		trailingTerminators = 0
	)

	for {
		token, err := t.NextToken()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, t.DoneString(), err
		}

		block := Block{
			Token: token,
		}

		// Count the trailing terminators for sanitization later
		if token == serial_tokenizer.TOK_SEP1 {
			trailingTerminators++
		} else {
			trailingTerminators = 0
		}

		switch token {
		case serial_tokenizer.TOK_SEP1:
			// Nothing to do

		case serial_tokenizer.TOK_SEP2:
			// Nothing to do

		case serial_tokenizer.TOK_VARINT:
			v, err := varint.Read(br)
			if err != nil {
				return nil, t.DoneString(), err
			}
			block.Value = v

		case serial_tokenizer.TOK_VARBIT:
			v, err := varbit.Read(br)
			if err != nil {
				return nil, t.DoneString(), err
			}
			block.Value = v

		case serial_tokenizer.TOK_PART:
			p, err := part.Read(t)
			if err != nil {
				return nil, t.DoneString(), err
			}

			block.Part = p

		case serial_tokenizer.TOK_STRING:
			str, err := b4string.Read(br)
			if err != nil {
				return nil, t.DoneString(), err
			}

			block.ValueStr = str
		default:
			return nil, t.DoneString(), fmt.Errorf("unknown token %d", token)
		}

		blocks = append(blocks, block)
	}

	// Sanitization: we probably read the zero-padding as terminators.
	// Only one terminator is needed, remove the extra ones
	if trailingTerminators > 1 {
		blocks = blocks[:len(blocks)-(trailingTerminators-1)]
	}

	return blocks, t.DoneString(), nil
}
