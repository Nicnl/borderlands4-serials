package byte_mirror

var (
	Uint4Mirror  [16]byte
	Uint5Mirror  [32]byte
	Uint7Mirror  [128]byte
	Uint8Mirror  [256]byte
	Uint3Mirror  [8]byte
	Uint11Mirror [2048]uint32
	Uint2Mirror  [4]byte
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

	for i := range Uint7Mirror {
		b := byte(i)
		Uint7Mirror[i] = (b&0x01)<<6 | (b&0x02)<<4 | (b&0x04)<<2 | (b & 0x08) |
			(b&0x10)>>2 | (b&0x20)>>4 | (b&0x40)>>6
	}

	for i := range Uint3Mirror {
		b := byte(i)
		Uint3Mirror[i] = (b&0x01)<<2 | (b & 0x02) | (b&0x04)>>2
	}

	for i := range Uint11Mirror {
		b := uint32(i)
		Uint11Mirror[i] = (b&0x001)<<10 | (b&0x002)<<8 | (b&0x004)<<6 | (b&0x008)<<4 | (b&0x010)<<2 |
			(b & 0x020) | (b&0x040)>>2 | (b&0x080)>>4 | (b&0x100)>>6 | (b&0x200)>>8 | (b&0x400)>>10
	}

	for i := range Uint2Mirror {
		b := byte(i)
		Uint2Mirror[i] = (b&0x01)<<1 | (b&0x02)>>1
	}
}

func GenericMirror(input uint32, bitCount int) uint32 {
	var output uint32
	for i := 0; i < bitCount; i++ {
		if (input & (1 << i)) != 0 {
			output |= 1 << (bitCount - 1 - i)
		}
	}
	return output
}
