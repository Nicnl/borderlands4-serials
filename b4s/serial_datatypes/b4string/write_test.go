package b4string

import (
	"borderlands_4_serials/lib/bit"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteString(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			expectedBits := tt.bin
			expectedBits = strings.ReplaceAll(expectedBits, " ", "")
			expectedBits = strings.ReplaceAll(expectedBits, ":", "")
			expectedBits = strings.ReplaceAll(expectedBits, "[", "")
			expectedBits = strings.ReplaceAll(expectedBits, "]", "")
			expectedBits = strings.ReplaceAll(expectedBits, "-", "")

			bw := bit.NewWriter()
			Write(bw, tt.text)

			assert.Equal(t, expectedBits, bw.String())
		})
	}
}
