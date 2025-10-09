package serial_tokenizer

import (
	"borderlands_4_serials/lib/bit_reader"
	"fmt"
)

type Token byte

const (
	TOK_SEP1    Token = iota // "01" soft separator
	TOK_SEP2                 // "00" hard separator
	TOK_VARINT               // "100" ... nibble varint
	TOK_VARBIT               // "110" ... varbit
	TOK_KV_MODE              // "101" ... enter KV section
	TOK_KV_ADD               // "111" ... new KV entry inside section
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
	splitPositions := make([]int, 0)

	output := ""
	defer func() {
		fmt.Println("Data:", output)

		// split done positions, from right to left or we'll mess up the indexes
		doneString := t.bs.StringBefore()
		for i := len(splitPositions) - 1; i >= 0; i-- {
			pos := splitPositions[i]
			doneString = doneString[:pos] + "  " + doneString[pos:]
		}
		fmt.Println("Done:", doneString)
		fmt.Println("Fail:", t.bs.StringAfter())
		fmt.Println()
	}()
	for {
		splitPositions = append(splitPositions, t.bs.Pos())
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
		case TOK_KV_MODE:
			output += " {"
		case TOK_KV_ADD:
			output += " +"
		default:
			return fmt.Errorf("unknown token %d", token)
		}
	}

	return nil
}
