package codex

import (
	"borderlands_4_serials/b4s/b85"
	"fmt"
	"sort"
	"strings"
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
		loadedItems, nbFails, err = Codex.Load("resources/bl4-serial-matches.json")
		assert.NoError(t, err)
	})

	t.Run(fmt.Sprintf("%d_fails", nbFails), func(t *testing.T) {
		for _, item := range loadedItems {
			if item.Error != "" {
				t.Run(item.Type+"__"+item.Error+"__"+item.Name, func(t *testing.T) {
					t.Fail()

					fmt.Println("Problematic item:", item.Name+" ("+item.Type+")")
					fmt.Println("Serial:", item.Serial)
					fmt.Println("Bits:", item.Bits)
					fmt.Println("Decoded:", item.DebugOutput)
					fmt.Println("Error:", item.Error)

					if strings.Contains(strings.ReplaceAll(item.Bits, " ", ""), "101100001011100001011110001010111101001000101011111111000010101000010010001010100001100000101010101110000010101111111000001010101111100000101010001101000010101110110100001000") {
						//panic("found")
					}
				})
			}
		}
	})

	t.Run(fmt.Sprintf("%d_success", len(loadedItems)-int(nbFails)), func(t *testing.T) {
		for _, item := range loadedItems {
			if item.Error == "" {
				t.Run(item.Type+"__"+item.Error+"__"+item.Name, func(t *testing.T) {
					fmt.Println("Problematic item:", item.Name+" ("+item.Type+")")
					fmt.Println("Serial:", item.Serial)
					fmt.Println("Bits:", item.Bits)
					fmt.Println("Decoded:", item.DebugOutput)
					fmt.Println("Error:", item.Error)

					if strings.Contains(strings.ReplaceAll(item.Bits, " ", ""), "101100001011100001011110001010111101001000101011111111000010101000010010001010100001100000101010101110000010101111111000001010101111100000101010001101000010101110110100001000") {
						//panic("found")
					}
				})
			}
		}
	})

	if len(loadedItems) > 0 {
		// Sort loadedItems by shortest serial length
		sort.Slice(loadedItems, func(i, j int) bool {
			return len(loadedItems[i].Serial) > len(loadedItems[j].Serial)
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
					fmt.Println("Bits:", item.Bits)
					fmt.Println("Decoded:", item.DebugOutput)
					fmt.Println("Error:", item.Error)
				})
			}
		})

		// For each loadedItems, calculate their score
		// The score is how many continous zeroes starting from right
		for i := range loadedItems {
			item := &loadedItems[i]
			bytes, err := b85.Decode(item.Serial)
			if err != nil {
				continue
			}
			binStr := ""
			for _, b := range bytes {
				binStr += fmt.Sprintf("%08b", b)
			}
			score := int64(0)
			for j := len(binStr) - 1; j >= 0; j-- {
				if binStr[j] == '0' {
					score++
				} else {
					break
				}
			}
			item.Score = score
		}

		// Sort loadedItems by score
		sort.Slice(loadedItems, func(i, j int) bool {
			return loadedItems[i].Score < loadedItems[j].Score
		})

		t.Run("TOP_10_HIGHEST_SCORE", func(t *testing.T) {
			limit := 10
			if len(loadedItems) < limit {
				limit = len(loadedItems)
			}
			for i := 0; i < limit; i++ {
				item := loadedItems[i]
				t.Run(fmt.Sprintf("%d__%s__%s", item.Score, item.Type, item.Name), func(t *testing.T) {
					if item.Error != "" {
						t.Fail()
					}
					fmt.Println("Problematic item:", item.Name+" ("+item.Type+")")
					fmt.Println("Score:", item.Score)
					fmt.Println("Serial:", item.Serial)
					fmt.Println("Bits:", item.Bits)
					fmt.Println("Decoded:", item.DebugOutput)
					fmt.Println("Error:", item.Error)
				})
			}
		})
	}
}
