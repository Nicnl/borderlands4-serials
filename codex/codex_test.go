package codex

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodex(t *testing.T) {
	var (
		loadedItems []_loadedItem
		err         error
		nbFails     int64
	)
	t.Run("LOAD", func(t *testing.T) {
		loadedItems, nbFails, err = Codex.Load("C:\\Users\\Nicnl\\GolandProjects\\borderlands_4_serials\\codex\\resources\\bl4-serial-matches.json")
		assert.NoError(t, err)
	})

	t.Run(fmt.Sprintf("%d_fails", nbFails), func(t *testing.T) {
		for _, item := range loadedItems {
			t.Run(item.Type+"__"+item.Error+"__"+item.Name, func(t *testing.T) {
				if item.Error != "" {
					t.Fail()
				}
				fmt.Println("Problematic item:", item.Name+" ("+item.Type+")")
				fmt.Println("Serial:", item.Serial)
				fmt.Println("Bits:", item.DoneString)
				fmt.Println("Decoded:", item.DebugOutput)
				fmt.Println("Error:", item.Error)
			})
		}
	})

	// Sort loadedItems by shortest serial length
	if len(loadedItems) > 0 {
		sort.Slice(loadedItems, func(i, j int) bool {
			return len(loadedItems[i].Serial) < len(loadedItems[j].Serial)
		})

		// Print top 10 shortest problematic items

		t.Run("TOP_10_SHORTEST_FAILS", func(t *testing.T) {
			limit := 2500
			if len(loadedItems) < limit {
				limit = len(loadedItems)
			}
			for i := 0; i < limit; i++ {
				item := loadedItems[i]
				t.Run(item.Type+"__"+item.Error+"__"+item.Name, func(t *testing.T) {
					if item.Error != "" {
						t.Fail()
					}
					fmt.Println("Problematic item:", item.Name+" ("+item.Type+")")
					fmt.Println("Serial:", item.Serial)
					fmt.Println("Bits:", item.DoneString)
					fmt.Println("Decoded:", item.DebugOutput)
					fmt.Println("Error:", item.Error)
				})
			}
		})
	}
}
