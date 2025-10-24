package serial_tokenizer

import (
	"borderlands_4_serials/lib/bit"
	"strings"
)

type Token byte

const (
	TOK_SEP1            Token = iota // "00" hard separator (terminator?)
	TOK_SEP2                         // "01" soft separator
	TOK_VARINT                       // "100" ... nibble varint
	TOK_VARBIT                       // "110" ... varbit
	TOK_PART                         // "101" ... complex part block
	TOK_UNSUPPORTED_111              // "111" is linked to DLC items, we DO NOT touch that
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
	original := t.br.FullString()
	if len(t.splitPositions) == 0 {
		return original
	}

	var sb strings.Builder
	sb.Grow(len(original) + len(t.splitPositions)*2)

	lastPos := 0
	for _, pos := range t.splitPositions {
		sb.WriteString(original[lastPos:pos])
		sb.WriteString("  ")
		lastPos = pos
	}
	sb.WriteString(original[lastPos:])

	return sb.String()
}

func (t *Tokenizer) BitReader() *bit.Reader {
	return t.br
}
