package codex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodex(t *testing.T) {
	var (
		problematicItems []_problematicItem
		err              error
	)
	t.Run("LOAD", func(t *testing.T) {
		problematicItems, err = Codex.Load("C:\\Users\\Nicnl\\GolandProjects\\borderlands_4_serials\\codex\\resources\\bl4-serial-matches.json")
		assert.NoError(t, err)
	})

	t.Run(fmt.Sprintf("%d_fails", len(problematicItems)), func(t *testing.T) {
		for _, item := range problematicItems {
			t.Run(item.Name, func(t *testing.T) {
				t.Fail()
				fmt.Println("Problematic item:", item.Name)
				fmt.Println("Serial:", item.Serial)
				fmt.Println("Bits:", item.DoneString)
				fmt.Println("Error:", item.Error)
			})
		}
	})

}
