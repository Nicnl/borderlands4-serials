package codex

import (
	"borderlands_4_serials/lib/b85"
	"borderlands_4_serials/lib/serial_tokenizer"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	Name         string `json:"itemName"`
	Type         string `json:"type"`
	Serial       string `json:"partString"`
	Decoded      string `json:"-"`
	Manufacturer string `json:"manufacturer"`
	WeaponType   string `json:"weaponType"`
	Rarity       string `json:"rarity"`
}

type _codex struct {
	Items []Item
}

var Codex = &_codex{}

type _loadedItem struct {
	Type        string
	Name        string
	Serial      string
	DoneString  string
	Error       string
	DebugOutput string

	Score int64
}

func (c *_codex) Load(jsonPath string) ([]_loadedItem, int64, error) {
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

	loadedItems := make([]_loadedItem, 0)
	for _, item := range items {
		item.Serial = strings.Trim(item.Serial, "\"")
		item.Serial = strings.Trim(item.Serial, ",")

		fmt.Println("Decoding", item.Name, item.Serial)

		data, err := b85.Decode(item.Serial)
		if err != nil {
			fmt.Fprint(os.Stderr, "B85 decode error:", err)
			//nbFail++
			continue
		}

		tokenizer := serial_tokenizer.NewTokenizer(data)
		_, decoded, err := tokenizer.Parse()
		if err != nil {
			fmt.Fprint(os.Stderr, "Tokenize error:", err)
			nbFail++
			loadedItems = append(loadedItems, _loadedItem{
				Type:        item.Type,
				Name:        item.Name,
				Serial:      item.Serial,
				DoneString:  tokenizer.DoneString(),
				Error:       err.Error(),
				DebugOutput: decoded,
			})
			continue
		} else {
			loadedItems = append(loadedItems, _loadedItem{
				Type:        item.Type,
				Name:        item.Name,
				Serial:      item.Serial,
				DoneString:  tokenizer.DoneString(),
				Error:       "",
				DebugOutput: decoded,
			})
		}

		item.Decoded = decoded
		fmt.Println("Decoded:", item.Decoded)
		nbOk++
	}

	fmt.Println("nbOk:", nbOk)
	fmt.Println("nbFail:", nbFail)

	return loadedItems, nbFail, nil
}
