package serial

import (
	"borderlands_4_serials/b4s/b85"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitBlocks(t *testing.T) {
	data, err := b85.Decode("@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00")
	if err != nil {
		t.Fatal(err)
	}

	s, bits, err := Deserialize(data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("blocks:", s)
	fmt.Println("bits:", bits)
	fmt.Println("debug:", s.String())

	blocks, parts, err := s.SplitBlocks()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("split left:", Serial(blocks).String())
	fmt.Println("split right:", Serial(parts).String())
}

func TestSortBlocks(t *testing.T) {
	tests := []struct {
		name           string
		serialBefore   string
		partsStrBefore string
		partsStrAfter  string
		serialAfter    string
	}{
		{
			name:           "L50 Legendary Cooking Ambushing Truck",
			serialBefore:   "@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			partsStrBefore: "24, 0, 1, 50| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16} {25} {44} {49}|",
			partsStrAfter:  "24, 0, 1, 50| 2, 3379|| {76} {2} {3} {16} {25} {44} {49} {57} {59} {60} {75}|",
			serialAfter:    "@Ugy3L+2}TYg%$yC%i7M2g!l34$a-qhd=ArJP@}X`b00",
		},
		{
			name:           "Knife 4 skill",
			serialBefore:   "@Ugr$WBm/$!m!X=5&qXxA;nj3OOD#<4R",
			partsStrBefore: "267, 0, 1, 22| 2, 274|| {7} {1} {245:[23 39 69 79]}|",
			partsStrAfter:  "267, 0, 1, 22| 2, 274|| {7} {1} {245:[23 39 69 79]}|",
			serialAfter:    "@Ugr$WBm/$!m!X=5&qXxA;nj3OOD#<4R",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serialBefore)
			assert.NoError(t, err)

			s, bits, err := Deserialize(data)
			assert.NoError(t, err)
			fmt.Println("blocks:", s)
			fmt.Println("bits:", bits)
			fmt.Println("debug:", s.String())
			assert.Equal(t, tt.partsStrBefore, s.String())

			err = s.Sort()
			assert.NoError(t, err)
			fmt.Println("sorted:", s.String())
			assert.Equal(t, tt.partsStrAfter, s.String())

			data = Serialize(s)
			reserialized := b85.Encode(data)
			fmt.Println("reserialized:", reserialized)
			assert.Equal(t, tt.serialAfter, reserialized)
		})
	}
}
