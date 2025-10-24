package serial

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/lib/bit"
	"borderlands_4_serials/lib/byte_mirror"
	"borderlands_4_serials/lib/helpers"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMesCouilles(t *testing.T) {
	binary := "0010001100000100100011111101100000101001010111010110001111110111011011000011111111010000110101011011111010011011101001111100111010011011101100100111000011010011110011111111101011001110010110100111110011100101111001011101101100001110100111101001111001111100111100101111110110111011"

	bytes := helpers.BinToBytes(binary)
	br := bit.NewReader(bytes)

	for i := 0; i < len(binary)/7; i++ {
		val, ok := br.ReadN(7)
		if !ok {
			fmt.Println("Failed to read bits")
			return
		}

		val = byte_mirror.GenericMirror(val, 7)

		// to ascii character
		fmt.Printf("%c", val)
	}

}

func TestMesCouilles2(t *testing.T) {
	binary := "0010001100000100100011111101100000101001010111010000011110000110100111001011111111010100011100001101001110100111101001100110111111101000011010001101111101011001110010110100111110011100101111001011101101100001110100111101001111001111100111100101111110110111011"

	bytes := helpers.BinToBytes(binary)
	br := bit.NewReader(bytes)

	for i := 0; i < len(binary)/7; i++ {
		val, ok := br.ReadN(7)
		if !ok {
			fmt.Println("Failed to read bits")
			return
		}

		val = byte_mirror.GenericMirror(val, 7)

		// to ascii character
		fmt.Printf("%c", val)
	}

}

func TestDeserializeItemSet1(t *testing.T) {
	var tests = []struct {
		name      string
		serial    string
		expected  string
		bitstream string
	}{
		{
			"ORIGINAL L50 Legendary Cooking Ambushing Truck",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			"24, 0, 1, 50| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16} {25} {44} {49}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100110011100110110  00  00  1010011100100010  10101000010  10111000010  1011101100100010  1011001111000010  1010011111000010  1011101111000010  1010000110000010  1011001110000010  1010011101000010  1011000111000010  00  00  00  00  0",
		},
		{
			"ORIGINAL L50 Legendary Cooking Ambushing Truck SMALL",
			"@Ugy3L+35F42=4?<-RG/)a6EzQ&4/NX}1~mtj3pEY_",
			"24, 0, 1, 12| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16} {25} {44} {49}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  10000110  00  10001000  01  100110011100110110  00  00  1010011100100010  10101000010  10111000010  1011101100100010  1011001111000010  1010011111000010  1011101111000010  1010000110000010  1011001110000010  1010011101000010  1011000111000010  00  00  00  ",
		},
		{
			"BROKEN TRUCK 1",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!XN+",
			"24, 0, 1, 50| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100110011100110110  00  00  1010011100100010  10101000010  10111000010  1011101100100010  1011001111000010  1010011111000010  1011101111000010  1010000110000010  00  00  00  00  0",
		},
		{
			"BROKEN TRUCK 3",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!XN+",
			"24, 0, 1, 50| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100110011100110110  00  00  1010011100100010  10101000010  10111000010  1011101100100010  1011001111000010  1010011111000010  1011101111000010  1010000110000010  00  00  00  00  0",
		},
		{
			"Knife 1 skill",
			"@Ugr$WBm/$!m!X=5&qXq#",
			"267, 0, 1, 22| 2, 274|| {7} {1}|",
			"0010000  11010010110100001  01  10000000  01  10010000  01  1000110110000  00  10001000  01  11010010010010001  00  00  10111100010  10110000010  00  00  00  ",
		},
		{
			"Knife 2 skill",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3Nj00",
			"267, 0, 1, 22| 2, 274|| {7} {1} {245:[23 39]}|",
			"0010000  11010010110100001  01  10000000  01  10010000  01  1000110110000  00  10001000  01  11010010010010001  00  00  10111100010  10110000010  1011010111110001  01  1001110110000  1001110101000  00  00  00  00  00  ",
		},
		{
			"Knife 3 skill",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3OODgg",
			"267, 0, 1, 22| 2, 274|| {7} {1} {245:[23 39 69]}|",
			"0010000  11010010110100001  01  10000000  01  10010000  01  1000110110000  00  10001000  01  11010010010010001  00  00  10111100010  10110000010  1011010111110001  01  1001110110000  1001110101000  1001010100100  00  00  0",
		},
		{
			"Knife 4 skill",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3OOD#<4R",
			"267, 0, 1, 22| 2, 274|| {7} {1} {245:[23 39 69 79]}|",
			"0010000  11010010110100001  01  10000000  01  10010000  01  1000110110000  00  10001000  01  11010010010010001  00  00  10111100010  10110000010  1011010111110001  01  1001110110000  1001110101000  1001010100100  1001111100100  00  00  00  00  ",
		},
		{
			"Common Unseen Xiuhcoatl",
			"@Ugy3L+2}TYgOyvyviz?KiBDJYGs9dOW2m",
			"24, 0, 1, 50| 2, 3269|| {95} {2} {3} {8} {13} {25} {44} {50}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100101010011100110  00  00  1011111110100010  10101000010  10111000010  10100010010  10110110010  1011001110000010  1010011101000010  1010100111000010  00  0",
		},
		{
			"Green Unseen Xiuhcoatl",
			"@Ugy3L+2}TMcjNb(cjVjck8WpL1s7>WTg+kRrl/uj",
			"24, 0, 1, 34| 2, 1470|| {96} {2} {3} {6} {8} {61} {13} {28} {32} {42} {50}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100101000  00  10001000  01  100011111101110100  00  00  1010000101100010  10101000010  10111000010  10101100010  10100010010  1011011111000010  10110110010  1010011110000010  1010000101000010  1010101101000010  1010100111000010  00  00  00  00  ",
		},
		{
			"Purple Looming Xiuhcoatl",
			"@Ugy3L+2}TYg4BQJUjVjck61AvE^+Sb3b!rZ(7U~=V",
			"24, 0, 1, 50| 2, 2992|| {98} {2} {3} {4} {5} {8} {62} {64} {14} {25} {42} {51}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100000011101111010  00  00  1010100101100010  10101000010  10111000010  10100100010  10110100010  10100010010  1010111111000010  1010000100100010  10101110010  1011001110000010  1010101101000010  1011100111000010  00  00  0",
		},
		{
			"Top Square Simple 1",
			"@Ugy3L+2}TYgOyvyviz?KiBDJYGs9dOW2m",
			"24, 0, 1, 50| 2, 3269|| {95} {2} {3} {8} {13} {25} {44} {50}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100101010011100110  00  00  1011111110100010  10101000010  10111000010  10100010010  10110110010  1011001110000010  1010011101000010  1010100111000010  00  0",
		},
		{
			"Top Square Simple 2",
			"@Ugy3L+2}TYgjMogxi7Hg07IhPq4>b?9sX3@zs9y*",
			"24, 0, 1, 50| 2, 3246|| {76} {2} {4} {5} {75} {60} {57} {13} {25} {44} {47}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100011110101100110  00  00  1010011100100010  10101000010  10100100010  10110100010  1011101100100010  1010011111000010  1011001111000010  10110110010  1011001110000010  1010011101000010  1011111101000010  00  0",
		},
		{
			"Top Square Simple 3",
			"@Ugy3L+2}TYg4BQJUjVjck61AvE^+Sb3b!rZ(7U~=V",
			"24, 0, 1, 50| 2, 2992|| {98} {2} {3} {4} {5} {8} {62} {64} {14} {25} {42} {51}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100000011101111010  00  00  1010100101100010  10101000010  10111000010  10100100010  10110100010  10100010010  1010111111000010  1010000100100010  10101110010  1011001110000010  1010101101000010  1011100111000010  00  00  0",
		},
		{
			"Top Square Simple 4",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			"24, 0, 1, 50| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16} {25} {44} {49}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100110011100110110  00  00  1010011100100010  10101000010  10111000010  1011101100100010  1011001111000010  1010011111000010  1011101111000010  1010000110000010  1011001110000010  1010011101000010  1011000111000010  00  00  00  00  0",
		},
		{
			"Top Square Simple 5",
			"@Ugy3L+2}TYgT#^cvMir`2hg#I5@}cgb=Ak+@2XzZ/4gm",
			"24, 0, 1, 50| 2, 338|| {98} {2} {6} {3} {5} {7} {60} {59} {57} {14} {25} {43} {51}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  11010010010010101  00  00  1010100101100010  10101000010  10101100010  10111000010  10110100010  10111100010  1010011111000010  1011101111000010  1011001111000010  10101110010  1011001110000010  1011101101000010  1011100111000010  00  00  00  ",
		},
		{
			"Side Long Smooth  1",
			"@Ugy3L+2}TYgOyvyviz?KiBDJYKs9dOW2m",
			"24, 0, 1, 50| 2, 3269|| {95} {2} {3} {8} {13} {29} {44} {50}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100101010011100110  00  00  1011111110100010  10101000010  10111000010  10100010010  10110110010  1011011110000010  1010011101000010  1010100111000010  00  0",
		},
		{
			"Side Long Smooth  2",
			"@Ugy3L+2}TYgjMogxi7Hg07IhPq4>b?9sXeG%s9y*",
			"24, 0, 1, 50| 2, 3246|| {76} {2} {4} {5} {75} {60} {57} {13} {29} {44} {47}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100011110101100110  00  00  1010011100100010  10101000010  10100100010  10110100010  1011101100100010  1010011111000010  1011001111000010  10110110010  1011011110000010  1010011101000010  1011111101000010  00  0",
		},
		{
			"Side Long Smooth  3",
			"@Ugy3L+2}TYg4BQJUjVjck61AvE^+Sb3b!rc)7U~=V",
			"24, 0, 1, 50| 2, 2992|| {98} {2} {3} {4} {5} {8} {62} {64} {14} {29} {42} {51}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100000011101111010  00  00  1010100101100010  10101000010  10111000010  10100100010  10110100010  10100010010  1010111111000010  1010000100100010  10101110010  1011011110000010  1010101101000010  1011100111000010  00  00  0",
		},
		{
			"Side Long Smooth  4",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l3q`a-qf{00",
			"24, 0, 1, 50| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16} {29} {44} {49}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100110011100110110  00  00  1010011100100010  10101000010  10111000010  1011101100100010  1011001111000010  1010011111000010  1011101111000010  1010000110000010  1011011110000010  1010011101000010  1011000111000010  00  00  00  00  0",
		},
		{
			"Side Long Smooth  5",
			"@Ugy3L+2}TYgT#^cvMir`2hg#I5@}cgb=Ak-u2XzZ/4gm",
			"24, 0, 1, 50| 2, 338|| {98} {2} {6} {3} {5} {7} {60} {59} {57} {14} {27} {43} {51}|",
			"0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  11010010010010101  00  00  1010100101100010  10101000010  10101100010  10111000010  10110100010  10111100010  1010011111000010  1011101111000010  1011001111000010  10101110010  1011101110000010  1011101101000010  1011100111000010  00  00  00  ",
		},
		{
			"moxx 1",
			"@Uge8Usm/*1K!b~@T2!pSK9`(MRphLZnMu&VCtx^C",
			"268, 0, 1, 50| \"ft\", 1| 2, 3704|| {7} {247:77} {3} {247:[26 239 170 5]}|",
			"0010000  11010010001100001  01  10000000  01  10010000  01  1000100111000  00  1110100001100110010111  01  10010000  00  10001000  01  100000111110101110  00  00  10111100010  101111011111011011100100000  10111000010  1011110111110001  01  1000101110000  1001111101110  1000101101010  10010100  00  00  00  0",
		},
		{
			"moxx 2",
			"@Ugr$`Rm/*1K!b~@T2(#sZ9`(MRph3Nl#=Gnicv1i",
			"299, 0, 1, 50| \"ft\", 1| 2, 1835|| {7} {247:77} {1} {247:[238 91 112 5]}|",
			"0010000  11010010110101001  01  10000000  01  10010000  01  1000100111000  00  1110100001100110010111  01  10010000  00  10001000  01  100110110100111100  00  00  10111100010  101111011111011011100100000  10110000010  1011110111110001  01  1000111101110  1001101110100  1000000111100  10010100  00  00  00  0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			start := time.Now()
			blocks, bits, err := Deserialize(data)
			end := time.Now()
			fmt.Printf("Deserialization took %v\n", end.Sub(start))
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, blocks.String())
			assert.Equal(t, tt.bitstream, bits)
			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			fmt.Println("Result:", blocks.String())
			fmt.Println("Bitstream:", bits)
			fmt.Println()
			fmt.Println()
		})
	}
}

func TestDeserializeItemWithString(t *testing.T) {
	var tests = []struct {
		name      string
		serial    string
		expected  string
		bitstream string
	}{
		{
			"1",
			"@Ugw$Yw5hh0gkD?i5ey2^>_}zwoVYM&2d@j4mVR`;*PRq;pw/rfp7Immb4H@o3Gsyf-n=W`*PXCI-a(S8kg~`6=w7h(O%h$z3b?Of46si>J7U~>o2nPT",
			"21, 0, 2, 537|| \"MAL_SM.comp_05_legendary_firework\" {2} {5} {3} {6} {1:12} \"MAL_SM.part_barrel_02_firework\" {72} {14} {27} {35} {34} {43} {51} {1:48}|",
			"0010000  1001010110000  01  10000000  01  10001000  01  100100111000101000  00  00  1111000101000101100110000010011001111110111001011011001011101011000111111011101101100001111111101000011010101101111101001101110100111110011101001101110110010011100001101001111001111111110101100111001011010011110100111110111111101101001111101011  10101000010  10110100010  10111000010  10101100010  10110000100110000  1110111110000101100110000010011001111110111001011011001011101000001111000011010011100101111111101010001110000110100111010011110100110011011111110100001100100110111110101100111001011010011110100111110111111101101001111101011  1010001100100010  10101110010  1011101110000010  1011100101000010  1010100101000010  1011101101000010  1011100111000010  1011000010000111000000  00  00  00  0",
		},
		{
			"2",
			"@Uge8jxm/%P$!X^K3ImB=I`Foa^-C=V1dk&M&Wx8wpZo/K@+Lv8Em)+#BJpZN3!/r)IJ)fuF?h3WOH{f)L-}3YKEHAsm<ns3%CZEf67d$Mdf8pUg`yD2K>+W>CUjEYMVfVb9p3l>7ck83UBmn>",
			"278, 0, 1, 30| 2, 510|| \"borg_grenade_gadget.comp_05_legendary_transmission\" {2} {245:23} \"borg_grenade_gadget.part_payload_unique_transmission\" {245:[72 1]}|",
			"0010000  11010010011010001  01  10000000  01  10010000  01  1000111110000  00  10001000  01  11010010011111111  00  00  111010011100001000111111011010011111100111111101111001101001111010011011101110000110010011101001111111011110011100001100100111110011101001100101110111010110001111110111011011000011111111010000110101011011111010011011101001111100111010011011101100100111000011010011110011111111101001011101001111000011011101111001111011011100101111001111100111100101111110110111011  10101000010  101101011111011110110000000  11100101110000100011111101101001111110011111110111100110100111101001101110111000011001001110100111111101111001110000110010011111001110100110010111011101000001111000011010011100101111111101000011110000111001111001101111110111000011001001111111011010111011101110010111000111101011110100111111101001011101001111000011011101111001111011011100101111001111100111100101111110110111011  1011010111110001  01  1000001100100  10010000  00  00  00  00  00  0",
		},
		{
			"3",
			"@Ugfs(8Fme!KG*)mdh(Y2vD0PkBZTJ^f`?AaDvYQ;1=fCE(d`?fd;qvl)o_@O<RH6#K3SyA>4N6_`u$=xChvo7z`wNSG&1w0Zo^HeC<@r4Qc2}uMU1/<073vi#7it&k9cm<MB?1",
			"13, 0, 1, 50| 2, 2698|| \"DAD_AR.comp_05_legendary_firstimpression\" {1} {4} {2} \"DAD_AR.part_barrel_01_firstimpression\" {10} {9} {11} {25} {36} {39} {44} {45} {55} {65} {69}|",
			"0010000  10010110  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100010110001101010  00  00  11100011010000010001100000100100011111101100000101001010111010110001111110111011011000011111111010000110101011011111010011011101001111100111010011011101100100111000011010011110011111111101011001110010110100111110011100101111001011101101100001110100111101001111001111100111100101111110110111011  10110000010  10100100010  10101000010  11110101010000010001100000100100011111101100000101001010111010000011110000110100111001011111111010100011100001101001110100111101001100110111111101000011010001101111101011001110010110100111110011100101111001011101101100001110100111101001111001111100111100101111110110111011  10101010010  10110010010  10111010010  1011001110000010  1010010101000010  1011110101000010  1010011101000010  1011011101000010  1011110111000010  1011000100100010  1011010100100010  00  00  0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			start := time.Now()
			blocks, bits, err := Deserialize(data)
			end := time.Now()
			fmt.Printf("Deserialization took %v\n", end.Sub(start))
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, blocks.String())
			assert.Equal(t, tt.bitstream, bits)
			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			fmt.Println("Result:", blocks.String())
			fmt.Println("Bitstream:", bits)
			fmt.Println()
			fmt.Println()
		})
	}
}

func TestDeserializeBench(t *testing.T) {
	const (
		numIterations = 25
		numDeserials  = 4000
	)

	measurements := make([]time.Duration, 0, numIterations)

	for range numIterations {
		start := time.Now()
		for range numDeserials {
			data, _ := b85.Decode("@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00")
			_, _, _ = Deserialize(data)
		}
		duration := time.Since(start)
		fmt.Println("4000:", duration)
		measurements = append(measurements, duration)
	}

	total := time.Duration(0)
	for _, d := range measurements {
		total += d
	}
	avg := total / time.Duration(len(measurements))
	fmt.Println("Average for 4000 deserializations:", avg)
}

var skinTests = []struct {
	name         string
	serial       string
	deserialized string
}{
	{
		"ORIGINAL L50 Lurking Cuca",
		"@UgbV{rFme!KI4sa#RG}W#sX3@xsFnx",
		`3, 0, 1, 50| 2, 1292|| {95} {2} {7} {14} {25} {42} {70}|`,
	},
	{
		"Skin: Solar Flair",
		"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vGE1",
		`3, 0, 1, 50| 2, 1292|| {95} {2} {7} {14} {25} {42} {70}| "c", 110|`,
	},
	{
		"Skin: Carcade Shooter",
		"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vB?G",
		`3, 0, 1, 50| 2, 1292|| {95} {2} {7} {14} {25} {42} {70}| "c", 105|`,
	},
	{
		"Skin: Itty Bitty Kitty Committee",
		"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vB3r",
		`3, 0, 1, 50| 2, 1292|| {95} {2} {7} {14} {25} {42} {70}| "c", 104|`,
	},
	{
		"Skin: With the grain",
		"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vGD}",
		`3, 0, 1, 50| 2, 1292|| {95} {2} {7} {14} {25} {42} {70}| "c", 94|`,
	},
	{
		"Skin: The System",
		"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vAG2",
		`3, 0, 1, 50| 2, 1292|| {95} {2} {7} {14} {25} {42} {70}| "c", 87|`,
	},
	{
		"Skin: Devourer",
		"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_v9Sd",
		`3, 0, 1, 50| 2, 1292|| {95} {2} {7} {14} {25} {42} {70}| "c", 86|`,
	},
	{
		"Skin: Soused",
		"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_v5^G",
		`3, 0, 1, 50| 2, 1292|| {95} {2} {7} {14} {25} {42} {70}| "c", 82|`,
	},
	{
		"Skin: Bird of Prey",
		"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_v55r",
		`3, 0, 1, 50| 2, 1292|| {95} {2} {7} {14} {25} {42} {70}| "c", 81|`,
	},
	{
		"Skin: Eternal Defender",
		"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vH1i",
		`3, 0, 1, 50| 2, 1292|| {95} {2} {7} {14} {25} {42} {70}| "c", 79|`,
	},
	{
		"Skin: Weirdo",
		"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vGD`",
		`3, 0, 1, 50| 2, 1292|| {95} {2} {7} {14} {25} {42} {70}| "c", 78|`,
	},
	{
		"Skin: Smiley",
		"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vFQW",
		`3, 0, 1, 50| 2, 1292|| {95} {2} {7} {14} {25} {42} {70}| "c", 77|`,
	},
}

func TestDeserializeSkins(t *testing.T) {
	for _, tt := range skinTests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			blocks, bits, err := Deserialize(data)
			assert.NoError(t, err)
			assert.Equal(t, tt.deserialized, blocks.String())
			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			fmt.Println("Result:", blocks.String())
			fmt.Println("Bitstream:", bits)
			fmt.Println()
			fmt.Println()
		})
	}
}

func TestDeserializeVexClassMods(t *testing.T) {
	var tests = []struct {
		name     string
		serial   string
		expected string
	}{
		{
			"1 arms    1 bottle",
			"@Ug!pHG38o5YT`HzQ)h-nP",
			"254, 0, 1, 37|| {181} {328} {234:36}|",
		},
		{
			"1 cat     1 bottle",
			"@Ug!pHG38o5YZ7QZg)h-nP",
			"254, 0, 1, 37|| {213} {328} {234:36}|",
		},
		{
			"1 skullR  1 bottle",
			"@Ug!pHG38o5YOe&^9)h-nP",
			"254, 0, 1, 37|| {145} {328} {234:36}|",
		},
		{
			"1 bullet  1 bottle",
			"@Ug!pHG38o6@O)92A)h-nP",
			"254, 0, 1, 37|| {151} {328} {234:36}|",
		},
		{
			"1 square  1 bottle",
			"@Ug!pHG38o5YPb#KC)h-nP",
			"254, 0, 1, 37|| {157} {328} {234:36}|",
		},
		{
			"1 skullG  1 bottle",
			"@Ug!pHG38o5YMJlF2)h-nP",
			"254, 0, 1, 37|| {133} {328} {234:36}|",
		},
		{
			"1 feet    1 bottle",
			"@Ug!pHG38o4tO)92A)h-nP",
			"254, 0, 1, 37|| {148} {328} {234:36}|",
		},
		{
			"1 empty   1 bottle",
			"@Ug!pHG38o5Y4JxKV)h-nP",
			"254, 0, 1, 37|| {21} {328} {234:36}|",
		},
		{
			"1 Rarms   1 Bbottle  + 1Gboom",
			"@Ug!pHG38o5YT`HzQ)k4)S6#x",
			"254, 0, 1, 37|| {181} {328} {42} {234:36}|",
		},
		{
			"1 Rarms",
			"@Ug!pHG38o5YU8;7e00",
			"254, 0, 1, 37|| {181} {234:36}|",
		},
		{
			"1 Rarms 2 bottles",
			"@Ug!pHG38o5YT`HzQ#Wbker2+r",
			"254, 0, 1, 37|| {181} {328} {328} {234:36}|",
		},
		{
			"2 Rarms 1 bottles",
			"@Ug!pHG38o5YU20t_ra{#%6#x",
			"254, 0, 1, 37|| {181} {181} {328} {234:36}|",
		},
		{
			"melee (red skin)",
			"@Ug!pHG38o6DcBud",
			"254, 0, 1, 37|| {234:36}|",
		},
		{
			"melee (white skin)",
			"@Ug!pHG38o6DP_;`100",
			"254, 0, 1, 37|| {2} {234:36}|",
		},
		{
			"1 arm no melee",
			"@Ug!pHG38o5YT>=",
			"254, 0, 1, 37|| {181}|",
		},
		{
			"1 arm 1 bottle + firmware jacked",
			"@Ug!pHG38o5YT`HzQ)$V@)",
			"254, 0, 1, 37|| {181} {328} {234:255}|",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			blocks, bits, err := Deserialize(data)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, blocks.String())
			fmt.Println("Result:", blocks.String())
			fmt.Println("Bits:", bits)
		})
	}
}

func TestDeserializeFirmware(t *testing.T) {
	var tests = []struct {
		name     string
		serial   string
		expected string
	}{
		{
			"shield 1",
			"@Uge8^+m/)}}!c178NkyuCbwKf>IWYh",
			"300, 0, 1, 50| 2, 1283|| {9} {8} {246:26} {248:[7 6]}|",
		},
		{
			"shield 2",
			"@Uge8^+m/)}}!axR1DpKvM1BxF_41oav",
			"300, 0, 1, 50| 2, 192|| {9} {8} {246:26} {248:[17 16]}|",
		},
		{
			"rapid swarm + gadget ahoy",
			"@Uge8aum/)}}!qkqSNDXRzG&iINder)8E{Op",
			"272, 0, 1, 50| 2, 2261|| {8} {1} {245:[23 29]} {7} {245:[71 6]}|",
		},
		{
			"waterfall grenade + airstrike",
			"@Ugr$)Nm/)}}!YpV~ky;-O59uLV#F7vI",
			"291, 0, 1, 50| 2, 11|| {9} {8} {245:[28 29 42 70 4]}|",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			blocks, bits, err := Deserialize(data)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, blocks.String())
			fmt.Println("Result:", blocks.String())
			fmt.Println("Bits:", bits)
		})
	}
}

func TestDeserializeRandomItems(t *testing.T) {
	var tests = []struct {
		name     string
		serial   string
		expected string
	}{
		{
			"occulted ephemeris",
			"@Ugr$!Lm/)}}!u<K5M>VQ_G&h6`+T9-j",
			"287, 0, 1, 50| 2, 3903|| {7} {6} {246:[23 44]} {237:31}|",
		},
		{
			"looming balor",
			"@Ugd_t@Fme!KdTvl?RG/_Tse7ors5+=wsFVl",
			"9, 0, 1, 50| 2, 3485|| {96} {2} {4} {8} {59} {92} {26} {41} {68}|",
		},
		{
			"double creme omnibore",
			"@Uge(J0Fme!Kux-$2RG}7is6<7oB&t$xP@zz<P`yy=5C",
			"11, 0, 1, 50| 2, 3432|| {97} {2} {5} {3} {4} {8} {66} {10} {23} {32} {40} {46} {53}|",
		},
		{
			"shield 1",
			"@Uge8^+m/)}}!c178NkyuCbwKf>IWYh",
			"300, 0, 1, 50| 2, 1283|| {9} {8} {246:26} {248:[7 6]}|",
		},
		{
			"shield 2",
			"@Uge8^+m/)}}!axR1DpKvM1BxF_41oav",
			"300, 0, 1, 50| 2, 192|| {9} {8} {246:26} {248:[17 16]}|",
		},
		{
			"Knife 1 skill",
			"@Ugr$WBm/$!m!X=5&qXq#",
			"267, 0, 1, 22| 2, 274|| {7} {1}|",
		},
		{
			"Knife 2 skill",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3Nj00",
			"267, 0, 1, 22| 2, 274|| {7} {1} {245:[23 39]}|",
		},
		{
			"Knife 3 skill",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3OODgg",
			"267, 0, 1, 22| 2, 274|| {7} {1} {245:[23 39 69]}|",
		},
		{
			"Knife 4 skill",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3OOD#<4R",
			"267, 0, 1, 22| 2, 274|| {7} {1} {245:[23 39 69 79]}|",
		},
		{
			"ORIGINAL L50 Legendary Cooking Ambushing Truck",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			"24, 0, 1, 50| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16} {25} {44} {49}|",
		},
		{
			"no level spawned console",
			"@Ugr$lG7-8sL(4z`<KALPY4GrpidjS",
			"277, 0, 2, 2932|| {7} {2} {243:[105 99]} {1} {243:93}|",
		},
		{
			"Vindictive Evolver",
			"@Ugr$rIm/)}}!q`oqNWCv7s8Ex7AI%h@D>DE",
			"281, 0, 1, 50| 2, 1206|| {8} {247:76} {2} {9} {247:[35 180 19]}|",
		},
		{
			"Retributive Devourer",
			"@Uge8;)m/)}}!sxA_MZGU4Xi$ZEAI&bYFAo3",
			"296, 0, 1, 50| 2, 2746|| {5} {247:76} {1} {2} {247:[91 246 15]}|",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			blocks, bits, err := Deserialize(data)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, blocks.String())
			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			fmt.Println("Result:", blocks.String())
			fmt.Println("Bits:", bits)
		})
	}
}

func TestDeserializeProblematicSerials1(t *testing.T) {
	var tests = []struct {
		name     string
		serial   string
		expected string
	}{
		{
			"(Overshielding Junker Principal) => invalid token, got 111 at position 177",
			"@Uge8^+m/)}}!hAcRNkyuCbwKeuQ2X8i00",
			"300, 0, 1, 50| 2, 879|| {9} {8} {246:26} {248:3} {246:55}|",
		},
		{
			"(Watts 4 Dinner) => invalid token, got 111 at position 149",
			"@Uge8^+m/)}}!X>8AqZ-w/IH35^oFHmn1po",
			"300, 0, 1, 50| 2, 332|| {7} {6} {246:24} {248:[7 8]} {246:10}|",
		},
		{
			"(Resonant Itchy Sparky Shield) => invalid token, got 111 at position 148",
			"@Uge92<m/)}}!hGMLL{+MNG&aNz+T9-j",
			"306, 0, 1, 50| 2, 3567|| {4} {10} {246:[22 54]} {237:31}|",
		},
		{
			"(High-Capacity Junker Cask) => invalid token, got 111 at position 177",
			"@Uge8^+m/)}}!rTKTs#5Kn1B%~++II&4",
			"300, 0, 1, 50| 2, 119|| {4} {10} {246:25} {248:7} {246:27}|",
		},
		{
			"(Resonant Dextrous Scar) => invalid token, got 111 at position 144",
			"@Ugr$oHm/)}}!gO1pMir`kaX/6gQ2WLK00",
			"279, 0, 1, 50| 2, 1389|| {6} {2} {246:24} {248:13} {246:49}|",
		},
		{
			"(Fleeting Vigorous Stanchion) => invalid token, got 111 at position 149",
			"@Ugr$oHm/)}}!t7k2N)@VobwKf(&>#Q",
			"279, 0, 1, 50| 2, 1435|| {10} {2} {246:26} {248:9} {1}|",
		},
		{
			"(Trigger-Happy Extra Medium) => invalid token, got 111 at position 149",
			"@Uge9B?m/)}}!t?`8DpKvM1BxF_3k?7",
			"312, 0, 1, 50| 2, 125|| {9} {8} {246:26} {248:[13 20]}|",
		},
		{
			"(Fleeting Anxious Super Soldier) => invalid token, got 111 at position 144",
			"@Ugr$uJm/)}}!q`oqM>VQ_Z$P^nQ2XWq00",
			"283, 0, 1, 50| 2, 1206|| {7} {6} {246:23} {237:1} {246:57}|",
		},
		{
			"(Scrapper Emollient Hoarder) => invalid token, got 111 at position 182",
			"@Ugr$uJm/)}}!bF&$M>VQ_Z9uylQ2XWq00",
			"283, 0, 1, 50| 2, 2305|| {7} {6} {246:22} {237:1} {246:57}|",
		},
		{
			"(Scrapper Emollient Hoarder) => invalid token, got 111 at position 182",
			"@Ugr$uJm/)}}!tBVPL{+MNb3nTr18UzM00",
			"283, 0, 1, 50| 2, 3099|| {4} {10} {246:25} {237:17} {246:29}|",
		},
		{
			"(Bladed Chitinous Domy) => invalid token, got 111 at position 185",
			"@Uge9B?m/)}}!f5KCL{+MNaX/5-sUb7~",
			"312, 0, 1, 50| 2, 3658|| {4} {10} {246:24} {248:[21 20]}|",
		},
		{
			"(Vigilant Bunker) => invalid token, got 111 at position 149",
			"@Ugr%Scm/)}}!a#_iM^&nQYe2i3Q2X/P8UX",
			"321, 0, 1, 50| 2, 2048|| {7} {10} {246:21} {237:9} {246:29} {6}|",
		},
		{
			"(Barrage Chitinous Firewerks) => invalid token, got 111 at position 180",
			"@Uge8^+m/)}}!gz?FNkyuCbwKf>nIVqa7X/<",
			"300, 0, 1, 50| 2, 2062|| {9} {8} {246:26} {248:[19 14]} {246:16}|",
		},
		{
			"(Watts 4 Dinner) => invalid token, got 111 at position 149",
			"@Ugr$-Om/)}}!g!mYLlvrhbwKgEQ2Ry!00",
			"293, 0, 1, 50| 2, 2414|| {3} {2} {246:26} {248:11} {246:33}|",
		},
		{
			"(Nova Pointed Protean Cell) => invalid token, got 111 at position 149",
			"@Ugr$-Om/)}}!g%VSN+qg&aX/5-DcOMr0R",
			"293, 0, 1, 50| 2, 3662|| {10} {4} {246:24} {248:[5 27]} {1}|",
		},
		{
			"(Berserkr Meandering Cindershelly) => invalid token, got 111 at position 149",
			"@Uge8^+m/)}}!dO_KNkyuCbwKf>sUeKoR{#",
			"300, 0, 1, 50| 2, 1286|| {9} {8} {246:26} {248:[21 12]} {246:2}|",
		},
		{
			"(Watts 4 Dinner) => invalid token, got 111 at position 149",
			"@Ugr$uJm/)}}!qj@8M>VQ_Ye2gjQ2XWq00",
			"283, 0, 1, 50| 2, 1877|| {7} {6} {246:21} {237:1} {246:57}|",
		},
		{
			"(Scrapper Emollient Hoarder) => invalid token, got 111 at position 182",
			"@Ugr$oHm/)}}!uW`wNfoMnbwKgEQ2XA2A^`",
			"279, 0, 1, 50| 2, 2078|| {9} {2} {246:26} {248:11} {246:55} {8}|",
		},
		{
			"(Nova Mutable Pandoran Memento) => invalid token, got 111 at position 149",
			"@Ugr$oHm/)}}!r*P7N)@VoaX/5V&>#Q",
			"279, 0, 1, 50| 2, 1720|| {10} {2} {246:24} {248:7} {1}|",
		},
		{
			"(Resonant Extra Medium) => invalid token, got 111 at position 149",
			"@Ugr$oHm/)}}!pO~_N)@Vob3pN1Q2Ww=1_1",
			"279, 0, 1, 50| 2, 3250|| {10} {2} {246:25} {248:5} {246:20} {1}|",
		},
		{
			"(Berserkr Extra Medium) => invalid token, got 111 at position 149",
			"@Uge9B?m/)}}!qjb`L{+MNb3pOiQ2XWq00",
			"312, 0, 1, 50| 2, 1717|| {4} {10} {246:25} {248:13} {246:57}|",
		},
		{
			"(Fleeting Emollient Bunker) => invalid token, got 111 at position 149",
			"@Uge8^+m/)}}!uYwMM>VQ_Z9wrGQ2TxX00",
			"300, 0, 1, 50| 2, 2878|| {7} {6} {246:22} {248:1} {246:47}|",
		},
		{
			"(Traveler Vitamin Sparky Shield) => invalid token, got 111 at position 180",
			"@Ugr$uJm/)}}!uW5XM>VQ_Z$P^nQ2XWq00",
			"283, 0, 1, 50| 2, 1790|| {7} {6} {246:23} {237:1} {246:57}|",
		},
		{
			"(Scrapper Emollient Hoarder) => invalid token, got 111 at position 182",
			"@Ugr$oHm/)}}!pu&fNfoMnb3pN<i6Mp}0R",
			"279, 0, 1, 50| 2, 1235|| {9} {2} {246:25} {248:[17 6]} {8}|",
		},
		{
			"(Amp Tolling Pandoran Memento) => invalid token, got 111 at position 149",
			"@Ugr$uJm/)}}!f216M>VQ_bwIltQ2XWq00",
			"283, 0, 1, 50| 2, 2282|| {7} {6} {246:26} {237:1} {246:57}|",
		},
		{
			"(Scrapper Emollient Hoarder) => invalid token, got 111 at position 182",
			"@Ugr$-Om/)}}!pz#BLlvrhbwKgEQ2Ry!00",
			"293, 0, 1, 50| 2, 3411|| {3} {2} {246:26} {248:11} {246:33}|",
		},
		{
			"(Nova Pointed Protean Cell) => invalid token, got 111 at position 149",
			"@Uge9B?m/)}}!X>CsqAJzCIiUCrsC}mZ",
			"312, 0, 1, 50| 2, 336|| {4} {10} {246:25} {248:1} {246:35}|",
		},
		{
			"(Traveler Weatherproof Bunker) => invalid token, got 111 at position 148",
			"@Ugr$!Lm/)}}!cfVeL{+MNYe2i318UzM00",
			"287, 0, 1, 50| 2, 3108|| {4} {10} {246:21} {237:25} {246:29}|",
		},
		{
			"(Boxer Chitinous Laminate) => invalid token, got 111 at position 185",
			"@Ugr$uJm/)}}!ljeYqZ-w/H=x}OsD1MQ",
			"283, 0, 1, 50| 2, 295|| {7} {6} {246:23} {237:1} {246:57}|",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			blocks, bits, err := Deserialize(data)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, blocks.String())
			fmt.Println("Result:", blocks.String())
			fmt.Println("Bits:", bits)
		})
	}
}

func TestDeserializeProblematicSerials2(t *testing.T) {
	var tests = []struct {
		name     string
		serial   string
		expected string
	}{
		{
			"Sho Kunai (gadget)",
			"@Ugr$WBm/)}}!bEtWObu#%Z$Os-",
			"267, 0, 1, 50| 2, 1793|| {12} {1} {245:23} {11}|",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!bmKjM-}RPG}*)&8r1p10ss",
			"290, 0, 1, 50| 2, 642|| {7} {2} {243:[105 102]} {1} {243:80}|",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!t}17M-}RPG}&Yt8r1pj0ss",
			"290, 0, 1, 50| 2, 2781|| {7} {2} {243:[105 99]} {1} {243:91}|",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!rWS*M-}RPG}(k28r1pH0ss",
			"290, 0, 1, 50| 2, 1367|| {7} {2} {243:[105 100]} {1} {243:84}|",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!gSxDM-}RPG}&Yt8r1p10ss",
			"290, 0, 1, 50| 2, 3565|| {7} {2} {243:[105 99]} {1} {243:80}|",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!f+~}M-}RPG}*)&8r1n{b_oI",
			"290, 0, 1, 50| 2, 588|| {7} {2} {243:[105 102]} {1} {243:[91 8]}|",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!u(vIM-}RPG}&Yt8r1n{aLED",
			"290, 0, 1, 50| 2, 1439|| {7} {2} {243:[105 99]} {1} {243:[88 11]}|",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!bIGlM-}RPG}&Yt8r1n{Y{?7&",
			"290, 0, 1, 50| 2, 3457|| {7} {2} {243:[105 99]} {1} {243:[86 19]}|",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!qBXsM-}RPG}(k28r1n{WC;Q",
			"290, 0, 1, 50| 2, 2708|| {7} {2} {243:[105 100]} {1} {243:[82 8]}|",
		},
		{
			"Ballistic Bonn-91 (weapon)",
			"@Uga`vnFme!Kq<v6nRG}8tsG&oVnu!XD8i/_J7}Pw}9aJgQDby^~IMg`=\"",
			"2, 0, 1, 50| 2, 3938|| {98} {2} {5} {6} {1:14} {8} {73} {64} {65} {9} {17} {57} {27} {34} {35} {41} {49} {51}|",
		},
		{
			"discord gun",
			"@UgxFw!2}TYgjMN5-iz-y2-lD>yI@JfY3Uv#04iyiT4*>",
			"22, 0, 1, 50| 2, 1698|| {88} {2} {4} {87} {80} {14} {30} {37} {43} {51} {56} {60}|",
		},
		{
			"kickballer ok-ish",
			"@UgeU_{Fme!K@IFv#RG}7is9{5o3W<7&8i^W%x`R4}nuR)t8pZ$s",
			"10, 0, 1, 50| 2, 1976|| {100} {2} {5} {3} {1:13} {7} {64} {71} {65} {17} {27} {35} {41} {51} {1:17}|",
		},
		{
			"kickballer ok-ish",
			"@UgcJizFme!KY=H8i3bm+44I8?L>WO-Xs)u@m3WZvQN`?A`DvJO",
			"5, 0, 1, 50| 2, 22|| {60} {2} {5} {6} {1:13} {59} {78} {55} {58} {23} {32} {37} {36} {47} {82}|",
		},
		{
			"discord gun 2",
			"@UgbV{rFme!Kc0JHoRG/*Fs6;g?Eb1I89cmqFQ)5tdP=8Q;P_t0KP$Ll!V}k<",
			"3, 0, 1, 50| 2, 1883|| {81} {2} {3} {5} {4} {6} {80} {51} {52} {53} {13} {17} {26} {31} {30} {41} {47} {65}|",
		},
		{
			"discord gun 3",
			"@Ugfs(8Fme!KYJAX4)SwEr6Sb)^s88iVr9#C*-9ptutwZHR00",
			"13, 0, 1, 50| 2, 3861|| {73} {1} {2} {77} {13} {17} {15} {28} {36} {40} {43} {42} {53} {76}|",
		},
		{
			"discord gun 3 clean",
			"@UgbV{rFme!Kc0JHoRG/*Fs6;g?Eb1I89cmqFQ)5tdP=8Q;P_t0KP$Ll",
			"3, 0, 1, 50| 2, 1883|| {81} {2} {3} {5} {4} {6} {80} {51} {52} {53} {13} {17} {26} {31} {30} {41} {47} {65}|",
		},
		{
			"discord gun 3 // NO MAGAZINE",
			"@UgbV{rFme!Kc0JHoRG/*Fs6;g?Eb1I89cmqFQ*}^(P<>FdP`^+k5d",
			"3, 0, 1, 50| 2, 1883|| {81} {2} {3} {5} {4} {6} {80} {51} {52} {53} {13} {26} {31} {30} {41} {47} {65}|",
		},
		{
			"discord serial from black market",
			"@Ugydj=2}TYgT+$BRLlx>!iE7l4p){zHsFA3ds4%EKs8Xm=s9dODh=;Mc0s",
			"25, 0, 1, 50| 2, 474|| {59} {2} {3} {4} {6} {1:12} {20} {66} {65} {73} {16} {29} {34} {33} {44} {47}|",
		},
		{
			"discord serial cleared",
			"@Ugydj=2}TYgT+$BRLlx>!iE7l4p){zHsFA3ds4%EKs8Xm=s9dODh=;Mc0s",
			"25, 0, 1, 50| 2, 474|| {59} {2} {3} {4} {6} {1:12} {20} {66} {65} {73} {16} {29} {34} {33} {44} {47}|",
		},
		{
			"discord serial NO SCOPE",
			"@Ugydj=2}TYgT+$BRLlx>!iE7l4p){zHsFA3ds4%Efs8Og~s9y*",
			"25, 0, 1, 50| 2, 474|| {59} {2} {3} {4} {6} {1:12} {20} {66} {65} {73} {16} {34} {33} {44} {47}|",
		},
		{
			"cyclopean multistrike / sequencer + piercer / gun crit + gun magazine",
			"@Uge8Usm/)}}!hF1-NWCv7Xi$@SA5Ao&D+B-",
			"268, 0, 1, 50| 2, 2959|| {8} {247:76} {1} {9} {247:[97 180 4]}|",
		},
		{
			"cyclopean multistrike / sequencer + piercer / gun crit",
			"@Uge8Usm/)}}!hF1-NWCv7Xi$@S9}Qh000",
			"268, 0, 1, 50| 2, 2959|| {8} {247:76} {1} {9} {247:[180 4]}|",
		},
		{
			"cyclopean multistrike / sequencer + piercer / gun crit+ gun damage",
			"@Uge8Usm/)}}!hF1-NWCv7Xi$@SAI&bID+B-",
			"268, 0, 1, 50| 2, 2959|| {8} {247:76} {1} {9} {247:[91 180 4]}|",
		},
		{
			"discord enhancement",
			"@Uge8^+m/)}}!t6/-_/YH$",
			"300, 0, 1, 50| 2, 1179|| {248:[8]}|",
		},
		{
			"discord gun + hyperion grip",
			"@UgbV{rFme!K?j_JzRG}7is6;g?ENUGp9qJrvQ+H6OP^D10P$>}",
			"3, 0, 1, 50| 2, 1143|| {81} {2} {5} {3} {4} {6} {80} {53} {52} {51} {13} {27} {35} {34} {44} {68}|",
		},
		{
			"discord gun + no grip",
			"@UgbV{rFme!K?j_JzRG}7is6;g?ENUGp9qJrvQ+H6OP^C~Q5d",
			"3, 0, 1, 50| 2, 1143|| {81} {2} {5} {3} {4} {6} {80} {53} {52} {51} {13} {27} {35} {34} {68}|",
		},
		{
			"discord destination with unwanted grip",
			"@UgbV{rFme!Kc0JHoRG/*Fs6;g?Eb1I89cmqFQ)5tdP=8Q;P_t0KP$Ll",
			"3, 0, 1, 50| 2, 1883|| {81} {2} {3} {5} {4} {6} {80} {51} {52} {53} {13} {17} {26} {31} {30} {41} {47} {65}|",
		},
		{
			"discord destination with no grip",
			"@UgbV{rFme!Kc0JHoRG/*Fs6;g?Eb1I89cmqFQ)5tdP=8Q;P`^+k5d",
			"3, 0, 1, 50| 2, 1883|| {81} {2} {3} {5} {4} {6} {80} {51} {52} {53} {13} {17} {26} {31} {30} {47} {65}|",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			blocks, bits, err := Deserialize(data)
			assert.NoError(t, err)
			fmt.Println("Result:", blocks.String())
			fmt.Println("Bits:", bits)
		})
	}
}

func TestDeserializePhospheneSkin(t *testing.T) {
	var tests = []struct {
		serialNormal    string
		serialPhosphene string
	}{
		{
			// 111 10000 110001101 100 00110 00000(eof)
			"@UgdhV<Fme!Kq%_bvRG/))sC1~Bs7#GP%/V4i%/iV`?L+_",
			"@UgdhV<Fme!Kq%_bvRG/))sC1~Bs7#GP%/V4i%/iV`?L<6`4Fd",
		},
		{
			// 111 10000 110001101 100 00110 00(eof)
			"@UgdhV<Fme!Kye7~(RG/*FsC1~Bs7ie*C#nwW4{8=_9I7V*",
			"@UgdhV<Fme!Kye7~(RG/*FsC1~Bs7ie*C#nwW4{8=_9I7YcVQd%",
		},
		{
			// 111 10000 110001101 100 0010110000 0000000(eof)
			"@Ugd77*Fme!KY)?>XRG}8ts6-7H3W<t/YKfYO3WLgnI)zGw%7rS200",
			"@Ugd77*Fme!KY)?>XRG}8ts6-7H3W<t/YKfYO3WLgnI)zGw%7rS2co-WR00",
		},
		{
			// 111 10000 110001101 100 1001101000 00000000(eof)
			"@UgggUGFme!K_(jl7R87>O3Kb7^4^<D<sX>)OjY7pjwL/~",
			"@UgggUGFme!K_(jl7R87>O3Kb7^4^<D<sX>)OjY7pjwM0COO%?z",
		},
		{
			// 111 10000 110001101 100 0001101000 00(eof)
			"@Ugydj=2}TYgT#^p$Llx>!jY`yzp?s*9sF/pisFkQq?Ln17jY920r9%zj00",
			"@Ugydj=2}TYgT#^p$Llx>!jY`yzp?s*9sF/pisFkQq?Ln17jY920r9%zj01sn>1p",
		},
		{
			// 111 10000 110001101 100 0111110000 000(eof)
			"@Ugy3L+2}TYgOx*?=RG/)46V(s35A~@ss5z)xsB#D",
			"@Ugy3L+2}TYgOx*?=RG/)46V(s35A~@ss5z)xsB(yhvGD-",
		},
		{
			// 111 10000 110001101 100 1001101000 00000(eof)
			"@UgggUGFme!KB5%-5R87>O8da!ysCcM)sC%eR4XPCC6{;5k",
			"@UgggUGFme!KB5%-5R87>O8da!ysCcM)sC%eR4XPCC6{;8FVQjJh",
		},
		{
			// 111 10000 110001101 100 00110 000000000(eof)
			"@UgdhV<Fme##(tBtfs!)kahpLIX)ELwqR4LRcR4r6ER8Ir",
			"@UgdhV<Fme##(tBtfs!)kahpLIX)ELwqR4LRcR4r6ER8Pdi*f0P",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			{
				data, err := b85.Decode(tt.serialNormal)
				assert.NoError(t, err)

				blocks, bits, err := Deserialize(data)
				assert.NoError(t, err)
				fmt.Println("Serial:", tt.serialNormal)
				fmt.Println("Result:", blocks.String())
				fmt.Println("Bits:", bits)
				fmt.Println()
			}

			{
				data, err := b85.Decode(tt.serialPhosphene)
				assert.NoError(t, err)

				blocks, bits, err := Deserialize(data)
				assert.NoError(t, err)
				fmt.Println("Serial:", tt.serialPhosphene)
				fmt.Println("Result:", blocks.String())
				fmt.Println("Bits:", bits)
				fmt.Println()
			}

		})
	}
}

func TestDeserializeCompareBuybacks(t *testing.T) {
	var tests = []struct {
		name           string
		serialOriginal string
		serialBuyback  string
		parsedOriginal string
		parsedBuyback  string
	}{
		{
			"L50 Legendary Cooking Ambushing Truck",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			"@Ugy3L+2}Ta0Od!I{*`S=LLLKTRY91;d>K-Z#Y7QzFY8(O",
			"24, 0, 1, 50| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16} {25} {44} {49}|",
			"24, 0, 1, 50| 10, 1| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16} {25} {44} {49}|",
		},
		{
			"L50 Legendary Ambushing Truck",
			"@Ugy3L+2}TYgjMogxi7Hg07IhPq4>b?9sX3@zs9y*",
			"@Ugy3L+2}Ta0Od!H/&7hp9LM3WZH&OXe^H7_bgUW^ag#Z",
			"24, 0, 1, 50| 2, 3246|| {76} {2} {4} {5} {75} {60} {57} {13} {25} {44} {47}|",
			"24, 0, 1, 50| 10, 1| 2, 3246|| {76} {2} {4} {5} {75} {60} {57} {13} {25} {44} {47}|",
		},
		{
			"L49 Uncommon Playful Kitty",
			"@Ugct)%FmVuJXn{hb3U#POJ!&6nQ*lsxP_0lm5d",
			"@Ugct)%FmVuN0uhE5C^V{2hg#I5_MtWv2ek*)3Uw0!",
			"6, 0, 1, 49| 2, 84|| {96} {2} {3} {5} {7} {61} {13} {24} {29} {38} {75}|",
			"6, 0, 1, 49| 10, 1| 2, 84|| {96} {2} {3} {5} {7} {61} {13} {24} {29} {38} {75}|",
		},
		{
			"L49 Common Karkadann",
			"@UgzR8/2__CAOuq;Eiz?Kj9yO^ss8y(TsD20",
			"@UgzR8/2__DrOd!Jad!WClLM`f1lbVBCg=&ZDhX4",
			"27, 0, 1, 49| 2, 1917|| {95} {2} {5} {7} {9} {25} {37} {54} {62}|",
			"27, 0, 1, 49| 10, 1| 2, 1917|| {95} {2} {5} {7} {9} {25} {37} {54} {62}|",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataOriginal, err := b85.Decode(tt.serialOriginal)
			assert.NoError(t, err)

			start := time.Now()
			parsedOriginal, _, err := Deserialize(dataOriginal)
			end := time.Now()
			fmt.Printf("Deserialization took %v\n", end.Sub(start))

			assert.NoError(t, err)
			assert.Equal(t, tt.parsedOriginal, parsedOriginal.String())

			dataBuyback, err := b85.Decode(tt.serialBuyback)
			assert.NoError(t, err)
			parsedBuyback, _, err := Deserialize(dataBuyback)
			assert.NoError(t, err)
			assert.Equal(t, tt.parsedBuyback, parsedBuyback.String())

			fmt.Println("Original:", parsedOriginal.String())
			fmt.Println("Buyback: ", parsedBuyback.String())

			// Remove the section added by the buyback, and compare
			buybakWithoutBuyback := strings.ReplaceAll(parsedBuyback.String(), "| 10, 1|", "|")
			assert.Equal(t, parsedOriginal.String(), buybakWithoutBuyback)
		})
	}
}
