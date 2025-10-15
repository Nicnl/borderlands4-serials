package codex

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/serial"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JsonItem struct {
	Name                   string `json:"itemName"`
	Type                   string `json:"type"`
	Serial                 string `json:"partString"`
	Decoded                string `json:"-"`
	Manufacturer           string `json:"manufacturer"`
	ManufacturerWeaponType string `json:"manufacturerWeaponType"`
	WeaponType             string `json:"weaponType"`
	Rarity                 string `json:"rarity"`

	Item *Item `json:"-"`
}

type _codex struct {
	JsonItems []JsonItem
}

var (
	Codex           = &_codex{}
	SkipFailedItems = false
)

func (c *_codex) Load(jsonPath string) ([]Item, int64, error) {
	rawJson, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, -1, err
	}

	var jsonItems []JsonItem
	err = json.Unmarshal(rawJson, &jsonItems)
	if err != nil {
		return nil, -1, err
	}

	var (
		nbOk   int64
		nbFail int64
	)

	Codex.JsonItems = make([]JsonItem, 0, len(jsonItems))
	for _, jsonItem := range jsonItems {
		jsonItem.Serial = strings.Trim(jsonItem.Serial, "\"")
		jsonItem.Serial = strings.Trim(jsonItem.Serial, ",")
		jsonItem.Serial = strings.Trim(jsonItem.Serial, "\"")
		jsonItem.Serial = strings.Trim(jsonItem.Serial, ",")

		//fmt.Println("Decoding", item.Name, item.Serial)

		item, err := Deserialize(jsonItem.Serial)
		if err != nil {
			fmt.Fprint(os.Stderr, "Serial decode error:", err)
			//nbFail++
			continue
		}
		jsonItem.Item = item

		if err != nil {
			fmt.Fprint(os.Stderr, "Tokenize error:", err)
			nbFail++

			if !SkipFailedItems {
				Codex.JsonItems = append(Codex.JsonItems, jsonItem)
			}
		} else {
			Codex.JsonItems = append(Codex.JsonItems, jsonItem)
		}

		nbOk++
	}

	fmt.Println("nbOk:", nbOk)
	fmt.Println("nbFail:", nbFail)

	return loadedItems, nbFail, nil
}
