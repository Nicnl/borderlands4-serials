package serial

import (
	"borderlands_4_serials/b4s/b85"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			"ORIGINAL L50 Legendary Cooking Ambushing Truck SMALL",
			"@Ugy3L+35F42=4?<-RG/)a6EzQ&4/NX}1~mtj3pEY_",
			"1",
		},
		{
			"BROKEN TRUCK 1",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!XN+",
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

			parsed, err := Deserialize(data)
			assert.NoError(t, err)
			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			fmt.Println("Result:", parsed.String())
			fmt.Println("Bitstream:", parsed.Bits)
			fmt.Println()
			fmt.Println()
		})
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
		{
			"1 empty   1 bottle",
			"@Ug!pHG38o5Y4JxKV)h-nP",
			"1",
		},
		{
			"1 Rarms   1 Bbottle  + 1Gboom",
			"@Ug!pHG38o5YT`HzQ)k4)S6#x",
			"1",
		},
		{
			"1 Rarms ",
			"@Ug!pHG38o5YU8;7e00",
			"1",
		},
		{
			"1 Rarms 2 bottles",
			"@Ug!pHG38o5YT`HzQ#Wbker2+r",
			"1",
		},
		{
			"2 Rarms 1 bottles",
			"@Ug!pHG38o5YU20t_ra{#%6#x",
			"1",
		},
		{
			"melee (red skin)",
			"@Ug!pHG38o6DcBud",
			"1",
		},
		{
			"melee (white skin)",
			"@Ug!pHG38o6DP_;`100",
			"1",
		},
		{
			"1 arm no melee",
			"@Ug!pHG38o5YT>=",
			"1",
		},
		{
			"1 arm 1 bottle + firmware jacked",
			"@Ug!pHG38o5YT`HzQ)$V@)",
			"1",
		},
		{
			"unseen",
			"@Ugy3L+2}TYgOyvyviz?KiBDJYGs9dOW2m",
			"1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			parsed, err := Deserialize(data)
			assert.NoError(t, err)
			fmt.Println("Result:", parsed.String())
			fmt.Println("Bits:", parsed.Bits)
		})
	}
}

func TestSerialTokenizeFirmware(t *testing.T) {
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
		{
			"rapid swarm + gadget ahoy",
			"@Uge8aum/)}}!qkqSNDXRzG&iINder)8E{Op",
			"1",
		},
		{
			"waterfall grenade + airstrike",
			"@Ugr$)Nm/)}}!YpV~ky;-O59uLV#F7vI",
			"1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			parsed, err := Deserialize(data)
			assert.NoError(t, err)
			fmt.Println("Result:", parsed.String())
			fmt.Println("Bits:", parsed.Bits)
		})
	}
}

func TestSerialRandom(t *testing.T) {
	var tests = []struct {
		name     string
		serial   string
		expected string
	}{
		{
			"occulted ephemeris",
			"@Ugr$!Lm/)}}!u<K5M>VQ_G&h6`+T9-j",
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
		{
			"no level spawned console",
			"@Ugr$lG7-8sL(4z`<KALPY4GrpidjS",
			"1",
		},
		{
			"Vindictive Evolver",
			"@Ugr$rIm/)}}!q`oqNWCv7s8Ex7AI%h@D>DE",
			"1",
		},
		{
			"Retributive Devourer",
			"@Uge8;)m/)}}!sxA_MZGU4Xi$ZEAI&bYFAo3",
			"1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			parsed, err := Deserialize(data)
			assert.NoError(t, err)
			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			fmt.Println("Result:", parsed.String())
			fmt.Println("Bits:", parsed.Bits)
		})
	}
}

func TestSerialProblematicSerials2(t *testing.T) {
	var tests = []struct {
		name     string
		serial   string
		expected string
	}{
		{
			"(Overshielding Junker Principal) => invalid token, got 111 at position 177",
			"@Uge8^+m/)}}!hAcRNkyuCbwKeuQ2X8i00",
			"1",
		},
		{
			"(Watts 4 Dinner) => invalid token, got 111 at position 149",
			"@Uge8^+m/)}}!X>8AqZ-w/IH35^oFHmn1po",
			"1",
		},
		{
			"(Resonant Itchy Sparky Shield) => invalid token, got 111 at position 148",
			"@Uge92<m/)}}!hGMLL{+MNG&aNz+T9-j",
			"1",
		},
		{
			"(High-Capacity Junker Cask) => invalid token, got 111 at position 177",
			"@Uge8^+m/)}}!rTKTs#5Kn1B%~++II&4",
			"1",
		},
		{
			"(Resonant Dextrous Scar) => invalid token, got 111 at position 144",
			"@Ugr$oHm/)}}!gO1pMir`kaX/6gQ2WLK00",
			"1",
		},
		{
			"(Fleeting Vigorous Stanchion) => invalid token, got 111 at position 149",
			"@Ugr$oHm/)}}!t7k2N)@VobwKf(&>#Q",
			"1",
		},
		{
			"(Trigger-Happy Extra Medium) => invalid token, got 111 at position 149",
			"@Uge9B?m/)}}!t?`8DpKvM1BxF_3k?7",
			"1",
		},
		{
			"(Fleeting Anxious Super Soldier) => invalid token, got 111 at position 144",
			"@Ugr$uJm/)}}!q`oqM>VQ_Z$P^nQ2XWq00",
			"1",
		},
		{
			"(Scrapper Emollient Hoarder) => invalid token, got 111 at position 182",
			"@Ugr$uJm/)}}!bF&$M>VQ_Z9uylQ2XWq00",
			"1",
		},
		{
			"(Scrapper Emollient Hoarder) => invalid token, got 111 at position 182",
			"@Ugr$uJm/)}}!tBVPL{+MNb3nTr18UzM00",
			"1",
		},
		{
			"(Bladed Chitinous Domy) => invalid token, got 111 at position 185",
			"@Uge9B?m/)}}!f5KCL{+MNaX/5-sUb7~",
			"1",
		},
		{
			"(Vigilant Bunker) => invalid token, got 111 at position 149",
			"@Ugr%Scm/)}}!a#_iM^&nQYe2i3Q2X/P8UX",
			"1",
		},
		{
			"(Barrage Chitinous Firewerks) => invalid token, got 111 at position 180",
			"@Uge8^+m/)}}!gz?FNkyuCbwKf>nIVqa7X/<",
			"1",
		},
		{
			"(Watts 4 Dinner) => invalid token, got 111 at position 149",
			"@Ugr$-Om/)}}!g!mYLlvrhbwKgEQ2Ry!00",
			"1",
		},
		{
			"(Nova Pointed Protean Cell) => invalid token, got 111 at position 149",
			"@Ugr$-Om/)}}!g%VSN+qg&aX/5-DcOMr0R",
			"1",
		},
		{
			"(Berserkr Meandering Cindershelly) => invalid token, got 111 at position 149",
			"@Uge8^+m/)}}!dO_KNkyuCbwKf>sUeKoR{#",
			"1",
		},
		{
			"(Watts 4 Dinner) => invalid token, got 111 at position 149",
			"@Ugr$uJm/)}}!qj@8M>VQ_Ye2gjQ2XWq00",
			"1",
		},
		{
			"(Scrapper Emollient Hoarder) => invalid token, got 111 at position 182",
			"@Ugr$oHm/)}}!uW`wNfoMnbwKgEQ2XA2A^`",
			"1",
		},
		{
			"(Nova Mutable Pandoran Memento) => invalid token, got 111 at position 149",
			"@Ugr$oHm/)}}!r*P7N)@VoaX/5V&>#Q",
			"1",
		},
		{
			"(Resonant Extra Medium) => invalid token, got 111 at position 149",
			"@Ugr$oHm/)}}!pO~_N)@Vob3pN1Q2Ww=1_1",
			"1",
		},
		{
			"(Berserkr Extra Medium) => invalid token, got 111 at position 149",
			"@Uge9B?m/)}}!qjb`L{+MNb3pOiQ2XWq00",
			"1",
		},
		{
			"(Fleeting Emollient Bunker) => invalid token, got 111 at position 149",
			"@Uge8^+m/)}}!uYwMM>VQ_Z9wrGQ2TxX00",
			"1",
		},
		{
			"(Traveler Vitamin Sparky Shield) => invalid token, got 111 at position 180",
			"@Ugr$uJm/)}}!uW5XM>VQ_Z$P^nQ2XWq00",
			"1",
		},
		{
			"(Scrapper Emollient Hoarder) => invalid token, got 111 at position 182",
			"@Ugr$oHm/)}}!pu&fNfoMnb3pN<i6Mp}0R",
			"1",
		},
		{
			"(Amp Tolling Pandoran Memento) => invalid token, got 111 at position 149",
			"@Ugr$uJm/)}}!f216M>VQ_bwIltQ2XWq00",
			"1",
		},
		{
			"(Scrapper Emollient Hoarder) => invalid token, got 111 at position 182",
			"@Ugr$-Om/)}}!pz#BLlvrhbwKgEQ2Ry!00",
			"1",
		},
		{
			"(Nova Pointed Protean Cell) => invalid token, got 111 at position 149",
			"@Uge9B?m/)}}!X>CsqAJzCIiUCrsC}mZ",
			"1",
		},
		{
			"(Traveler Weatherproof Bunker) => invalid token, got 111 at position 148",
			"@Ugr$!Lm/)}}!cfVeL{+MNYe2i318UzM00",
			"1",
		},
		{
			"(Boxer Chitinous Laminate) => invalid token, got 111 at position 185",
			"@Ugr$uJm/)}}!ljeYqZ-w/H=x}OsD1MQ",
			"1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			parsed, err := Deserialize(data)
			assert.NoError(t, err)
			fmt.Println("Result:", parsed.String())
			fmt.Println("Bits:", parsed.Bits)
		})
	}
}

func TestSerialProblematicSerials3(t *testing.T) {
	var tests = []struct {
		name     string
		serial   string
		expected string
	}{
		{
			"Sho Kunai (gadget)",
			"@Ugr$WBm/)}}!bEtWObu#%Z$Os-",
			"1",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!bmKjM-}RPG}*)&8r1p10ss",
			"1",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!t}17M-}RPG}&Yt8r1pj0ss",
			"1",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!rWS*M-}RPG}(k28r1pH0ss",
			"1",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!gSxDM-}RPG}&Yt8r1p10ss",
			"1",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!f+~}M-}RPG}*)&8r1n{b_oI",
			"1",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!u(vIM-}RPG}&Yt8r1n{aLED",
			"1",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!bIGlM-}RPG}&Yt8r1n{Y{?7&",
			"1",
		},
		{
			"Kill Sprint Repkit",
			"@Uge8#%m/)}}!qBXsM-}RPG}(k28r1n{WC;Q",
			"1",
		},
		{
			"Ballistic Bonn-91 (weapon)",
			"@Uga`vnFme!Kq<v6nRG}8tsG&oVnu!XD8i/_J7}Pw}9aJgQDby^~IMg`=\"",
			"1",
		},
		{
			"discord gun",
			"@UgxFw!2}TYgjMN5-iz-y2-lD>yI@JfY3Uv#04iyiT4*>",
			"1",
		},
		{
			"kickballer ok-ish",
			"@UgeU_{Fme!K@IFv#RG}7is9{5o3W<7&8i^W%x`R4}nuR)t8pZ$s",
			"1",
		},
		{
			"kickballer ok-ish",
			"@UgcJizFme!KY=H8i3bm+44I8?L>WO-Xs)u@m3WZvQN`?A`DvJO",
			"1",
		},
		{
			"discord gun 2",
			"@UgbV{rFme!Kc0JHoRG/*Fs6;g?Eb1I89cmqFQ)5tdP=8Q;P_t0KP$Ll!V}k<",
			"1",
		},
		{
			"discord gun 3",
			"@Ugfs(8Fme!KYJAX4)SwEr6Sb)^s88iVr9#C*-9ptutwZHR00",
			"1",
		},
		{
			"discord gun 3 clean",
			"@UgbV{rFme!Kc0JHoRG/*Fs6;g?Eb1I89cmqFQ)5tdP=8Q;P_t0KP$Ll",
			"1",
		},
		{
			"discord gun 3 // NO MAGAZINE",
			"@UgbV{rFme!Kc0JHoRG/*Fs6;g?Eb1I89cmqFQ*}^(P<>FdP`^+k5d",
			"1",
		},
		{
			"discord serial from black market",
			"@Ugydj=2}TYgT+$BRLlx>!iE7l4p){zHsFA3ds4%EKs8Xm=s9dODh=;Mc0s",
			"1",
		},
		{
			"discord serial cleared",
			"@Ugydj=2}TYgT+$BRLlx>!iE7l4p){zHsFA3ds4%EKs8Xm=s9dODh=;Mc0s",
			"1",
		},
		{
			"discord serial NO SCOPE",
			"@Ugydj=2}TYgT+$BRLlx>!iE7l4p){zHsFA3ds4%Efs8Og~s9y*",
			"1",
		},
		{
			"cyclopean multistrike / sequencer + piercer / gun crit + gun magazine",
			"@Uge8Usm/)}}!hF1-NWCv7Xi$@SA5Ao&D+B-",
			"1",
		},
		{
			"cyclopean multistrike / sequencer + piercer / gun crit",
			"@Uge8Usm/)}}!hF1-NWCv7Xi$@S9}Qh000",
			"1",
		},
		{
			"cyclopean multistrike / sequencer + piercer / gun crit+ gun damage",
			"@Uge8Usm/)}}!hF1-NWCv7Xi$@SAI&bID+B-",
			"1",
		},
		{
			"discord enhancement",
			"@Uge8^+m/)}}!t6/-_/YH$",
			"1",
		},
		{
			"discord gun + hyperion grip",
			"@UgbV{rFme!K?j_JzRG}7is6;g?ENUGp9qJrvQ+H6OP^D10P$>}",
			"1",
		},
		{
			"discord gun + no grip",
			"@UgbV{rFme!K?j_JzRG}7is6;g?ENUGp9qJrvQ+H6OP^C~Q5d",
			"1",
		},
		{
			"discord destination with unwanted grip",
			"@UgbV{rFme!Kc0JHoRG/*Fs6;g?Eb1I89cmqFQ)5tdP=8Q;P_t0KP$Ll",
			"1",
		},
		{
			"discord destination with no grip",
			"@UgbV{rFme!Kc0JHoRG/*Fs6;g?Eb1I89cmqFQ)5tdP=8Q;P`^+k5d",
			"1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := b85.Decode(tt.serial)
			assert.NoError(t, err)

			fmt.Println("Name:", tt.name)
			fmt.Println("Serial:", tt.serial)
			parsed, err := Deserialize(data)
			assert.NoError(t, err)
			fmt.Println("Result:", parsed.String())
			fmt.Println("Bits:", parsed.Bits)
		})
	}
}

func TestSerialCompareBuybacks(t *testing.T) {
	var tests = []struct {
		name           string
		serialOriginal string
		serialBuyback  string
	}{
		{
			"L50 Legendary Cooking Ambushing Truck",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			"@Ugy3L+2}Ta0Od!I{*`S=LLLKTRY91;d>K-Z#Y7QzFY8(O",
		},
		{
			"L50 Legendary Ambushing Truck",
			"@Ugy3L+2}TYgjMogxi7Hg07IhPq4>b?9sX3@zs9y*",
			"@Ugy3L+2}Ta0Od!H/&7hp9LM3WZH&OXe^H7_bgUW^ag#Z",
		},
		{
			"L49 Uncommon Playful Kitty",
			"@Ugct)%FmVuJXn{hb3U#POJ!&6nQ*lsxP_0lm5d",
			"@Ugct)%FmVuN0uhE5C^V{2hg#I5_MtWv2ek*)3Uw0!",
		},
		{
			"L49 Common Karkadann",
			"@UgzR8/2__CAOuq;Eiz?Kj9yO^ss8y(TsD20",
			"@UgzR8/2__DrOd!Jad!WClLM`f1lbVBCg=&ZDhX4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataOriginal, err := b85.Decode(tt.serialOriginal)
			assert.NoError(t, err)
			parsedOriginal, err := Deserialize(dataOriginal)
			assert.NoError(t, err)

			dataBuyback, err := b85.Decode(tt.serialBuyback)
			assert.NoError(t, err)
			parsedBuyback, err := Deserialize(dataBuyback)
			assert.NoError(t, err)

			fmt.Println("Original:", parsedOriginal.String())
			fmt.Println("Buyback: ", parsedBuyback.String())
		})
	}
}
