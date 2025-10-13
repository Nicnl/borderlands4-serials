package b85

import (
	"borderlands_4_serials/lib/bit_reader"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	serial string
	hex    string
}{
	{
		"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
		"21070601906270443339b05391542b85764567854f857785430567054e85638400",
	},
	{
		"@Ugy3L+2}Ta0Od!I{*`S=LLLKTRY91;d>K-Z#Y7QzFY8(O",
		"2107060190627045320443339b05391542b85764567854f857785430567054e8563840",
	},
	{
		"@Ugy3L+2}TYgjMogxi7Hg07IhPq4>b?9sX3@zs9y*",
		"210706019062704431eb305391542a4568aec8a9f0acf0ad959c153a15fa10",
	},
	{
		"@Ugy3L+2}Ta0Od!H/&7hp9LM3WZH&OXe^H7_bgUW^ag#Z",
		"21070601906270453204431eb305391542a4568aec8a9f0acf0ad959c153a15fa100",
	},
	{
		"@Ugct)%FmVuJXn{hb3U#POJ!&6nQ*lsxP_0lm5d",
		"2118c0320c8e0886168142c550ae15a2bc56f856ca8e0ade0ab50aec88",
	},
	{
		"@Ugct)%FmVuN0uhE5C^V{2hg#I5_MtWv2ek*)3Uw0!",
		"2118c0320c8e08a640886168142c550ae15a2bc56f856ca8e0ade0ab50aec880",
	},
	{
		"@UgwSAs2}TYgOz#USjp~P5)S(jfsJ*DNsIaI@g+bLpr9!Pj^+J_H00",
		"210b06019062704432f7d054b157e15a2b8548af15bd154d150d15d2a182ab82a542b542af42a9c200",
	},
	{
		"@UglGc",
		"2116c0",
	},
	{
		"@Ugr$lGm/)}}!dNJvM-}RPG}?q38r1nh0{{",
		"21a5516019062704431a1405e2a8573e2c5b2158582b9f42c000",
	},
	{
		"@Ugr$lGm/)}}!dNJvM-}RPG}?q38r1nh0{{",
		"21a5516019062704431a1405e2a8573e2c5b2158582b9f42c000",
	},
}

func TestDecode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.serial, func(t *testing.T) {
			decoded, err := Decode(tt.serial)
			if err != nil {
				t.Fatalf("Decode() error = %v", err)
			}
			if fmt.Sprintf("%x", decoded) != tt.hex {
				t.Errorf("Decode() = %x, want %s", decoded, tt.hex)
			}
		})
	}
}

func TestDecodeRobust(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.serial, func(t *testing.T) {
			decoded, err := Decode(tt.serial + "\"\"\"\"\",,,,,,,,,,,,")
			if err != nil {
				t.Fatalf("Decode() error = %v", err)
			}
			if fmt.Sprintf("%x", decoded) != tt.hex {
				t.Errorf("Decode() = %x, want %s", decoded, tt.hex)
			}
		})
	}
}

func TestDecode2(t *testing.T) {
	serials := []string{
		"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
		"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf`00",
		"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$av=Z",
		"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$G64",
		"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l33L00",
		"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!XN+",
		"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}ce_00",
		"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}ce^00",
		"@Ugy3L+2}TYg%$yC%i7M2gZldO)@*w~",
		"@Ugy3L+2}TYg%$yC%i7M2gZldO)@*n^",
		"@Ugy3L+2}TYg%$yC%i7M2gZldNP00",
		"@Ugy3L+2}TYg%$yC%i7M2gZldNO00",
		"@Ugy3L+2}TYg%$yC%i7M2gZXy5",
		"@Ugy3L+2}TYg%$yC%i7M2gE&%",
		"@Ugy3L+2}TYg%$yC%i7M0~00",
		"@Ugy3L+2}TYg%$yC%i7Es",
		"@Ugy3L+2}TYg%$yC%i2w",
		"@Ugy3L+2}TY8",
		"@Ugy3L+2@aC}/NsC0/Nnmg",
		"@Ugy3L+2}S?",
		"@Ugy3L+2?hW",
		"@Ugx~-",
		"@Ugdh",
	}

	for _, serial := range serials {
		data, err := Decode(serial)
		assert.NoError(t, err)

		bitStream := bit_reader.NewBitReader(data)
		fmt.Println(bitStream.StringAfter())
	}
}
