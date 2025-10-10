package serial_tokenizer

import (
	"borderlands_4_serials/lib/b85"
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

func TestSerialTokenize1(t *testing.T) {
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
		//t.Run(tt.name, func(t *testing.T) {
		data, err := b85.Decode(tt.serial)
		assert.NoError(t, err)

		fmt.Println("Name:", tt.name)
		fmt.Println("Serial:", tt.serial)
		tok := NewTokenizer(data)
		debugOutput, err := tok.Parse()
		assert.NoError(t, err)
		fmt.Println("Result:", debugOutput)
		fmt.Println()
		fmt.Println()
		//})
	}
}

func TestSerialTokenizeVexClassMods(t *testing.T) {
	var tests = []struct {
		name     string
		serial   string
		expected string
	}{
		{
			"1 arms    1 bottle",
			"@Ug!pHG38o5YT`HzQ)h-nP",
			"1",
		},
		{
			"1 cat     1 bottle",
			"@Ug!pHG38o5YZ7QZg)h-nP",
			"1",
		},
		{
			"1 skullR  1 bottle",
			"@Ug!pHG38o5YOe&^9)h-nP",
			"1",
		},
		{
			"1 bullet  1 bottle",
			"@Ug!pHG38o6@O)92A)h-nP",
			"1",
		},
		{
			"1 square  1 bottle",
			"@Ug!pHG38o5YPb#KC)h-nP",
			"1",
		},
		{
			"1 skullG  1 bottle",
			"@Ug!pHG38o5YMJlF2)h-nP",
			"1",
		},
		{
			"1 feet    1 bottle",
			"@Ug!pHG38o4tO)92A)h-nP",
			"1",
		},
		{
			"1 empty   1 bottle",
			"@Ug!pHG38o5Y4JxKV)h-nP",
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
			debugOutput, err := tok.Parse()
			assert.NoError(t, err)
			fmt.Println("Result:", debugOutput)
		})
	}
}

func TestSerialTokenizeShieldFirmware(t *testing.T) {
	var tests = []struct {
		name     string
		serial   string
		expected string
	}{
		{
			"shield 1",
			"@Uge8^+m/)}}!c178NkyuCbwKf>IWYh",
			"1",
		},
		{
			"shield 2",
			"@Uge8^+m/)}}!axR1DpKvM1BxF_41oav",
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
			debugOutput, err := tok.Parse()
			assert.NoError(t, err)
			fmt.Println("Result:", debugOutput)
		})
	}
}

func TestSerialBordel(t *testing.T) {
	var tests = []struct {
		name     string
		serial   string
		expected string
	}{
		{
			"occulted ephemeris",
			"@Ugw$Yw2}TYg4Ali1jVjck64j`oLytO%+EgA?C{!)fIMna}",
			"1",
		},
		{
			"looming balor",
			"@Ugd_t@Fme!KdTvl?RG/_Tse7ors5+=wsFVl",
			"1",
		},
		{
			"double creme omnibore",
			"@Uge(J0Fme!Kux-$2RG}7is6<7oB&t$xP@zz<P`yy=5C",
			"1",
		},
		{
			"shield 1",
			"@Uge8^+m/)}}!c178NkyuCbwKf>IWYh",
			"1",
		},
		{
			"shield 2",
			"@Uge8^+m/)}}!axR1DpKvM1BxF_41oav",
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
			"ORIGINAL L50 Legendary Cooking Ambushing Truck",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
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
			debugOutput, err := tok.Parse()
			assert.NoError(t, err)
			fmt.Println("Result:", debugOutput)
		})
	}
}
