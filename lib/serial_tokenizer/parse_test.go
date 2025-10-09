package serial_tokenizer

import (
	"borderlands_4_serials/lib/b85"
	"encoding/json"
	"fmt"
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
		name     string
		serial   string
		expected string
	}{
		{
			"ORIGINAL L50 Legendary Cooking Ambushing Truck",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			"1",
		},
		{
			"BROKEN TRUCK 1",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!XN+",
			"1",
		},
		{
			"BROKEN TRUCK 2",
			"@Ugy3L+2}TYg93=h!/NsC0/Nl@<RG/)a6EzQ&4/NX}1_1",
			"1",
		},
		{
			"BROKEN TRUCK 3",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!XN+",
			"1",
		},
		{
			"Knife 1 skill",
			"@Ugr$WBm/$!m!X=5&qXq#",
			"1",
		},
		{
			"Knife 2 skill",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3Nj00",
			"1",
		},
		{
			"Knife 3 skill",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3OODgg",
			"1",
		},
		{
			"Knife 4 skill",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3OOD#<4R",
			"1",
		},
		{
			"Common Unseen Xiuhcoatl",
			"@Ugy3L+2}TYgOyvyviz?KiBDJYGs9dOW2m",
			"1",
		},
		{
			"Green Unseen Xiuhcoatl",
			"@Ugy3L+2}TMcjNb(cjVjck8WpL1s7>WTg+kRrl/uj",
			"1",
		},
		{
			"Purple Looming Xiuhcoatl",
			"@Ugy3L+2}TYg4BQJUjVjck61AvE^+Sb3b!rZ(7U~=V",
			"1",
		},
		{
			"Top Square Simple 1",
			"@Ugy3L+2}TYgOyvyviz?KiBDJYGs9dOW2m",
			"1",
		},
		{
			"Top Square Simple 2",
			"@Ugy3L+2}TYgjMogxi7Hg07IhPq4>b?9sX3@zs9y*",
			"1",
		},
		{
			"Top Square Simple 3",
			"@Ugy3L+2}TYg4BQJUjVjck61AvE^+Sb3b!rZ(7U~=V",
			"1",
		},
		{
			"Top Square Simple 4",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			"1",
		},
		{
			"Top Square Simple 5",
			"@Ugy3L+2}TYgT#^cvMir`2hg#I5@}cgb=Ak+@2XzZ/4gm",
			"1",
		},
		{
			"Side Long Smooth  1",
			"@Ugy3L+2}TYgOyvyviz?KiBDJYKs9dOW2m",
			"1",
		},
		{
			"Side Long Smooth  2",
			"@Ugy3L+2}TYgjMogxi7Hg07IhPq4>b?9sXeG%s9y*",
			"1",
		},
		{
			"Side Long Smooth  3",
			"@Ugy3L+2}TYg4BQJUjVjck61AvE^+Sb3b!rc)7U~=V",
			"1",
		},
		{
			"Side Long Smooth  4",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l3q`a-qf{00",
			"1",
		},
		{
			"Side Long Smooth  5",
			"@Ugy3L+2}TYgT#^cvMir`2hg#I5@}cgb=Ak-u2XzZ/4gm",
			"1",
		},
		{
			"test debug",
			"@Ugy3L+2}TYg%$yC%i7Es",
			"1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			tok := NewTokenizer(data)
			output, err := tok.Parse()
			assert.NoError(t, err)

			prettyJson, err := json.MarshalIndent(output, "", "  ")
			assert.NoError(t, err)
			fmt.Println("Output:", string(prettyJson))
			fmt.Println()
		})
	}
}
