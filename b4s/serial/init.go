package serial

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/codex_loader/codex"
	_ "embed"
	"encoding/json"
)

var (
	initDone = false

	//go:embed database/item_types.json
	itemTypesJson []byte
)

func init() {
	if initDone {
		return
	}

	itemTypes := make([]codex.ItemType, 0, 100)

	//fmt.Println("# CODEX: Initializing item types...")
	// Decode the embedded JSON data
	err := json.Unmarshal(itemTypesJson, &itemTypes)
	if err != nil {
		panic(err)
	}

	for i := range itemTypes {
		var (
			itemType   = itemTypes[i]
			serialData Serial
		)

		// Deserialize the item data
		{
			data, err := b85.Decode(itemType.Serial)
			if err != nil {
				panic("error decoding b85 for item type: " + itemType.Manufacturer + ":" + itemType.Type + " : " + err.Error())
			}

			serialData, err = Deserialize(data)
			if err != nil {
				panic("error deserializing serial for item type: " + itemType.Manufacturer + ":" + itemType.Type + " : " + err.Error())
			}

			if len(serialData.Blocks) == 0 {
				panic("no blocks in serial data for item type: " + itemType.Manufacturer + ":" + itemType.Type)
			}

			itemType.Index = serialData.Blocks[0].Value
		}

		// Map
		//fmt.Println("  - ", itemType.Manufacturer+":"+itemType.Type, "=>", itemType.Index)
		codex.PushItemType(itemType)
	}

	initDone = true
}
