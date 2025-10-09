package byte_mirror

var (
	Uint8Mirror [256]byte
	Uint4Mirror [16]byte
	Uint5Mirror [32]byte
)

func init() {
	for i := range Uint8Mirror {
		b := byte(i)
		Uint8Mirror[i] = (b&0x01)<<7 | (b&0x02)<<5 | (b&0x04)<<3 | (b&0x08)<<1 |
			(b&0x10)>>1 | (b&0x20)>>3 | (b&0x40)>>5 | (b&0x80)>>7
	}

	for i := range Uint4Mirror {
		b := byte(i)
		Uint4Mirror[i] = (b&0x01)<<3 | (b&0x02)<<1 | (b&0x04)>>1 | (b&0x08)>>3
	}

	for i := range Uint5Mirror {
		b := byte(i)
		Uint5Mirror[i] = (b&0x01)<<4 | (b&0x02)<<2 | (b & 0x04) | (b&0x08)>>2 | (b&0x10)>>4
	}
}
