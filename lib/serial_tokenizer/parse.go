package serial_tokenizer

import (
	"borderlands_4_serials/lib/bit_reader"
	"fmt"
	"io"
)

type Token byte

const (
	TOK_SEP1            Token = iota // "01" soft separator
	TOK_SEP2                         // "00" hard separator
	TOK_VARINT                       // "100" ... nibble varint
	TOK_VARBIT                       // "110" ... varbit
	TOK_VARINT_EXTENDED              // "101" ... enter KV section
	TOK_111
)

type Tokenizer struct {
	bs             *bit_reader.BitReader
	splitPositions []int
}

func NewTokenizer(data []byte) *Tokenizer {
	return &Tokenizer{
		bs:             bit_reader.NewBitReader(data),
		splitPositions: make([]int, 0),
	}
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

func (t *Tokenizer) DoneString() string {
	splitted := t.bs.StringBefore()
	for i := len(t.splitPositions) - 1; i >= 0; i-- {
		pos := t.splitPositions[i]
		splitted = splitted[:pos] + "  " + splitted[pos:]
	}
	return splitted
}

func (t *Tokenizer) Parse() (string, error) {
	if err := t.expect("magic header", 0, 0, 1, 0, 0, 0, 0); err != nil {
		return "", err
	}

	debugOutput := ""
	defer func() {
		if strAfter := t.bs.StringAfter(); strAfter != "" || true {
			fmt.Println("Debug", debugOutput)
			fmt.Println("Data remaining", strAfter)
		}
	}()

	for {
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
			v, err := t.readVarint()
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
		case TOK_VARINT_EXTENDED:
			v, err := t.readVarintExtended()
			if err != nil {
				return "", err
			}
			debugOutput += fmt.Sprintf(" {%d}", v)
		//case TOK_111:
		//	debugOutput += " <111>"
		default:
			return "", fmt.Errorf("unknown token %d", token)
		}
	}

	return debugOutput, nil
}
