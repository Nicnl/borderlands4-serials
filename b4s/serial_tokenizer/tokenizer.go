package serial_tokenizer

import (
	"borderlands_4_serials/lib/bit"
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
	br             *bit.Reader
	splitPositions []int
}

func NewTokenizer(data []byte) *Tokenizer {
	return &Tokenizer{
		br:             bit.NewReader(data),
		splitPositions: make([]int, 0),
	}
}

func (t *Tokenizer) DoneString() string {
	splitted := t.br.StringBefore()
	for i := len(t.splitPositions) - 1; i >= 0; i-- {
		pos := t.splitPositions[i]
		splitted = splitted[:pos] + "  " + splitted[pos:]
	}
	return splitted
}

func (t *Tokenizer) BitReader() *bit.Reader {
	return t.br
}
