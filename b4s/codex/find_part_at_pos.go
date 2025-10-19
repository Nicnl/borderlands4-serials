package codex

import (
	"borderlands_4_serials/b4s/serial_datatypes/part"
	"borderlands_4_serials/b4s/serial_tokenizer"
)

func (i *Item) HasPart(p part.Part) bool {
	for _, block := range i.Serial {
		if block.Token != serial_tokenizer.TOK_PART {
			continue
		}

		if block.Part.Index != p.Index {
			continue
		}

		if block.Part.SubType != p.SubType {
			continue
		}

		switch block.Part.SubType {
		case part.SUBTYPE_NONE:
			return true
		case part.SUBTYPE_INT:
			return block.Part.Value == p.Value
		case part.SUBTYPE_LIST:
			if len(block.Part.Values) != len(p.Values) {
				continue
			}
			matchedAll := true
			for idx, val := range block.Part.Values {
				if val != p.Values[idx] {
					matchedAll = false
					break
				}
			}
			if matchedAll {
				return true
			}
		}
	}

	return false
}

func (i *Item) FindPartAtPos(pos int, splitLists bool) *part.Part {
	for _, block := range i.Serial {
		if block.Token != serial_tokenizer.TOK_PART {
			continue
		}

		switch block.Part.SubType {
		case part.SUBTYPE_NONE, part.SUBTYPE_INT:
			if pos == 0 {
				return &block.Part
			} else {
				pos -= 1
			}
		case part.SUBTYPE_LIST:
			if !splitLists {
				if pos == 0 {
					return &block.Part
				} else {
					pos -= 1
				}
			} else {
				for _, value := range block.Part.Values {
					subPart := part.Part{
						Index:   block.Part.Index,
						SubType: part.SUBTYPE_INT,
						Value:   value,
					}
					if pos == 0 {
						return &subPart
					} else {
						pos -= 1
					}
				}
			}
		}
	}

	return nil
}
