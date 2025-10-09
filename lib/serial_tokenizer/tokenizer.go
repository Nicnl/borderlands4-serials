package serial_tokenizer

import (
	"borderlands_4_serials/lib/bit_reader"
	"fmt"
)

type Tokenizer struct {
	bs *bit_reader.BitReader
}

func NewTokenizer(data []byte) *Tokenizer {
	return &Tokenizer{bs: bit_reader.NewBitReader(data)}
}

func (t *Tokenizer) expect(msg string, bits ...byte) error {
	for _, bit := range bits {
		b, ok := t.bs.Read()
		if !ok {
			return fmt.Errorf("unexpected end of data")
		}
		if b != bit {
			return fmt.Errorf(msg+" => expected bit %d, got %d", bit, b)
		}
	}
	return nil
}

func (t *Tokenizer) Parse() error {
	t.expect("magic", 0, 0, 1)

	first101 := false
	output := ""
	defer func() {
		fmt.Println("Data:", output)
		fmt.Println("Done:", t.bs.StringBefore())
		fmt.Println("Fail:", t.bs.StringAfter())
		fmt.Println()
	}()
	for {

		token, err := t.nextToken()
		if err != nil {
			return err
		}

		switch token {
		case TOK_SEP1:
			output += "|"
		case TOK_SEP2:
			output += ","
		case TOK_VARINT:
			v, err := t.readVarInt()
			if err != nil {
				return err
			}
			output += fmt.Sprintf(" %d", v)
		case TOK_VARBIT:
			v, err := t.readVarBit()
			if err != nil {
				return err
			}
			output += fmt.Sprintf(" %d", v)
		case TOK_UNK1:
			if !first101 {
				first101 = true
				output = ""
			}

			output += " <"
		case TOK_UNK2:
			output += " D"
		default:
			return fmt.Errorf("unknown token %d", token)
		}
	}

	return nil
}
