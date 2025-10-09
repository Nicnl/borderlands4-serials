package serial_tokenizer

import (
	"fmt"
	"testing"
)

func TestByWalkingAnalysis(t *testing.T) {
	serials := []string{
		"0000101101011101001010100011001011000001010101011011101001010100000000000",
		"0000101101011011001010100011001011000001010101011011101001010100000000000",
		"0000101100011001001010100011001011000001010101011011101001010100000000000",
		"0000101111011001001010100011001011000001010101011011101001010100000000000",
		"0000101101111001001010100011001011000001010101011011101001010100000000000",
		"0000101101010001001010100011001011000001010101011011101001010100000000000",
		"0000101001011001001010100011001011000001010101011011101001010100000000000",
		"0000101101011000001010100011001011000001010101011011101001010100000000000",
	}

	serialInts := make(map[string][]int64)

OUTER:
	for _, s := range serials {
		data := binToBytes(s)
		tok := NewTokenizer(data)

		for i := 0; i < len(s); i++ {
			tok.bs.SetPos(i)

			token, err := tok.nextToken()
			if err != nil {
				break OUTER
			}

			if token == TOK_VARINT {
				val, err := tok.readVarInt()
				if err != nil {
					t.Errorf("Error reading varint at pos %d in serial %s: %v", i, s, err)
					continue OUTER
				}
				serialInts[s] = append(serialInts[s], int64(val))
			} else if token == TOK_VARBIT {
				val, err := tok.readVarBit()
				if err != nil {
					t.Errorf("Error reading varbit at pos %d in serial %s: %v", i, s, err)
					continue OUTER
				}
				serialInts[s] = append(serialInts[s], int64(val))
			}
		}
		fmt.Println("Serial:", s)
		fmt.Println("Ints:", serialInts[s])
		fmt.Println()
	}

}
