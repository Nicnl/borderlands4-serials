package codex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPartAtPos_NoSplit(t *testing.T) {
	var tests = []struct {
		name         string
		b85          string
		deserialized string
		splitLists   bool
		parts        []string
	}{
		{
			"Knife 4 Skill / NoSplit",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3OOD#<4R",
			"267, 0, 1, 22| 2, 274|| {7} {1} {245:[23 39 69 79]}|",
			false,
			[]string{
				"{7}",
				"{1}",
				"{245:[23 39 69 79]}",
			},
		},
		{
			"Knife 4 Skill / Slit",
			"@Ugr$WBm/$!m!X=5&qXxA;nj3OOD#<4R",
			"267, 0, 1, 22| 2, 274|| {7} {1} {245:[23 39 69 79]}|",
			true,
			[]string{
				"{7}",
				"{1}",
				"{245:23}",
				"{245:39}",
				"{245:69}",
				"{245:79}",
			},
		},
		{
			"Kill Sprint Repkit / NoSplit",
			"@Uge8#%m/)}}!qBXsM-}RPG}(k28r1n{WC;Q",
			"290, 0, 1, 50| 2, 2708|| {7} {2} {243:[105 100]} {1} {243:[82 8]}|",
			false,
			[]string{
				"{7}",
				"{2}",
				"{243:[105 100]}",
				"{1}",
				"{243:[82 8]}",
			},
		},
		{
			"Kill Sprint Repkit / Split",
			"@Uge8#%m/)}}!qBXsM-}RPG}(k28r1n{WC;Q",
			"290, 0, 1, 50| 2, 2708|| {7} {2} {243:[105 100]} {1} {243:[82 8]}|",
			true,
			[]string{
				"{7}",
				"{2}",
				"{243:105}",
				"{243:100}",
				"{1}",
				"{243:82}",
				"{243:8}",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, err := Deserialize(tt.b85)
			assert.NoError(t, err)
			assert.Equal(t, tt.deserialized, item.Serial.String())

			pos := 0
			for {
				part := item.FindPartAtPos(pos, tt.splitLists)
				pos++
				if part == nil {
					break
				}

				expected := tt.parts[pos-1]
				t.Run(fmt.Sprintf("FindPartAtPos(%d)_"+expected, pos), func(t *testing.T) {
					assert.Equal(t, expected, part.String())
				})
			}

			if pos != len(tt.parts)+1 {
				t.Errorf("expected to find %d parts, but found %d", len(tt.parts), pos-1)
			}
		})
	}
}
