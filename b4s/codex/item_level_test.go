package codex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLevel(t *testing.T) {
	item, err := Deserialize("@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00")
	assert.NoError(t, err)

	level, found := item.FindLevel()
	assert.True(t, found)
	assert.Equal(t, uint32(50), level)
}
