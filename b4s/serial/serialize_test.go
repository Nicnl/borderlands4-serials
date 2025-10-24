package serial

import (
	"borderlands_4_serials/b4s/b85"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializeFromString(t *testing.T) {
	var tests = []struct {
		name         string
		expected     string
		deserialized string
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
		{
			"test debug 2",
			"@UgdhV<Fme!Kq%_bvRG/))sC1~Bs7#GP%/V4i%/iV`?L+_",
			"8, 0, 1, 50| 2, 1570|| {53} {2} {3} {4} {52} {74} {12} {17} {25} {32} {41} {47} {77}|",
		},
		{
			"test debug 3",
			"@UgdhV<Fme##(tBtfs!)kahpLIX)ELwqR4LRcR4r6ER8Ir",
			"8, 0, 1, 50| 2, 495|| {53} {2} {4} {52} {74} {11} {17} {27} {34} {35} {42} {48} {78}|",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Deserialized:", tt.deserialized)

			s := Serial{}

			err := s.FromString(tt.deserialized)
			assert.NoError(t, err)

			serializedData := Serialize(s)
			fmt.Println("Serialized data:", fmt.Sprintf("%x", serializedData))

			serializedB85 := b85.Encode(serializedData)
			fmt.Println("Serialized B85:", serializedB85)
			fmt.Println("Expected B85:  ", tt.expected)

			assert.Equal(t, tt.expected, serializedB85)
		})
	}
}

func TestSerializePartsRoundtrip(t *testing.T) {
	var tests = []struct {
		name string
		b85  string
	}{
		{
			"ORIGINAL L50 Legendary Cooking Ambushing Truck",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
		},
		{
			"Knife 1 skill",
			"@Ugr$WBm/$!m!X=5&qXq#",
		},
		{
			"Knife 2 skill",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3Nj00",
		},
		{
			"Knife 3 skill",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3OODgg",
		},
		{
			"Knife 4 skill",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3OOD#<4R",
		},
		{
			"Common Unseen Xiuhcoatl",
			"@Ugy3L+2}TYgOyvyviz?KiBDJYGs9dOW2m",
		},
		{
			"Green Unseen Xiuhcoatl",
			"@Ugy3L+2}TMcjNb(cjVjck8WpL1s7>WTg+kRrl/uj",
		},
		{
			"Purple Looming Xiuhcoatl",
			"@Ugy3L+2}TYg4BQJUjVjck61AvE^+Sb3b!rZ(7U~=V",
		},
		{
			"Top Square Simple 1",
			"@Ugy3L+2}TYgOyvyviz?KiBDJYGs9dOW2m",
		},
		{
			"Top Square Simple 2",
			"@Ugy3L+2}TYgjMogxi7Hg07IhPq4>b?9sX3@zs9y*",
		},
		{
			"Top Square Simple 3",
			"@Ugy3L+2}TYg4BQJUjVjck61AvE^+Sb3b!rZ(7U~=V",
		},
		{
			"Top Square Simple 4",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
		},
		{
			"Top Square Simple 5",
			"@Ugy3L+2}TYgT#^cvMir`2hg#I5@}cgb=Ak+@2XzZ/4gm",
		},
		{
			"Side Long Smooth  1",
			"@Ugy3L+2}TYgOyvyviz?KiBDJYKs9dOW2m",
		},
		{
			"Side Long Smooth  2",
			"@Ugy3L+2}TYgjMogxi7Hg07IhPq4>b?9sXeG%s9y*",
		},
		{
			"Side Long Smooth  3",
			"@Ugy3L+2}TYg4BQJUjVjck61AvE^+Sb3b!rc)7U~=V",
		},
		{
			"Side Long Smooth  4",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l3q`a-qf{00",
		},
		{
			"Side Long Smooth  5",
			"@Ugy3L+2}TYgT#^cvMir`2hg#I5@}cgb=Ak-u2XzZ/4gm",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			data, err := b85.Decode(tt.b85)
			assert.NoError(t, err)

			serial, _, err := Deserialize(data)
			assert.NoError(t, err)
			fmt.Println("Parsed:", serial.String())

			serializedData := Serialize(serial)
			assert.NoError(t, err)

			reserializedB85 := b85.Encode(serializedData)
			assert.Equal(t, tt.b85, reserializedB85)
			fmt.Println("Original:    ", tt.b85)
			fmt.Println("Reserialized:", reserializedB85)
		})
	}
}

func TestSerializeSkins(t *testing.T) {
	for _, tt := range skinTests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
