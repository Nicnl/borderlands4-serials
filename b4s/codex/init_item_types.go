package codex

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/serial"
	_ "embed"
	"encoding/json"
	"fmt"
)

const CODEX_PRINT_INIT_ITEM_TYPES = false

var (
	initDone = false

	//go:embed database/item_types.json
	itemTypesJson []byte
)

func init() {
	if initDone {
		return
	}

	itemTypes := make([]ItemType, 0, 100)

	if CODEX_PRINT_INIT_ITEM_TYPES {
		fmt.Println("# CODEX: Initializing item types...")
	}
	// Decode the embedded JSON data
	err := json.Unmarshal(itemTypesJson, &itemTypes)
	if err != nil {
		panic(err)
	}

	for i := range itemTypes {
		var (
			itemType = itemTypes[i]
			blocks   serial.Serial
		)

		// Deserialize the item data
		{
			data, err := b85.Decode(itemType.Serial)
			if err != nil {
				panic("error decoding b85 for item type: " + itemType.Manufacturer + ":" + itemType.Type + " : " + err.Error())
			}

			blocks, _, err = serial.Deserialize(data)
			if err != nil {
				panic("error deserializing serial for item type: " + itemType.Manufacturer + ":" + itemType.Type + " : " + err.Error())
			}

			if len(blocks) == 0 {
				panic("no blocks in serial data for item type: " + itemType.Manufacturer + ":" + itemType.Type)
			}

			itemType.Index = blocks[0].Value
		}

		// Map
		if CODEX_PRINT_INIT_ITEM_TYPES {
			fmt.Println("  - ", itemType.Manufacturer+":"+itemType.Type, "=>", itemType.Index)
		}
		PushItemType(itemType)
	}

	initDone = true
}
