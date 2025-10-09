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

func TestSerialTokenize(t *testing.T) {
	var tests = []struct {
		serial   string
		expected string
	}{
		{
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			"1",
		},
		{
			"@Ugr$WBm/!Fz!X=5&qXxA;nj3OOD#<4R",
			"1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.serial, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			tok := NewTokenizer(data)
			err = tok.Parse()
			assert.NoError(t, err)
		})
	}
}
