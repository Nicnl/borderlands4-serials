package serial

import (
	"borderlands_4_serials/b4s/b85"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeserializeItemSet1(t *testing.T) {
	var tests = []struct {
		name     string
		serial   string
		expected string
	}{
		{
			"ORIGINAL L50 Legendary Cooking Ambushing Truck",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			"24, 0, 1, 50| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16} {25} {44} {49}|",
		},
		{
			"ORIGINAL L50 Legendary Cooking Ambushing Truck SMALL",
			"@Ugy3L+35F42=4?<-RG/)a6EzQ&4/NX}1~mtj3pEY_",
			"24, 0, 1, 12| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16} {25} {44} {49}|",
		},
		{
			"BROKEN TRUCK 1",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!XN+",
			"24, 0, 1, 50| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16}|",
		},
		{
			"BROKEN TRUCK 3",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!XN+",
			"24, 0, 1, 50| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16}|",
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
			"Common Unseen Xiuhcoatl",
			"@Ugy3L+2}TYgOyvyviz?KiBDJYGs9dOW2m",
			"24, 0, 1, 50| 2, 3269|| {95} {2} {3} {8} {13} {25} {44} {50}|",
		},
		{
			"Green Unseen Xiuhcoatl",
			"@Ugy3L+2}TMcjNb(cjVjck8WpL1s7>WTg+kRrl/uj",
			"24, 0, 1, 34| 2, 1470|| {96} {2} {3} {6} {8} {61} {13} {28} {32} {42} {50}|",
		},
		{
			"Purple Looming Xiuhcoatl",
			"@Ugy3L+2}TYg4BQJUjVjck61AvE^+Sb3b!rZ(7U~=V",
			"24, 0, 1, 50| 2, 2992|| {98} {2} {3} {4} {5} {8} {62} {64} {14} {25} {42} {51}|",
		},
		{
			"Top Square Simple 1",
			"@Ugy3L+2}TYgOyvyviz?KiBDJYGs9dOW2m",
			"24, 0, 1, 50| 2, 3269|| {95} {2} {3} {8} {13} {25} {44} {50}|",
		},
		{
			"Top Square Simple 2",
			"@Ugy3L+2}TYgjMogxi7Hg07IhPq4>b?9sX3@zs9y*",
			"24, 0, 1, 50| 2, 3246|| {76} {2} {4} {5} {75} {60} {57} {13} {25} {44} {47}|",
		},
		{
			"Top Square Simple 3",
			"@Ugy3L+2}TYg4BQJUjVjck61AvE^+Sb3b!rZ(7U~=V",
			"24, 0, 1, 50| 2, 2992|| {98} {2} {3} {4} {5} {8} {62} {64} {14} {25} {42} {51}|",
		},
		{
			"Top Square Simple 4",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			"24, 0, 1, 50| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16} {25} {44} {49}|",
		},
		{
			"Top Square Simple 5",
			"@Ugy3L+2}TYgT#^cvMir`2hg#I5@}cgb=Ak+@2XzZ/4gm",
			"24, 0, 1, 50| 2, 338|| {98} {2} {6} {3} {5} {7} {60} {59} {57} {14} {25} {43} {51}|",
		},
		{
			"Side Long Smooth  1",
			"@Ugy3L+2}TYgOyvyviz?KiBDJYKs9dOW2m",
			"24, 0, 1, 50| 2, 3269|| {95} {2} {3} {8} {13} {29} {44} {50}|",
		},
		{
			"Side Long Smooth  2",
			"@Ugy3L+2}TYgjMogxi7Hg07IhPq4>b?9sXeG%s9y*",
			"24, 0, 1, 50| 2, 3246|| {76} {2} {4} {5} {75} {60} {57} {13} {29} {44} {47}|",
		},
		{
			"Side Long Smooth  3",
			"@Ugy3L+2}TYg4BQJUjVjck61AvE^+Sb3b!rc)7U~=V",
			"24, 0, 1, 50| 2, 2992|| {98} {2} {3} {4} {5} {8} {62} {64} {14} {29} {42} {51}|",
		},
		{
			"Side Long Smooth  4",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l3q`a-qf{00",
			"24, 0, 1, 50| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16} {29} {44} {49}|",
		},
		{
			"Side Long Smooth  5",
			"@Ugy3L+2}TYgT#^cvMir`2hg#I5@}cgb=Ak-u2XzZ/4gm",
			"24, 0, 1, 50| 2, 338|| {98} {2} {6} {3} {5} {7} {60} {59} {57} {14} {27} {43} {51}|",
		},
		{
			"test debug",
			"@Ugy3L+2}TYg%$yC%i7Es",
			"24, 0, 1, 50| 2, 3379|| {76} {2}|",
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
			parsedOriginal, _, err := Deserialize(dataOriginal)
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
