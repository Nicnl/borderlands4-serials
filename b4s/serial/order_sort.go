package serial

import (
	"borderlands_4_serials/b4s/serial_datatypes/part"
	"borderlands_4_serials/b4s/serial_tokenizer"
	"fmt"
	"slices"
)

func (s Serial) SplitBlocks() ([]Block, []Block, error) {
	// We want to split the blocks in two slices:
	// - All the blocks before the first 101 part block (included)
	// - All the parts after the one

	blocks := make([]Block, 0, len(s))
	parts := make([]Block, 0, len(s))

	foundPartBlock := false
	for _, b := range s {
		if foundPartBlock {

			if b.Token == serial_tokenizer.TOK_PART {
				parts = append(parts, b)
			} else if b.Token == serial_tokenizer.TOK_SEP1 {
				// End of parts
				break
			} else {
				return nil, nil, fmt.Errorf("unexpected token after parts: %v", b.Token)
			}

		} else {

			if b.Token != serial_tokenizer.TOK_PART {
				blocks = append(blocks, b)
			} else {
				blocks = append(blocks, b)
				foundPartBlock = true
			}

		}
	}

	return blocks, parts, nil
}

func (s *Serial) Sort() error {
	// We want to build a a new list of blocks, but with the parts sorted
	blocks, parts, err := s.SplitBlocks()
	if err != nil {
		return err
	}

	slices.SortFunc(parts, func(a, b Block) int {
		// Sort by Index, then by SubType, then by Value, then by Values (two different variables
		if a.Part.Index != b.Part.Index {
			return int(a.Part.Index) - int(b.Part.Index)
		}
		if a.Part.SubType != b.Part.SubType {
			return int(a.Part.SubType) - int(b.Part.SubType)
		}
		if a.Part.SubType == part.SUBTYPE_NONE && b.Part.SubType == part.SUBTYPE_NONE {
			return 0
		}
		if a.Part.SubType == part.SUBTYPE_INT && b.Part.SubType == part.SUBTYPE_INT {
			return int(a.Part.Value) - int(b.Part.Value)
		}
		if a.Part.SubType == part.SUBTYPE_LIST && b.Part.SubType == part.SUBTYPE_LIST {
			minLen := len(a.Part.Values)
			if len(b.Part.Values) < minLen {
				minLen = len(b.Part.Values)
			}
			for i := 0; i < minLen; i++ {
				if a.Part.Values[i] != b.Part.Values[i] {
					return int(a.Part.Values[i]) - int(b.Part.Values[i])
				}
			}
			return len(a.Part.Values) - len(b.Part.Values)
		}

		// Should not reach here, there are no other subtypes
		// But just in case, fallback to SubType comparison
		return int(a.Part.SubType) - int(b.Part.SubType)
	})

	blocks = append(blocks, parts...)
	blocks = append(blocks, Block{Token: serial_tokenizer.TOK_SEP1})
	*s = blocks
	return nil
}
