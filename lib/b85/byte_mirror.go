package b85

var mirrorLookup [256]byte

func init() {
	for i := range mirrorLookup {
		b := byte(i)
		mirrorLookup[i] = (b&0x01)<<7 | (b&0x02)<<5 | (b&0x04)<<3 | (b&0x08)<<1 |
			(b&0x10)>>1 | (b&0x20)>>3 | (b&0x40)>>5 | (b&0x80)>>7
	}
}
