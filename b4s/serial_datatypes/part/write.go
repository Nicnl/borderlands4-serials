package part

import (
	"borderlands_4_serials/b4s/serial_datatypes/varbit"
	"borderlands_4_serials/b4s/serial_datatypes/varint"
	"borderlands_4_serials/lib/bit"
)

func bestTypeForValue(v uint32) ([]byte, []byte) {
	bwVarint := bit.NewWriter()
	varint.Write(bwVarint, v)

	bwVarbit := bit.NewWriter()
	varbit.Write(bwVarbit, v)

	if bwVarint.Pos() > bwVarbit.Pos() {
		return []byte{1, 1, 0}, bwVarbit.Bits()
	} else {
		return []byte{1, 0, 0}, bwVarint.Bits()
	}
}

func Write(bw *bit.Writer, p Part) {
	// First, write the index
	varint.Write(bw, p.Index)

	// Next flag (and data) depends on the type of the part
	switch p.SubType {
	case SUBTYPE_NONE:
		// No data, write 0:10
		bw.WriteBits(0, 1, 0)
	case SUBTYPE_INT:
		// 1 indicates "next is varint"
		bw.WriteBit(1)

		// Write the varint
		varint.Write(bw, p.Value)

		// End with 000
		bw.WriteBits(0, 0, 0)
	case SUBTYPE_LIST:
		// 0:01 indicates "next is a list"
		bw.WriteBits(0, 0, 1)

		// List starts with a soft separator (0x01)
		bw.WriteBits(0, 1)

		// Write all varints with their type
		for _, v := range p.Values {
			typeBits, valueBits := bestTypeForValue(v)
			bw.WriteBits(typeBits...)
			bw.WriteBits(valueBits...)
		}

		// End the list with a hard separator (0x00)
		bw.WriteBits(0, 0)
	}
}
