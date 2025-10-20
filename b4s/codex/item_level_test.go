package codex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLevel(t *testing.T) {
	var tests = []struct {
		name          string
		b85           string
		found         bool
		expectedLevel uint32
	}{
		{
			"Legendary Cooking Ambushing Truck",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			true,
			50,
		},
		{
			"Uncommon Playful Kitty",
			"@Ugct)%FmVuJXn{hb3U#POJ!&6nQ*lsxP_0lm5d",
			true,
			49,
		},
		{
			"Knife 3 skill",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3OODgg",
			true,
			22,
		},
		{
			"Green Unseen Xiuhcoatl",
			"@Ugy3L+2}TMcjNb(cjVjck8WpL1s7>WTg+kRrl/uj",
			true,
			34,
		},
		{
			"Legendary Cooking Ambushing Truck",
			"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			true,
			50,
		},
		{
			"Jakobs Sniper Top Square Simple 5",
			"@Ugy3L+2}TYgT#^cvMir`2hg#I5@}cgb=Ak+@2XzZ/4gm",
			true,
			50,
		},
		{
			"Legendary Cooking Ambushing Truck SMALL",
			"@Ugy3L+35F42=4?<-RG/)a6EzQ&4/NX}1~mtj3pEY_",
			true,
			12,
		},
		{
			"Vex Classmod: 1 square  1 bottle",
			"@Ug!pHG38o5YPb#KC)h-nP",
			true,
			37,
		},
		{
			"Vex Classmod: 1 arm 1 bottle + firmware jacked",
			"@Ug!pHG38o5YT`HzQ)$V@)",
			true,
			37,
		},
		{
			"valkyr, weapon with no level, received via discord",
			"@Ugr$lG7-8sL(4z`<KALPY4GrpidjS",
			false,
			0,
		},
		{
			"Knife Sho Kunai 4 Skill",
			"@Ugr$WBm/)}}!bEtWObu#%Z$Os-",
			true,
			50,
		},
		{
			"Nova Pointed Pandoran Memento",
			"@Ugr$oHm/)}}!q~Z>NfoMnaX/6AQ2S`2h@waU",
			true,
			50,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("L%d_%s", tt.expectedLevel, tt.name), func(t *testing.T) {
			item, err := Deserialize(tt.b85)
			assert.NoError(t, err)

			level, found := item.Level()
			assert.Equal(t, tt.found, found)
			assert.Equal(t, tt.expectedLevel, level)
		})
	}
}
