package codex

import (
	"borderlands_4_serials/b4s/serial_datatypes/part"
	"borderlands_4_serials/b4s/serial_tokenizer"
)

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
						SubType: part.SUBTYPE_LIST,
						Values:  []uint32{value},
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
