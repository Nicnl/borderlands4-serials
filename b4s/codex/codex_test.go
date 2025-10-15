package codex

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/serial"
	"borderlands_4_serials/b4s/serial_datatypes/part"
	"borderlands_4_serials/b4s/serial_tokenizer"
	"fmt"
	"testing"
)

func TestCodesMachin(t *testing.T) {
	// Collect all parts
	mapParts := make(map[string]part.Part)
	for _, jsonItem := range Codex.JsonItems {
		pos := 1
		for {
			p := jsonItem.Item.FindPartAtPos(pos, true)
			pos++
			if p == nil {
				break
			}

			if p.SubType != part.SUBTYPE_NONE {
				continue
			}

			mapParts[p.String()] = *p
		}
	}

	// Group parts by barrel base
	baseToParts := make(map[BaseBarrel]map[uint32]bool)
	for _, jsonItem := range Codex.JsonItems {
		base, found := jsonItem.Item.BaseBarrel()
		if !found {
			//fmt.Println("no base barrel found for " + jsonItem.Name)
			continue
		}

		if _, exists := baseToParts[base.BaseBarrel]; !exists {
			baseToParts[base.BaseBarrel] = make(map[uint32]bool)
		}

		pos := 1
		for {
			p := jsonItem.Item.FindPartAtPos(pos, true)
			pos++
			if p == nil {
				break
			}

			if p.SubType != part.SUBTYPE_NONE {
				continue
			}

			if p.Index == base.BaseBarrel.BarrelIndex {
				// Skip the barrel part itself
				continue
			}

			if _, exists := baseToParts[base.BaseBarrel][p.Index]; !exists {
				baseToParts[base.BaseBarrel][p.Index] = true
			}
		}
	}

	combinations := 0
	fmt.Println("Total parts", len(mapParts))
	fmt.Println("Total bases", len(baseToParts))
	fmt.Println()
	fmt.Print("ALl parts =")
	for _, part := range mapParts {
		fmt.Print(" ", part.String())
	}
	fmt.Println()

	fmt.Println()
	for baseBarrel, parts := range baseToParts {
		combinations += len(parts)
		infos := Barrels[baseBarrel]
		fmt.Println("Base:", infos.Name, infos.BaseBarrel.ManufacturerIndex, infos.BaseBarrel.BaseIndex, infos.BaseBarrel.BarrelIndex)

		generatedSerials := make([]string, 0)
		for partIndex := range parts {
			encoded := b85.Encode(serial.Serialize([]serial.Block{
				{Token: serial_tokenizer.TOK_VARINT, Value: baseBarrel.ManufacturerIndex},
				{Token: serial_tokenizer.TOK_SEP2},
				{Token: serial_tokenizer.TOK_VARINT, Value: 0}, // Unknown, always zero
				{Token: serial_tokenizer.TOK_SEP2},
				{Token: serial_tokenizer.TOK_VARINT, Value: 1}, // Unknown, always one before the level
				{Token: serial_tokenizer.TOK_SEP2},
				{Token: serial_tokenizer.TOK_VARINT, Value: 50}, // Level 50
				{Token: serial_tokenizer.TOK_SEP1},
				{Token: serial_tokenizer.TOK_SEP1},
				{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: baseBarrel.BaseIndex, SubType: part.SUBTYPE_NONE}},
				{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: baseBarrel.BarrelIndex, SubType: part.SUBTYPE_NONE}},
				{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: partIndex, SubType: part.SUBTYPE_NONE}},
				{Token: serial_tokenizer.TOK_SEP1},
			}))
			generatedSerials = append(generatedSerials, encoded)
		}

		_serialsToYaml(generatedSerials)
	}

	fmt.Println("Total combinations:", combinations)
}

func TestAddSamePartALlBases(t *testing.T) {
	generatedSerials := make([]string, 0)
	for baseBarrel, _ := range Barrels {

		encoded := b85.Encode(serial.Serialize([]serial.Block{
			{Token: serial_tokenizer.TOK_VARINT, Value: baseBarrel.ManufacturerIndex},
			{Token: serial_tokenizer.TOK_SEP2},
			{Token: serial_tokenizer.TOK_VARINT, Value: 0}, // Unknown, always zero
			{Token: serial_tokenizer.TOK_SEP2},
			{Token: serial_tokenizer.TOK_VARINT, Value: 1}, // Unknown, always one before the level
			{Token: serial_tokenizer.TOK_SEP2},
			{Token: serial_tokenizer.TOK_VARINT, Value: 50}, // Level 50
			{Token: serial_tokenizer.TOK_SEP1},
			{Token: serial_tokenizer.TOK_SEP1},
			{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: baseBarrel.BaseIndex, SubType: part.SUBTYPE_NONE}},
			{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: baseBarrel.BarrelIndex, SubType: part.SUBTYPE_NONE}},
			{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: 25, SubType: part.SUBTYPE_NONE}},
			{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: 29, SubType: part.SUBTYPE_NONE}},
			{Token: serial_tokenizer.TOK_SEP1},
		}))
		generatedSerials = append(generatedSerials, encoded)
	}

	_serialsToYaml(generatedSerials)
}
