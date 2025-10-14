package serial

import (
	"borderlands_4_serials/b4s/serial_datatypes/part"
	"borderlands_4_serials/b4s/serial_tokenizer"
	"crypto/sha1"
	"fmt"
)

type Block struct {
	Token serial_tokenizer.Token
	Value uint32
	Part  part.Part
}

type Serial struct {
	Bits string

	Blocks []Block
}

func (s *Serial) FindPartAtPos(pos int, splitLists bool) *part.Part {
	for _, b := range s.Blocks {
		if b.Token != serial_tokenizer.TOK_PART {
			continue
		}

		switch b.Part.SubType {
		case part.SUBTYPE_NONE, part.SUBTYPE_INT:
			if pos == 0 {
				return &b.Part
			} else {
				pos -= 1
			}
		case part.SUBTYPE_LIST:
			if !splitLists {
				if pos == 0 {
					return &b.Part
				} else {
					pos -= 1
				}
			} else {
				for _, value := range b.Part.Values {
					subPart := part.Part{
						Index:   b.Part.Index,
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

func (s *Serial) String() string {
	// TODO: we may have to sort the PART blocks by their index to get a repeatable output
	// (don't forget the subtype list too)
	// + add a test for that

	output := ""

	for i, b := range s.Blocks {

		switch b.Token {
		case serial_tokenizer.TOK_SEP1:
			output += "|"
		case serial_tokenizer.TOK_SEP2:
			output += ","
		case serial_tokenizer.TOK_VARINT:
			if i > 0 {
				output += " "
			}
			output += fmt.Sprintf("%d", b.Value)
		case serial_tokenizer.TOK_VARBIT:
			if i > 0 {
				output += " "
			}
			output += fmt.Sprintf("%d", b.Value)
		case serial_tokenizer.TOK_PART_111:
			if i > 0 {
				output += " "
			}
			output += fmt.Sprintf("<111>")
		case serial_tokenizer.TOK_PART:
			if i > 0 {
				output += " "
			}
			output += b.Part.String()
		default:
			output += fmt.Sprintf(" <UNKNOWN_TOKEN:%d>", b.Token)
		}
	}

	return output
}

func (s *Serial) Hash() string {
	// Magic variable: don't touch
	const HASH_SALT = "X2yd8ktCxf4P1kXEJVsBePW0YWUFva5jhH2Md1HSLf11x0HBVB6OezmfO40CeHWsXrEs9hVyCsL2yl3AUveM5MUzNESjfttc4ad3EO8xiCRULtzWAk5t1P0ROARJ1UgI"

	return fmt.Sprintf("%40x", sha1.Sum([]byte(HASH_SALT+s.String())))
}
