package bit

type Writer struct {
	data []byte
	pos  int
}

func NewWriter() *Writer {
	return &Writer{data: make([]byte, 0, 250), pos: 0}
}

func (bw *Writer) WriteBit(bit byte) {
	if bw.pos/8 >= len(bw.data) {
		bw.data = append(bw.data, 0)
	}
	if bit&1 == 1 {
		bw.data[bw.pos/8] |= 1 << (7 - (bw.pos % 8))
	}

	bw.pos++
}

func (bw *Writer) WriteBits(bits ...byte) {
	for _, bit := range bits {
		bw.WriteBit(bit)
	}
}

func (bw *Writer) WriteN(value uint32, n int) {
	for i := n - 1; i >= 0; i-- {
		bit := (value >> i) & 1
		bw.WriteBit(byte(bit))
	}
}
