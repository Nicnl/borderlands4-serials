package serial_tokenizer

import (
	"borderlands_4_serials/lib/b85"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func binToBytes(s string) []byte {
	s = strings.ReplaceAll(s, " ", "")

	n := (len(s) + 7) / 8
	data := make([]byte, n)

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			data[i/8] |= 1 << (7 - uint(i)%8)
		}
	}

	return data
}

func TestSerialDecode1(t *testing.T) {
	data, err := b85.Decode("@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00")
	assert.NoError(t, err)

	tok := NewTokenizer(data)
	assert.NoError(t, tok.Parse())
}

func TestSerialDecode2(t *testing.T) {
	data, err := b85.Decode("@Ugr$WBm/!Fz!X=5&qXxA;nj3OOD#<4R")
	assert.NoError(t, err)

	tok := NewTokenizer(data)
	assert.NoError(t, tok.Parse())
}
