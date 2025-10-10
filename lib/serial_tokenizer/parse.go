package serial_tokenizer

import (
	"borderlands_4_serials/lib/bit_reader"
	"fmt"
	"io"
)

type Token byte

const (
	TOK_SEP1   Token = iota // "01" soft separator
	TOK_SEP2                // "00" hard separator
	TOK_VARINT              // "100" ... nibble varint
	TOK_VARBIT              // "110" ... varbit
	TOK_PART                // "101" ... enter KV section
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

func (t *Tokenizer) Parse() (string, error) {
	splitPositions := make([]int, 0)

	if err := t.expect("magic header", 0, 0, 1, 0, 0, 0, 0); err != nil {
		return "", err
	}

	defer func() {
		fmt.Println("AFTER", t.bs.StringAfter())
	}()

	debugOutput := ""

	for {
		splitPositions = append(splitPositions, t.bs.Pos())
		token, err := t.nextToken()
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}

		switch token {
		case TOK_SEP1:
			debugOutput += "|"
		case TOK_SEP2:
			debugOutput += ","
		case TOK_VARINT:
			v, err := t.readVarInt()
			if err != nil {
				return "", err
			}
			debugOutput += fmt.Sprintf(" %d", v)
		case TOK_VARBIT:
			v, err := t.readVarBit()
			if err != nil {
				return "", err
			}
			debugOutput += fmt.Sprintf(" %d", v)
		case TOK_PART:
			v, err := t.readPart()
			if err != nil {
				return "", err
			}
			debugOutput += fmt.Sprintf(" {%d}", v)
		default:
			return "", fmt.Errorf("unknown token %d", token)
		}
	}

	return debugOutput, nil
}
