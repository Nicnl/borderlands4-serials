package serial

import (
	"borderlands_4_serials/b4s/b85"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLevel(t *testing.T) {
	data, err := b85.Decode("@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00")
	assert.NoError(t, err)

	s, err := Deserialize(data)
	assert.NoError(t, err)

	level, found := s.FindLevel()
	assert.True(t, found)
	assert.Equal(t, uint32(50), level)
}
