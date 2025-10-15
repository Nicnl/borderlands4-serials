package codex

import (
	"embed"
	"encoding/json"
	"fmt"
)

type BaseBarrel struct {
	ManufacturerIndex uint32
	BaseIndex         uint32
	BarrelIndex       uint32
}

type BaseBarrelJson struct {
	PartStr     string `json:"part_str"`
	Serial      string `json:"serial"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stats       string `json:"stats"`
}

type BaseBarrelInfos struct {
	Type        string     `json:"type"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Stats       string     `json:"stats"`
	Item        *Item      `json:"-"`
	BaseBarrel  BaseBarrel `json:"-"`
}

var (
	//go:embed database/parts/base_barrel
	partsBaseBarrel embed.FS

	Barrels = make(map[BaseBarrel]BaseBarrelInfos) // key: ManufacturerIndex<<16 | BarrelIndex
)

func init() {
	// Walk all JSON files in the embedded filesystem
	files, err := partsBaseBarrel.ReadDir("database/parts/base_barrel")
	if err != nil {
		panic(err)
	}

	var barrels []BaseBarrelInfos

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		data, err := partsBaseBarrel.ReadFile("database/parts/base_barrel/" + file.Name())
		if err != nil {
			panic(err)
		}

		var barrel BaseBarrelJson
		err = json.Unmarshal(data, &barrel)
		if err != nil {
			panic(err)
		}

		item, err := Deserialize(barrel.Serial)
		if err != nil {
			panic(err)
		}

		barrels = append(barrels, BaseBarrelInfos{
			Type:        barrel.Type,
			Name:        barrel.Name,
			Description: barrel.Description,
			Stats:       barrel.Stats,
			Item:        item,
		})
	}

	if CODEX_PRINT_INIT_ITEM_TYPES {
		fmt.Println("Loaded", len(barrels), "barrel parts")
	}

	for _, barrel := range barrels {
		manufacturerIndex, found := barrel.Item.FindIntAtPos(0)
		baseIndex := barrel.Item.FindPartAtPos(0, false)
		barrelIndex := barrel.Item.FindPartAtPos(1, false)

		if !found || baseIndex == nil || barrelIndex == nil {
			continue
		}

		if CODEX_PRINT_INIT_ITEM_TYPES {
			fmt.Println(barrel.Name, manufacturerIndex, baseIndex, barrelIndex, barrel.Stats)
		}

		baseBarrel := BaseBarrel{
			ManufacturerIndex: manufacturerIndex,
			BaseIndex:         baseIndex.Index,
			BarrelIndex:       barrelIndex.Index,
		}
		barrel.BaseBarrel = baseBarrel

		Barrels[baseBarrel] = barrel
	}
}
