package serial

import (
	"borderlands_4_serials/b4s/serial_datatypes/part"
	"borderlands_4_serials/b4s/serial_datatypes/varbit"
	"borderlands_4_serials/b4s/serial_datatypes/varint"
	"borderlands_4_serials/b4s/serial_tokenizer"
	"borderlands_4_serials/lib/bit"
)

func Serialize(s Serial) []byte {
	bw := bit.NewWriter()

	// Write magic header
	bw.WriteBits(0, 0, 1, 0, 0, 0, 0)

	for _, block := range s.Blocks {
		// Write the token (aka: major type)
		switch block.Token {
		case serial_tokenizer.TOK_SEP1:
			bw.WriteBits(0, 0)
		case serial_tokenizer.TOK_SEP2:
			bw.WriteBits(0, 1)
		case serial_tokenizer.TOK_VARINT:
			bw.WriteBits(1, 0, 0)
			varint.Write(bw, block.Value)
		case serial_tokenizer.TOK_VARBIT:
			bw.WriteBits(1, 1, 0)
			varbit.Write(bw, block.Value)
		case serial_tokenizer.TOK_PART:
			bw.WriteBits(1, 0, 1)
			part.Write(bw, block.Part)
		case serial_tokenizer.TOK_PART_111:
			// Unsupported, unknown
			// bw.WriteBits(1, 1, 1)
			// We do not handle this case
			break
		}
	}

	return bw.Data()
}
