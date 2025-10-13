package serial_parser

import (
	"borderlands_4_serials/b4s/serial_datatypes/part"
	"borderlands_4_serials/b4s/serial_datatypes/varbit"
	"borderlands_4_serials/b4s/serial_datatypes/varint"
	"borderlands_4_serials/b4s/serial_tokenizer"
	"fmt"
	"io"
)

type Parsed struct {
	Debug string
	Bits  string
}

func Parse(data []byte) (Parsed, error) {
	t := serial_tokenizer.NewTokenizer(data)
	br := t.BitReader()

	if err := t.Expect("magic header", 0, 0, 1, 0, 0, 0, 0); err != nil {
		return Parsed{}, err
	}

	output := Parsed{}

OUTER:
	for {
		token, err := t.NextToken()
		if err == io.EOF {
			break
		} else if err != nil {
			return Parsed{}, err
		}

		switch token {
		case serial_tokenizer.TOK_SEP1:
			output.Debug += "|"
		case serial_tokenizer.TOK_SEP2:
			output.Debug += ","
		case serial_tokenizer.TOK_VARINT:
			v, err := varint.Read(br)
			if err != nil {
				return Parsed{}, err
			}
			output.Debug += fmt.Sprintf(" %d", v)
		case serial_tokenizer.TOK_VARBIT:
			v, err := varbit.Read(br)
			if err != nil {
				return Parsed{}, err
			}
			output.Debug += fmt.Sprintf(" %d", v)
		case serial_tokenizer.TOK_PART_111:
			// UNSUPPORTED, unknown
			// Seems linked to DLC weapons
			// BUT it sometimes appears on items bought from the legendary vending machine
			// Luckily, it's always at the end of the data
			// => discard and forget about it
			//continue
			break OUTER

		case serial_tokenizer.TOK_PART:
			p, err := part.Read(t)
			if err != nil {
				return Parsed{}, err
			}

			output.Debug += p.String()
		default:
			return Parsed{}, fmt.Errorf("unknown token %d", token)
		}
	}

	return output, nil
}
