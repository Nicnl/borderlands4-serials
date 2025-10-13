package part

import (
	"borderlands_4_serials/lib/helpers"
	"borderlands_4_serials/lib/serial_tokenizer"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadPart(t *testing.T) {
	type test struct {
		bits     string
		expected Part
	}

	testsGroups := []struct {
		name  string
		tests []test
	}{
		{
			name: "Type simple, one block",
			tests: []test{
				{
					bits: "00000 0:10",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_NONE,
					},
				},
				{
					bits: "10000 0:10",
					expected: Part{
						Index:   1,
						SubType: SUBTYPE_NONE,
					},
				},
				{
					bits: "01000 0:10",
					expected: Part{
						Index:   2,
						SubType: SUBTYPE_NONE,
					},
				},
				{
					bits: "00100 0:10",
					expected: Part{
						Index:   4,
						SubType: SUBTYPE_NONE,
					},
				},
				{
					bits: "00010 0:10",
					expected: Part{
						Index:   8,
						SubType: SUBTYPE_NONE,
					},
				},
			},
		},

		{
			name: "Type simple, two blocks",
			tests: []test{
				{
					bits: "00001-10000 0:10",
					expected: Part{
						Index:   16,
						SubType: SUBTYPE_NONE,
					},
				},
				{
					bits: "00001-01000 0:10",
					expected: Part{
						Index:   32,
						SubType: SUBTYPE_NONE,
					},
				},
				{
					bits: "00001-00100 0:10",
					expected: Part{
						Index:   64,
						SubType: SUBTYPE_NONE,
					},
				},
				{
					bits: "00001-00010 0:10",
					expected: Part{
						Index:   128,
						SubType: SUBTYPE_NONE,
					},
				},
			},
		},

		{
			name: "Type int 1block + 1block",
			tests: []test{
				{
					bits: "00000 1:00000 0:00",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_INT,
						Value:   0,
					},
				},
				{
					bits: "10000 1:00000 0:00",
					expected: Part{
						Index:   1,
						SubType: SUBTYPE_INT,
						Value:   0,
					},
				},
				{
					bits: "01000 1:00000 0:00",
					expected: Part{
						Index:   2,
						SubType: SUBTYPE_INT,
						Value:   0,
					},
				},
				{
					bits: "00100 1:00000 0:00",
					expected: Part{
						Index:   4,
						SubType: SUBTYPE_INT,
						Value:   0,
					},
				},
				{
					bits: "00010 1:00000 0:00",
					expected: Part{
						Index:   8,
						SubType: SUBTYPE_INT,
						Value:   0,
					},
				},
				{
					bits: "00000 1:10000 0:00",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_INT,
						Value:   1,
					},
				},
				{
					bits: "00000 1:01000 0:00",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_INT,
						Value:   2,
					},
				},
				{
					bits: "00000 1:00100 0:00",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_INT,
						Value:   4,
					},
				},
				{
					bits: "00000 1:00010 0:00",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_INT,
						Value:   8,
					},
				},
				{
					bits: "00010 1:10000 0:00",
					expected: Part{
						Index:   8,
						SubType: SUBTYPE_INT,
						Value:   1,
					},
				},
				{
					bits: "00100 1:01000 0:00",
					expected: Part{
						Index:   4,
						SubType: SUBTYPE_INT,
						Value:   2,
					},
				},
				{
					bits: "01000 1:00100 0:00",
					expected: Part{
						Index:   2,
						SubType: SUBTYPE_INT,
						Value:   4,
					},
				},
				{
					bits: "10000 1:00010 0:00",
					expected: Part{
						Index:   1,
						SubType: SUBTYPE_INT,
						Value:   8,
					},
				},
			},
		},
		{
			name: "Type list 1block + 1 number",
			tests: []test{
				{
					bits: "00000 0:01 [01 100-00000 00] 0:00",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{0},
					},
				},
				{
					bits: "00000 0:01 [01 100-10000 00] 0:00",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{1},
					},
				},
				{
					bits: "00000 0:01 [01 100-01000 00] 0:00",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{2},
					},
				},
				{
					bits: "00000 0:01 [01 100-00100 00] 0:00",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{4},
					},
				},
				{
					bits: "00000 0:01 [01 100-00010 00] 0:00",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{8},
					},
				},
				{
					bits: "00000 0:01 [01 100-00001-10000 00] 0:00",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{16},
					},
				},
				{
					bits: "00000 0:01 [01 100-00001-01000 00] 0:00",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{32},
					},
				},
				{
					bits: "00000 0:01 [01 100-00001-00100 00] 0:00",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{64},
					},
				},
				{
					bits: "00000 0:01 [01 100-00001-00010 00] 0:00",
					expected: Part{
						Index:   0,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{128},
					},
				},
				{
					bits: "10000 0:01 [01 100-00001-00010 00] 0:00",
					expected: Part{
						Index:   1,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{128},
					},
				},
				{
					bits: "01000 0:01 [01 100-00001-00010 00] 0:00",
					expected: Part{
						Index:   2,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{128},
					},
				},
				{
					bits: "00100 0:01 [01 100-00001-00010 00] 0:00",
					expected: Part{
						Index:   4,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{128},
					},
				},
				{
					bits: "00010 0:01 [01 100-00001-00010 00] 0:00",
					expected: Part{
						Index:   8,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{128},
					},
				},
				{
					bits: "00001-10000 0:01 [01 100-00001-00010 00] 0:00",
					expected: Part{
						Index:   16,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{128},
					},
				},
				{
					bits: "00001-01000 0:01 [01 100-00001-00010 00] 0:00",
					expected: Part{
						Index:   32,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{128},
					},
				},
				{
					bits: "00001-00100 0:01 [01 100-00001-00010 00] 0:00",
					expected: Part{
						Index:   64,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{128},
					},
				},
				{
					bits: "00001-00010 0:01 [01 100-00001-00010 00] 0:00",
					expected: Part{
						Index:   128,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{128},
					},
				},
			},
		},
		{
			name: "Type list 1block + multi^me",
			tests: []test{
				{
					bits: "11110 0:01 [01 00] 0:00",
					expected: Part{
						Index:   15,
						SubType: SUBTYPE_LIST,
						Values:  nil,
					},
				},
				{
					bits: "11110 0:01 [01 100-00000 00] 0:00",
					expected: Part{
						Index:   15,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{0},
					},
				},
				{
					bits: "11110 0:01 [01 100-10000 100-00000 00] 0:00",
					expected: Part{
						Index:   15,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{1, 0},
					},
				},
				{
					bits: "11110 0:01 [01 100-01000 100-10000 100-00000 00] 0:00",
					expected: Part{
						Index:   15,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{2, 1, 0},
					},
				},
				{
					bits: "11110 0:01 [01 100-00100 100-01000 100-10000 100-00000 00] 0:00",
					expected: Part{
						Index:   15,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{4, 2, 1, 0},
					},
				},
				{
					bits: "11110 0:01 [01 100-00010 100-00100 100-01000 100-10000 100-00000 00] 0:00",
					expected: Part{
						Index:   15,
						SubType: SUBTYPE_LIST,
						Values:  []uint32{8, 4, 2, 1, 0},
					},
				},
			},
		},
	}

	for _, testGroup := range testsGroups {
		t.Run(testGroup.name, func(t *testing.T) {
			for _, tt := range testGroup.tests {
				t.Run(tt.bits+"__"+fmt.Sprintf("%+v", tt.expected), func(t *testing.T) {
					data := helpers.BinToBytes(tt.bits)
					tok := serial_tokenizer.NewTokenizer(data)
					part, err := Read(tok)
					if err != nil {
						t.Fatalf("unexpected error: %v", err)
					}
					assert.Equal(t, tt.expected.Index, part.Index)
					assert.Equal(t, tt.expected.SubType, part.SubType)
					assert.Equal(t, tt.expected.Value, part.Value)
					assert.Equal(t, tt.expected.Values, part.Values)
				})
			}
		})

	}
}
