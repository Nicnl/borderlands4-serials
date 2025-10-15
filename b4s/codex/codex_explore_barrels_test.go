package codex

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/serial"
	"borderlands_4_serials/b4s/serial_datatypes/part"
	"borderlands_4_serials/b4s/serial_tokenizer"
	"fmt"
	"strings"
	"testing"
)

func _serialsToYaml(serials []string) {
	output := strings.Builder{}
	slotCounter := 0
	for _, serial := range serials {
		if strings.HasPrefix(serial, "@U") {
			serial = strings.Trim(serial, "\r")
			serial = strings.Trim(serial, "\n")
			output.WriteString("        slot_" + fmt.Sprintf("%d", slotCounter) + ":\n")
			output.WriteString("          serial: '" + serial + "'\n")
			slotCounter++
		}
	}

	fmt.Println(output.String())
}

func TestConstructBasesWithParts(t *testing.T) {
	type WeaponBase struct {
		manufacturerIndex uint32
		baseIndex         uint32
	}

	allParts := make(map[string]part.Part)
	weaponBases := make(map[WeaponBase]map[string]bool)

	t.Run("EXTRACT_BASES", func(t *testing.T) {
		for _, jsonItem := range Codex.JsonItems {

			manufacturerIndex, found := jsonItem.Item.FindIntAtPos(0)
			if !found {
				continue
			}

			itemType, found := GetItemTypeByIndex(manufacturerIndex)
			if !found {
				continue
			}

			switch itemType.Type {
			case "pistol", "assault_rifle", "smg", "shotgun", "sniper":
				// Nothing to do
			default:
				continue
			}

			baseIndexPart := jsonItem.Item.FindPartAtPos(0, false)
			if baseIndexPart == nil {
				continue
			}
			if baseIndexPart.SubType != part.SUBTYPE_NONE {
				continue
			}

			weaponBase := WeaponBase{
				manufacturerIndex: manufacturerIndex,
				baseIndex:         baseIndexPart.Index,
			}

			if _, found := weaponBases[weaponBase]; !found {
				weaponBases[weaponBase] = make(map[string]bool)
			}

			pos := 1
			for {
				p := jsonItem.Item.FindPartAtPos(pos, true)
				if p == nil {
					break
				}
				allParts[p.String()] = *p
				weaponBases[weaponBase][p.String()] = true
				pos++
			}
		}
	})

	fmt.Println("weaponBases =")
	serialsToTest := make([]string, 0)
	t.Run("GENERATE_SERIALS", func(t *testing.T) {
		for weaponBase, parts := range weaponBases {
			//fmt.Printf("  // ManufacturerIndex: %d, BaseIndex: %d\n", weaponBase.manufacturerIndex, weaponBase.baseIndex)
			//fmt.Printf("  {\n")

			//for partStr := range parts {
			//	part := allParts[partStr]
			//	fmt.Printf("      %s,\n", part.String())
			//}
			//fmt.Printf("    },\n")

			for partStr := range parts {
				curPart, found := allParts[partStr]
				if !found {
					continue
				}

				encoded := b85.Encode(serial.Serialize([]serial.Block{
					{Token: serial_tokenizer.TOK_VARINT, Value: weaponBase.manufacturerIndex},
					{Token: serial_tokenizer.TOK_SEP2},
					{Token: serial_tokenizer.TOK_VARINT, Value: 0}, // Unknown, always zero
					{Token: serial_tokenizer.TOK_SEP2},
					{Token: serial_tokenizer.TOK_VARINT, Value: 1}, // Unknown, always one before the level
					{Token: serial_tokenizer.TOK_SEP2},
					{Token: serial_tokenizer.TOK_VARINT, Value: 50}, // Level 50
					{Token: serial_tokenizer.TOK_SEP1},
					{Token: serial_tokenizer.TOK_SEP1},
					{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: weaponBase.baseIndex, SubType: part.SUBTYPE_NONE}},
					{Token: serial_tokenizer.TOK_PART, Part: curPart},
					{Token: serial_tokenizer.TOK_SEP1},
				}))
				fmt.Printf("  // ManufacturerIndex: %d, BaseIndex: %d\n", weaponBase.manufacturerIndex, weaponBase.baseIndex)
				fmt.Printf("  \"%s\",\n", encoded)
				serialsToTest = append(serialsToTest, encoded)
			}

		}
	})

	t.Run("PRINT", func(t *testing.T) {
		_serialsToYaml(serialsToTest)
	})
}
