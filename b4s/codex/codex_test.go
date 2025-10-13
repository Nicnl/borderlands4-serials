package codex

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/serial"
	"borderlands_4_serials/b4s/serial_tokenizer"
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodex(t *testing.T) {
	var (
		loadedItems []LoadedItem
		err         error
		nbFails     int64
	)
	t.Run("LOAD", func(t *testing.T) {
		loadedItems, nbFails, err = Codex.Load("database/bl4-serial-matches.json")
		assert.NoError(t, err)
	})

	t.Run(fmt.Sprintf("%d_fails", nbFails), func(t *testing.T) {
		for _, item := range loadedItems {
			if item.Error != "" {
				t.Run(item.Type+"__"+item.Error+"__"+item.Name, func(t *testing.T) {
					t.Fail()

					fmt.Println("Item:", item.Name+" ("+item.Type+")"+item.Manufacturer)
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

					fmt.Println("Item:", item.Name+" ("+item.Type+")"+item.Manufacturer)
					fmt.Println("Serial:", item.Serial)
					fmt.Println("Bits:", item.Bits)
					fmt.Println("Decoded:", item.DebugOutput)
					fmt.Println("Error:", item.Error)
					fmt.Println("Name:", item.Name)
					fmt.Println("Manufacturer:", item.Manufacturer)
					fmt.Println("WeaponType:", item.WeaponType)
					fmt.Println("ManufacturerWeaponType:", item.ManufacturerWeaponType)

					if len(item.Parsed.Blocks) > 0 {
						manufacturerIndex := item.Parsed.Blocks[0].Value
						fmt.Println("manufacturerIndex:", manufacturerIndex)

						if _, ok := itemTypesByIndex[manufacturerIndex]; !ok {
							t.Fatalf("unknown manufacturer index: " + fmt.Sprint(manufacturerIndex))
						}
					}

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
					fmt.Println("Item:", item.Name+" ("+item.Type+")"+item.Manufacturer)
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
					fmt.Println("Item:", item.Name+" ("+item.Type+")"+item.Manufacturer)
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

func TestCodexReserializeRoundtrip(t *testing.T) {
	var (
		loadedItems []LoadedItem
		err         error
	)
	t.Run("LOAD", func(t *testing.T) {
		loadedItems, _, err = Codex.Load("database/bl4-serial-matches.json")
		assert.NoError(t, err)
	})

	t.Run("round_trip", func(t *testing.T) {
		for _, item := range loadedItems {
			if item.Error != "" {
				continue
			}

			t.Run(item.Type+"__"+item.Error+"__"+item.Name, func(t *testing.T) {
				expectedSerial := strings.Trim(item.Serial, "\"")

				serializedData := serial.Serialize(item.Parsed)
				assert.NoError(t, err)

				reserializedB85 := b85.Encode(serializedData)
				assert.Equal(t, expectedSerial, reserializedB85)

				fmt.Println("Decoded:     ", item.DebugOutput)
				fmt.Println("Original:    ", expectedSerial)
				fmt.Println("Reserialized:", reserializedB85)
			})
		}
	})
}

func TestCodesExtractPairSerialsCommonPart(t *testing.T) {
	// Build:
	// []part
	// part -> map[serial]bool
	// serial -> map[part]bool

	var (
		loadedItems []LoadedItem
		err         error
		_           int64
	)
	t.Run("LOAD", func(t *testing.T) {
		loadedItems, _, err = Codex.Load("database/bl4-serial-matches.json")
		assert.NoError(t, err)
	})

	partToItem := make(map[string][]*LoadedItem, 0)
	t.Run("GROUP_BY_PARTS_TO_ITEM", func(t *testing.T) {
		for _, item := range loadedItems {
			if item.Error != "" {
				continue
			}

			for _, block := range item.Parsed.Blocks {
				if block.Token == serial_tokenizer.TOK_PART {
					if len(block.Part.Values) > 0 {
						continue
					}

					partStr := block.Part.String()
					if _, ok := partToItem[partStr]; !ok {
						partToItem[partStr] = make([]*LoadedItem, 0)
					}

					partToItem[partStr] = append(partToItem[partStr], &item)
				}
			}
		}
		fmt.Println("len(partToItem) =", len(partToItem))
	})

	serialToParts := make(map[string][]string, 0)
	serialToItem := make(map[string]*LoadedItem, 0)
	t.Run("GROUP_BY_SERIAL_TO_PARTS", func(t *testing.T) {
		for _, item := range loadedItems {
			if item.Error != "" {
				continue
			}

			serialToItem[item.Serial] = &item

			if _, ok := serialToParts[item.Serial]; !ok {
				serialToParts[item.Serial] = make([]string, 0)
			}

			for _, block := range item.Parsed.Blocks {
				if block.Token == serial_tokenizer.TOK_PART {
					if len(block.Part.Values) > 0 {
						continue
					}

					partStr := block.Part.String()
					serialToParts[item.Serial] = append(serialToParts[item.Serial], partStr)
				}
			}
		}
		fmt.Println("len(serialToParts) =", len(serialToParts))
	})

	partPairs := make(map[string][][2]string, 0)

	for part := range partToItem {
		// For each part, we'll get its items
		itemsHavingThisPart := partToItem[part]

		pairs := make([][2]string, 0)
		partPairs[part] = pairs

		// We base on each item as a pivvot
		for _, pivotItem := range itemsHavingThisPart {
			// For each item having this part, we check if they have other parts in common
			for _, otherItem := range itemsHavingThisPart {
				if pivotItem == otherItem {
					continue
				}

				// We want to verify is the only part in common is `part`
				// If there are common parts, we stop
				hasOtherCommonPart := false

			OUTER:
				for _, pivotPart := range serialToParts[pivotItem.Serial] {
					if pivotPart == part {
						continue
					}

					item1 := serialToItem[pivotItem.Serial]

					for _, otherPart := range serialToParts[otherItem.Serial] {
						if otherPart == part {
							continue
						}

						item2 := serialToItem[otherItem.Serial]

						// If different manufacturers, skip
						if item1.Parsed.Blocks[0].Value != item2.Parsed.Blocks[0].Value {
							continue
						}

						if pivotPart == otherPart {
							hasOtherCommonPart = true
							break OUTER
						}
					}
				}

				//fmt.Println(part, "=>", hasOtherCommonPart)
				if !hasOtherCommonPart {
					partPairs[part] = append(partPairs[part], [2]string{
						pivotItem.Serial,
						otherItem.Serial,
					})
				}
			}
		}
	}

	// Print the results
	for part, pairs := range partPairs {
		if len(pairs) > 0 {
			fmt.Println("Part:", part)
			for _, pair := range pairs {
				// obtain the two items
				item1 := serialToItem[pair[0]]
				item2 := serialToItem[pair[1]]

				fmt.Println("  - ", pair[0], "("+item1.Name+")")
				fmt.Println("  - ", pair[1], "("+item2.Name+")")

				break
			}
			fmt.Println()
		}
	}
}
