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

type _problematicItem struct {
	Name       string
	Serial     string
	DoneString string
	Error      string
}

func (c *_codex) Load(jsonPath string) ([]_problematicItem, error) {
	rawJson, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}

	var items []Item
	err = json.Unmarshal(rawJson, &items)
	if err != nil {
		return nil, err
	}

	var (
		nbOk   int
		nbFail int
	)

	problematicItems := make([]_problematicItem, 0)
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
		decoded, err := tokenizer.Parse()
		if err != nil {
			fmt.Fprint(os.Stderr, "Tokenize error:", err)
			nbFail++
			problematicItems = append(problematicItems, _problematicItem{
				Name:       item.Name,
				Serial:     item.Serial,
				DoneString: tokenizer.DoneString(),
				Error:      err.Error(),
			})
			continue
		}

		item.Decoded = decoded
		fmt.Println("Decoded:", item.Decoded)
		nbOk++
	}

	fmt.Println("nbOk:", nbOk)
	fmt.Println("nbFail:", nbFail)

	return problematicItems, nil
}
