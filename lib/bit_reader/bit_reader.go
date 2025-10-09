package bit_reader

type BitReader struct {
	data []byte
	pos  int
}

func NewBitReader(data []byte) *BitReader {
	return &BitReader{data: data, pos: 0}
}

func (br *BitReader) Read() (byte, bool) {
	if br.pos >= len(br.data)*8 {
		return 0, false
	}
	b := br.data[br.pos/8]
	bit := (b >> (7 - (br.pos % 8))) & 1
	br.pos++
	return bit, true
}

func (br *BitReader) ReadN(n int) (uint32, bool) {
	if n <= 0 || n > 32 {
		return 0, false
	}

	if br.pos+n > len(br.data)*8 {
		return 0, false
	}

	var value uint32 = 0
	for i := 0; i < n; i++ {
		bit, _ := br.Read()
		value = (value << 1) | uint32(bit)
	}
	return value, true
}

func (br *BitReader) Pos() any {
	return br.pos
}

func (br *BitReader) Rewind(n int) bool {
	if n < 0 || br.pos-n < 0 {
		return false
	}
	br.pos -= n
	return true
}

func (br *BitReader) AsString() string {
	result := ""
	for i := br.pos; i < len(br.data)*8; i++ {
		bit, _ := br.Read()
		result += string('0' + bit)
	}
	return result
}
