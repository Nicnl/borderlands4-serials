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

	var shortestProblem _problematicItem
	t.Run(fmt.Sprintf("%d_fails", len(problematicItems)), func(t *testing.T) {
		for _, item := range problematicItems {
			t.Run(item.Type+"__"+item.Error+"__"+item.Name, func(t *testing.T) {
				t.Fail()
				fmt.Println("Problematic item:", item.Name, item.Type)
				fmt.Println("Serial:", item.Serial)
				fmt.Println("Bits:", item.DoneString)
				fmt.Println("Error:", item.Error)

				if shortestProblem.Serial == "" || len(item.Serial) < len(shortestProblem.Serial) {
					shortestProblem = item
				}
			})
		}
	})

	if shortestProblem.Serial != "" {
		t.Run("SHORTEST_FAIL", func(t *testing.T) {
			t.Fail()
			fmt.Println("Shortest problematic item:", shortestProblem.Name, shortestProblem.Type)
			fmt.Println("Serial:", shortestProblem.Serial)
			fmt.Println("Bits:", shortestProblem.DoneString)
			fmt.Println("Error:", shortestProblem.Error)
		})
	}

}
