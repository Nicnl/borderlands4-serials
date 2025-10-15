package codex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindItemType(t *testing.T) {
	item, err := Deserialize("@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00")
	assert.NoError(t, err)

	itemType, found := item.Type()
	assert.True(t, found)
	fmt.Println(itemType)
	assert.Equal(t, "jakobs", itemType.Manufacturer)
	assert.Equal(t, "sniper", itemType.Type)
}
