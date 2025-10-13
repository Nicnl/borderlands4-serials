package part

import (
	"borderlands_4_serials/b4s/serial_datatypes/varint"
	"borderlands_4_serials/lib/bit"
)

func Write(br *bit.Writer, p Part) {
	// First, write the index
	varint.Write(br, p.Index)

	// Next flag (and data) depends on the type of the part
	switch p.SubType {
	case SUBTYPE_NONE:
		// No data, write 0:10
		br.WriteBits(0, 1, 0)
	case SUBTYPE_INT:
		// 1 indicates "next is varint"
		br.WriteBit(1)

		// Write the varint
		varint.Write(br, p.Value)

		// End with 000
		br.WriteBits(0, 0, 0)
	case SUBTYPE_LIST:
		// 0:01 indicates "next is a list"
		br.WriteBits(0, 0, 1)

		// List starts with a soft separator (0x01)
		br.WriteBits(0, 1)

		// Write all varints with their type
		for _, v := range p.Values {
			// TODO: auto select varint or varbit depending on bitsize
			// For the time being, varints are hardcoded

			// Major type of varint is 100
			br.WriteBits(1, 0, 0)

			// Write the varint
			varint.Write(br, v)
		}

		// End the list with a hard separator (0x00)
		br.WriteBits(0, 0)
	}
}
