package part

import (
	"borderlands_4_serials/b4s/serial_tokenizer"
	"borderlands_4_serials/lib/bit"
	"borderlands_4_serials/lib/helpers"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWritePart(t *testing.T) {
	for _, testGroup := range testsGroups {
		t.Run(testGroup.name, func(t *testing.T) {
			for _, test := range testGroup.tests {
				t.Run(test.bits, func(t *testing.T) {
					expectedBits := test.bits
					expectedBits = strings.ReplaceAll(expectedBits, " ", "")
					expectedBits = strings.ReplaceAll(expectedBits, ":", "")
					expectedBits = strings.ReplaceAll(expectedBits, "[", "")
					expectedBits = strings.ReplaceAll(expectedBits, "]", "")
					expectedBits = strings.ReplaceAll(expectedBits, "-", "")

					bw := bit.NewWriter()
					Write(bw, test.expected)
					writtenBits := bw.String()
					assert.Equal(t, expectedBits, writtenBits)
				})
			}
		})
	}
}

func TestWritePartRoundtripSubtypeNone(t *testing.T) {
	maxIndex := helpers.IntPow(2, 16) - 1 // All possible values for varint (4 blocks of 4 usable bits)

	for index := 0; index <= maxIndex; index++ {
		part := Part{
			SubType: SUBTYPE_NONE,
			Index:   uint32(index),
		}

		// Write
		//fmt.Println("Write:", part.String())
		bw := bit.NewWriter()
		Write(bw, part)

		// Read
		tok := serial_tokenizer.NewTokenizer(bw.Data())
		readPart, err := Read(tok)
		assert.NoError(t, err)
		//fmt.Println("Read:", readPart.String())

		// Compare
		assert.Equal(t, part, readPart)

		//fmt.Println()
	}
}

func TestWritePartRoundtripSubtypeInt(t *testing.T) {
	var (
		maxIndex = helpers.IntPow(2, 9) - 1
		maxValue = helpers.IntPow(2, 9) - 1
	)

	for index := 0; index <= maxIndex; index++ {
		for value := 0; value <= maxValue; value++ {
			part := Part{
				SubType: SUBTYPE_INT,
				Index:   uint32(index),
				Value:   uint32(value),
			}

			// Write
			//fmt.Println("Write:", part.String())
			bw := bit.NewWriter()
			Write(bw, part)

			// Read
			tok := serial_tokenizer.NewTokenizer(bw.Data())
			readPart, err := Read(tok)
			assert.NoError(t, err)
			//fmt.Println("Read:", readPart.String())

			// Compare
			assert.Equal(t, part, readPart)

			//fmt.Println()
		}
	}
}

func TestWritePartRoundtripSubtypeList(t *testing.T) {
	var (
		maxIndex     = helpers.IntPow(2, 16) - 1
		maxListElems = 5
	)

	for index := 0; index <= maxIndex; index++ {

		for listElems := 0; listElems <= maxListElems; listElems++ {
			part := Part{
				SubType: SUBTYPE_LIST,
				Index:   uint32(index),
				Values:  nil,
			}

			for curElem := 0; curElem < listElems; curElem++ {
				// Generate a predictable value for the list element
				listVal := maxIndex - index/3 + curElem*2
				if listVal < 0 {
					listVal = 0
				}
				for listVal > maxIndex {
					// Varint max value is 2^16-1 may overflow, wrap around
					listVal -= maxIndex
				}

				part.Values = append(part.Values, uint32(listVal))
			}

			// Write
			//fmt.Println("Write:", part.String())
			bw := bit.NewWriter()
			Write(bw, part)
			//fmt.Println(bw.String())

			// Read
			tok := serial_tokenizer.NewTokenizer(bw.Data())
			readPart, err := Read(tok)
			assert.NoError(t, err)
			//fmt.Println("Read:", readPart.String())

			// Compare
			assert.Equal(t, part, readPart)

			//fmt.Println()
		}
	}
}
