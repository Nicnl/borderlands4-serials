package codex_loader

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/serial"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	Name                   string `json:"itemName"`
	Type                   string `json:"type"`
	Serial                 string `json:"partString"`
	Decoded                string `json:"-"`
	Manufacturer           string `json:"manufacturer"`
	ManufacturerWeaponType string `json:"manufacturerWeaponType"`
	WeaponType             string `json:"weaponType"`
	Rarity                 string `json:"rarity"`
}

type _codex struct {
	Items []Item
}

var (
	Codex           = &_codex{}
	SkipFailedItems = false
)

type LoadedItem struct {
	Type                   string
	Manufacturer           string
	WeaponType             string
	ManufacturerWeaponType string
	Name                   string
	Serial                 string
	Bits                   string
	Error                  string
	DebugOutput            string

	Parsed serial.Serial

	Score int64
}

func (c *_codex) Load(jsonPath string) ([]LoadedItem, int64, error) {
	rawJson, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, -1, err
	}

	var items []Item
	err = json.Unmarshal(rawJson, &items)
	if err != nil {
		return nil, -1, err
	}

	var (
		nbOk   int64
		nbFail int64
	)

	loadedItems := make([]LoadedItem, 0)
	for _, item := range items {
		item.Serial = strings.Trim(item.Serial, "\"")
		item.Serial = strings.Trim(item.Serial, ",")
		item.Serial = strings.Trim(item.Serial, "\"")
		item.Serial = strings.Trim(item.Serial, ",")

		//fmt.Println("Decoding", item.Name, item.Serial)

		data, err := b85.Decode(item.Serial)
		if err != nil {
			fmt.Fprint(os.Stderr, "B85 decode error:", err)
			//nbFail++
			continue
		}

		parsed, err := serial.Deserialize(data)
		if err != nil {
			fmt.Fprint(os.Stderr, "Tokenize error:", err)
			nbFail++

			if SkipFailedItems {
				continue
			}

			loadedItems = append(loadedItems, LoadedItem{
				Type:                   item.Type,
				Manufacturer:           item.Manufacturer,
				WeaponType:             item.WeaponType,
				ManufacturerWeaponType: item.ManufacturerWeaponType,
				Name:                   item.Name,
				Serial:                 item.Serial,
				Bits:                   parsed.Bits,
				Error:                  err.Error(),
				DebugOutput:            parsed.String(),
			})
			continue
		} else {
			loadedItems = append(loadedItems, LoadedItem{
				Type:                   item.Type,
				Manufacturer:           item.Manufacturer,
				WeaponType:             item.WeaponType,
				ManufacturerWeaponType: item.ManufacturerWeaponType,
				Name:                   item.Name,
				Serial:                 item.Serial,
				Bits:                   parsed.Bits,
				Error:                  "",
				DebugOutput:            parsed.String(),
				Parsed:                 parsed,
			})
		}

		item.Decoded = parsed.String()
		//fmt.Println("Decoded:", item.Decoded)
		nbOk++
	}

	fmt.Println("nbOk:", nbOk)
	fmt.Println("nbFail:", nbFail)

	return loadedItems, nbFail, nil
}
