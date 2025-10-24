package b4string

import (
	"borderlands_4_serials/b4s/serial_datatypes/varint"
	"borderlands_4_serials/lib/bit"
	"borderlands_4_serials/lib/byte_mirror"
)

func Write(bw *bit.Writer, str string) {
	varint.Write(bw, uint32(len(str)))

	for i := 0; i < len(str); i++ {
		bw.WriteN(uint32(byte_mirror.Uint7Mirror[str[i]]), 7)
	}
}
