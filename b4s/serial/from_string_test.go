package serial

import (
	"borderlands_4_serials/b4s/b85"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializePartsStringRoundtrip(t *testing.T) {
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
		{
			"with simple int subpart",
			"@Ug!pHG38o5YT`HzQ)h-nP",
		},
		{
			"with list subpart",
			"@Uge8^+m/)}}!c178NkyuCbwKf>IWYh",
		},
		{
			"1 arms    1 bottle",
			"@Ug!pHG38o5YT`HzQ)h-nP",
		},
		{
			"1 cat     1 bottle",
			"@Ug!pHG38o5YZ7QZg)h-nP",
		},
		{
			"1 skullR  1 bottle",
			"@Ug!pHG38o5YOe&^9)h-nP",
		},
		{
			"1 bullet  1 bottle",
			"@Ug!pHG38o6@O)92A)h-nP",
		},
		{
			"1 square  1 bottle",
			"@Ug!pHG38o5YPb#KC)h-nP",
		},
		{
			"1 skullG  1 bottle",
			"@Ug!pHG38o5YMJlF2)h-nP",
		},
		{
			"1 feet    1 bottle",
			"@Ug!pHG38o4tO)92A)h-nP",
		},
		{
			"1 empty   1 bottle",
			"@Ug!pHG38o5Y4JxKV)h-nP",
		},
		{
			"1 empty   1 bottle",
			"@Ug!pHG38o5Y4JxKV)h-nP",
		},
		{
			"1 Rarms   1 Bbottle  + 1Gboom",
			"@Ug!pHG38o5YT`HzQ)k4)S6#x",
		},
		{
			"1 Rarms ",
			"@Ug!pHG38o5YU8;7e00",
		},
		{
			"1 Rarms 2 bottles",
			"@Ug!pHG38o5YT`HzQ#Wbker2+r",
		},
		{
			"2 Rarms 1 bottles",
			"@Ug!pHG38o5YU20t_ra{#%6#x",
		},
		{
			"melee (red skin)",
			"@Ug!pHG38o6DcBud",
		},
		{
			"melee (white skin)",
			"@Ug!pHG38o6DP_;`100",
		},
		{
			"1 arm no melee",
			"@Ug!pHG38o5YT>=",
		},
		{
			"1 arm 1 bottle + firmware jacked",
			"@Ug!pHG38o5YT`HzQ)$V@)",
		},
		{
			"unseen",
			"@Ugy3L+2}TYgOyvyviz?KiBDJYGs9dOW2m",
		},
		{
			"shield 1",
			"@Uge8^+m/)}}!c178NkyuCbwKf>IWYh",
		},
		{
			"shield 2",
			"@Uge8^+m/)}}!axR1DpKvM1BxF_41oav",
		},
		{
			"rapid swarm + gadget ahoy",
			"@Uge8aum/)}}!qkqSNDXRzG&iINder)8E{Op",
		},
		{
			"waterfall grenade + airstrike",
			"@Ugr$)Nm/)}}!YpV~ky;-O59uLV#F7vI",
		},
		{
			"occulted ephemeris",
			"@Ugr$!Lm/)}}!u<K5M>VQ_G&h6`+T9-j",
		},
		{
			"looming balor",
			"@Ugd_t@Fme!KdTvl?RG/_Tse7ors5+=wsFVl",
		},
		{
			"double creme omnibore",
			"@Uge(J0Fme!Kux-$2RG}7is6<7oB&t$xP@zz<P`yy=5C",
		},
		{
			"shield 1",
			"@Uge8^+m/)}}!c178NkyuCbwKf>IWYh",
		},
		{
			"shield 2",
			"@Uge8^+m/)}}!axR1DpKvM1BxF_41oav",
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
			"ORIGINAL L50 Legendary Cooking Ambushing Truck",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
		},
		{
			"no level spawned console",
			"@Ugr$lG7-8sL(4z`<KALPY4GrpidjS",
		},
		{
			"Vindictive Evolver",
			"@Ugr$rIm/)}}!q`oqNWCv7s8Ex7AI%h@D>DE",
		},
		{
			"Retributive Devourer",
			"@Uge8;)m/)}}!sxA_MZGU4Xi$ZEAI&bYFAo3",
		},
		{
			"Skin: Solar Flair",
			"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vGE1",
		},
		{
			"Skin: Carcade Shooter",
			"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vB?G",
		},
		{
			"Skin: Itty Bitty Kitty Committee",
			"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vB3r",
		},
		{
			"Skin: With the grain",
			"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vGD}",
		},
		{
			"Skin: The System",
			"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vAG2",
		},
		{
			"Skin: Devourer",
			"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_v9Sd",
		},
		{
			"Skin: Soused",
			"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_v5^G",
		},
		{
			"Skin: Bird of Prey",
			"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_v55r",
		},
		{
			"Skin: Eternal Defender",
			"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vH1i",
		},
		{
			"Skin: Weirdo",
			"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vGD`",
		},
		{
			"Skin: Smiley",
			"@UgbV{rFme!KI4sa#RG}W#sX3@xsFsL_vFQW",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fmt.Println("Original:    ", tt.b85)

			var parts string
			t.Run("Deserialize", func(t *testing.T) {
				data, err := b85.Decode(tt.b85)
				assert.NoError(t, err)

				serial, _, err := Deserialize(data)
				assert.NoError(t, err)

				parts = serial.String()
				fmt.Println("Parsed:", parts)
			})

			var imported Serial
			t.Run("Import", func(t *testing.T) {
				err := imported.FromString(parts)
				assert.NoError(t, err)
			})

			t.Run("Reserialize", func(t *testing.T) {
				serializedData := Serialize(imported)
				reserializedB85 := b85.Encode(serializedData)
				assert.Equal(t, tt.b85, reserializedB85)
				fmt.Println("Reserialized:", reserializedB85)
			})
		})
	}
}

func TestImportStringCrash(t *testing.T) {
	// Found by a user:
	// 291, 0, 1, 50| 9, 1| 10, 1| 2, 1375|| {5} {8} {245:}|
	// It should generate an error (unfinished part at the end)
	// But it panics

	var s Serial
	err := s.FromString("291, 0, 1, 50| 9, 1| 10, 1| 2, 1375|| {5} {8} {245:}|")
	assert.Error(t, err)
	fmt.Println("Error:", err.Error())
}

func TestImportStringCrash2(t *testing.T) {
	// Found by a user:
	// 291, 0, 1, 50| 9, 1| 10, 1| 2, 1375|| {5} {8} {245:}|
	// It should generate an error (unfinished part at the end)
	// But it panics

	var s Serial
	err := s.FromString("291, 0, 1, 50| 9, 1| 10, 1| 2, 1375|| {5} {8} {245:[}|")
	assert.Error(t, err)
	fmt.Println("Error:", err.Error())
}

func TestImportStringSpecialChars(t *testing.T) {
	baseDeserialized := `"my name is \"The Boss\" and I use \\ in paths"`

	var s Serial
	err := s.FromString(baseDeserialized)
	assert.NoError(t, err)
	fmt.Println("Parts:", s.String())
	assert.Equal(t, `my name is "The Boss" and I use \ in paths`, s[0].ValueStr)

	// Reserialize
	serializedData := Serialize(s)
	reserializedB85 := b85.Encode(serializedData)
	assert.Equal(t, "@Uglo~xgWTbE8I+!bL{xMcBz({3B2d^(1/>oDc^Sk7rQINSn2w$U", reserializedB85)

	// Deserialize again to check
	data, err := b85.Decode(reserializedB85)
	assert.NoError(t, err)

	serial, _, err := Deserialize(data)
	assert.NoError(t, err)
	assert.Equal(t, `my name is "The Boss" and I use \ in paths`, serial[0].ValueStr)
}
