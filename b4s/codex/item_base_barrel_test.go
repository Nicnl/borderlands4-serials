package codex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemFindBaseBarrel(t *testing.T) {
	item, err := Deserialize("@UgzR8/2}TYgOx+18jVjck67{G`WolA$P_a<CP`gm;Q1cJ")
	assert.NoError(t, err)

	infos, found := item.BaseBarrel()
	assert.True(t, found)
	fmt.Println(infos)
}
