package bit

type Reader struct {
	data []byte
	pos  int
}

func NewReader(data []byte) *Reader {
	return &Reader{data: data, pos: 0}
}

func (br *Reader) Read() (byte, bool) {
	if br.pos >= len(br.data)*8 {
		return 0, false
	}
	b := br.data[br.pos/8]
	bit := (b >> (7 - (br.pos % 8))) & 1
	br.pos++
	return bit, true
}

func (br *Reader) Read2() (byte, byte, bool) {
	bit1, ok1 := br.Read()
	if !ok1 {
		return 0, 0, false
	}
	bit2, ok2 := br.Read()
	if !ok2 {
		return 0, 0, false
	}
	return bit1, bit2, true
}

func (br *Reader) ReadN(n int) (uint32, bool) {
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

func (br *Reader) Pos() int {
	return br.pos
}

func (br *Reader) SetPos(n int) bool {
	if n < 0 || n > len(br.data)*8 {
		return false
	}
	br.pos = n
	return true
}

func (br *Reader) Rewind(n int) bool {
	if n < 0 || br.pos-n < 0 {
		return false
	}
	br.pos -= n
	return true
}

func (br *Reader) StringBefore() string {
	//As binary
	oldPos := br.pos
	br.Rewind(oldPos)
	result := ""
	for i := 0; i < oldPos; i++ {
		bit, _ := br.Read()
		result += string('0' + bit)
	}
	br.pos = oldPos
	return result
}

func (br *Reader) StringAfter() string {
	result := ""
	for i := br.pos; i < len(br.data)*8; i++ {
		bit, _ := br.Read()
		result += string('0' + bit)
	}
	return result
}

func (br *Reader) FullString() string {
	result := make([]byte, len(br.data)*8, len(br.data)*8)

	oldPos := br.pos
	br.Rewind(oldPos)
	for i := 0; i < len(br.data)*8; i++ {
		bit, _ := br.Read()
		result[i] = '0' + bit
	}
	br.pos = oldPos
	return string(result)
}

func (br *Reader) Len() int {
	return len(br.data) * 8
}
