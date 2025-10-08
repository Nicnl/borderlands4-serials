package b85

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		serial string
		hex    string
	}{
		{
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			"21070601906270443339b05391542b85764567854f857785430567054e85638400",
		},
		{
			"@Ugy3L+2}Ta0Od!I{*`S=LLLKTRY91;d>K-Z#Y7QzFY8(O",
			"0106072145706290334304321539059b6457b842f854785630547857e8547056563840",
		},
		{
			"@Ugy3L+2}TYgjMogxi7Hg07IhPq4>b?9sX3@zs9y*",
			"01060721447062905330eb31452a5491a9c8ae68adf0acf03a159c9515fa10",
		},
		{
			"@Ugy3L+2}Ta0Od!H/&7hp9LM3WZH&OXe^H7_bgUW^ag#Z",
			"01060721457062901e430432153905b38a56a4420a9f8aec59d90acf5fa153c1a100",
		},
		{
			"@Ugct)%FmVuJXn{hb3U#POJ!&6nQ*lsxP_0lm5d",
			"32c0182186088e0cc5428116a215ae5056f856bcde0a8ecaec0ab50a88",
		},
		{
			"@Ugct)%FmVuN0uhE5C^V{2hg#I5_MtWv2ek*)3Uw0!",
			"32c01821a6088e0c686188400a552c14c52b5ae1a86c856fabe0ade080c8ae50",
		},
		{
			"@UgwSAs2}TYgOz#USjp~P5)S(jfsJ*DNsIaI@g+bLpr9!Pj^+J_H00",
			"01060b214470629054d0f7325ae157b1af48852b4d15bd15d2150d1582ab82a142b542a5c2a942af00",
		},
		{
			"@UglGc",
			"2116c0",
		},
		{
			"@Ugr$lGm/)}}!dNJvM-}RPG}?q38r1nh0{{",
			"6051a5210427061905141a433e57a8e258215b2c429f2b58c000",
		},
	}

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
