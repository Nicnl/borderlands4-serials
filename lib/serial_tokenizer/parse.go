package serial_tokenizer

import (
	"borderlands_4_serials/lib/bit_reader"
	"fmt"
	"io"
)

type Token byte

const (
	TOK_SEP1     Token = iota // "00" hard separator (terminator?)
	TOK_SEP2                  // "01" soft separator
	TOK_VARINT                // "100" ... nibble varint
	TOK_VARBIT                // "110" ... varbit
	TOK_PART                  // "101" ... enter KV section
	TOK_PART_111              // "111" weird, seems linked to DLC items, we DO NOT touch that
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

func (t *Tokenizer) Parse() (int, string, error) {
	if err := t.expect("magic header", 0, 0, 1, 0, 0, 0, 0); err != nil {
		return -1, "", err
	}

	debugOutput := ""
	defer func() {
		if strAfter := t.bs.StringAfter(); strAfter != "" || true {
			//fmt.Println("Debug", debugOutput)
			//fmt.Println("Data remaining", strAfter)
		}
	}()

	var (
		foundLevel       = -1
		currentParsedInt = 0
	)

OUTER:
	for {
		token, err := t.nextToken()
		if err == io.EOF {
			break
		} else if err != nil {
			return foundLevel, debugOutput, err
		}

		switch token {
		case TOK_SEP1:
			debugOutput += "|"
		case TOK_SEP2:
			debugOutput += ","
		case TOK_VARINT:
			v, err := t.readVarint()
			if err != nil {
				return foundLevel, debugOutput, err
			}
			debugOutput += fmt.Sprintf(" %d", v)
			currentParsedInt++
			if currentParsedInt == 4 {
				foundLevel = int(v)
			}
		case TOK_VARBIT:
			v, err := t.readVarBit()
			if err != nil {
				return foundLevel, debugOutput, err
			}
			debugOutput += fmt.Sprintf(" %d", v)
			currentParsedInt++
			if currentParsedInt == 4 {
				foundLevel = int(v)
			}
		case TOK_PART_111:
			// UNSUPPORTED, unknown
			// Seems linked to DLC weapons
			// BUT it sometimes appears on items bought from the legendary vending machine
			// Luckily, it's always at the end of the data
			// => discard and forget about it
			break OUTER

		case TOK_PART:
			part, err := t.readPartV2()
			if err != nil {
				return foundLevel, debugOutput, err
			}

			switch part.SubType {
			case PART_SUBTYPE_NONE:
				debugOutput += fmt.Sprintf(" {%d}", part.Index)
			case PART_SUBTYPE_INT:
				debugOutput += fmt.Sprintf(" {%d:%d}", part.Index, part.Value)
			case PART_SUBTYPE_LIST:
				debugOutput += fmt.Sprintf(" {%d:[", part.Index)
				for i, v := range part.Values {
					if i != 0 {
						debugOutput += " "
					}
					debugOutput += fmt.Sprintf("%d", v)
				}
				debugOutput += "]}"
			default:
				return foundLevel, debugOutput, fmt.Errorf("unknown part subtype %d", part.SubType)
			}
		//case TOK_111:
		//	debugOutput += " <111>"
		default:
			return foundLevel, debugOutput, fmt.Errorf("unknown token %d", token)
		}
	}

	return foundLevel, debugOutput, nil
}
