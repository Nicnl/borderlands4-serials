package serial

import (
	"borderlands_4_serials/b4s/b85"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindItemType(t *testing.T) {
	data, err := b85.Decode("@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00")
	assert.NoError(t, err)

	s, err := Deserialize(data)
	assert.NoError(t, err)

	itemType, found := s.FindItemType()
	assert.True(t, found)
	fmt.Println(itemType)
	assert.Equal(t, "jakobs", itemType.Manufacturer)
	assert.Equal(t, "sniper", itemType.Type)
}
