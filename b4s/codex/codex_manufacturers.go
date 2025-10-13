package codex

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/serial"
	_ "embed"
	"encoding/json"
	"fmt"
)

type ItemType struct {
	Manufacturer string `json:"manufacturer"`
	Type         string `json:"type"`
	Serial       string `json:"serial"`
	Index        uint32 `json:"-"` // Extracted from Serial
}

var (
	itemTypes        []ItemType
	itemTypesByIndex = make(map[uint32]*ItemType)
	itemTypesByHash  = make(map[string]*ItemType)

	//go:embed database/item_types.json
	itemTypesJson []byte
)

func init() {
	fmt.Println("# CODEX: Loading item types...")
	// Decode the embedded JSON data
	err := json.Unmarshal(itemTypesJson, &itemTypes)
	if err != nil {
		panic(err)
	}

	for i := range itemTypes {
		var (
			itemType   = &itemTypes[i]
			serialData serial.Serial
		)

		// Deserialize the item data
		{
			data, err := b85.Decode(itemType.Serial)
			if err != nil {
				panic("error decoding b85 for item type: " + itemType.Manufacturer + ":" + itemType.Type + " : " + err.Error())
			}

			serialData, err = serial.Deserialize(data)
			if err != nil {
				panic("error deserializing serial for item type: " + itemType.Manufacturer + ":" + itemType.Type + " : " + err.Error())
			}

			if len(serialData.Blocks) == 0 {
				panic("no blocks in serial data for item type: " + itemType.Manufacturer + ":" + itemType.Type)
			}
		}

		// Map
		if _, conflict := itemTypesByHash[itemType.Serial]; conflict {
			panic("conflict in itemTypesByHash for serial: " + itemType.Serial + " (" + itemType.Manufacturer + ":" + itemType.Type + ")")
		}

		itemType.Index = serialData.Blocks[0].Value
		itemTypesByIndex[itemType.Index] = itemType
		itemTypesByHash[itemType.Serial] = itemType
		fmt.Println("  - ", itemType.Manufacturer+":"+itemType.Type, "=>", itemType.Index)
	}
}
