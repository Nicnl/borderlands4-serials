package serial

import (
	"borderlands_4_serials/b4s/serial_datatypes/part"
	"borderlands_4_serials/b4s/serial_tokenizer"
	"crypto/sha1"
	"fmt"
	"strings"
)

type Block struct {
	Token    serial_tokenizer.Token
	Value    uint32
	ValueStr string
	Part     part.Part
}

type Serial []Block

func (s Serial) String() string {
	// TODO: we may have to sort the PART blocks by their index to get a repeatable output
	// (don't forget the subtype list too)
	// + add a test for that

	output := ""

	for i, b := range s {

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
		case serial_tokenizer.TOK_PART:
			if i > 0 {
				output += " "
			}
			output += b.Part.String()
		case serial_tokenizer.TOK_STRING:
			if i > 0 {
				output += " "
			}

			b.ValueStr = strings.ReplaceAll(b.ValueStr, "\\", "\\\\")
			b.ValueStr = strings.ReplaceAll(b.ValueStr, "\"", "\\\"")

			output += "\"" + b.ValueStr + "\""
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
