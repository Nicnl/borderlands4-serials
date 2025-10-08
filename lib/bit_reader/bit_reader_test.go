package bit_reader

import "testing"

func bit(br *BitReader, expected byte, t *testing.T) {
	bit, ok := br.Pop()
	if !ok || bit != expected {
		t.Errorf("Expected %d, got %d", expected, bit)
	}
}

func TestBitReader1(t *testing.T) {
	br := NewBitReader([]byte{0b10101010, 0b11001100})

	bit(br, 1, t)
	bit(br, 0, t)
	bit(br, 1, t)
	bit(br, 0, t)
	bit(br, 1, t)
	bit(br, 0, t)
	bit(br, 1, t)
	bit(br, 0, t)

	bit(br, 1, t)
	bit(br, 1, t)
	bit(br, 0, t)
	bit(br, 0, t)
	bit(br, 1, t)
	bit(br, 1, t)
	bit(br, 0, t)
	bit(br, 0, t)

	// Test out of bounds
	_, ok := br.Pop()
	if ok {
		t.Errorf("Expected false when popping out of bounds")
	}
}

func TestBitReaderPopN4(t *testing.T) {
	br := NewBitReader([]byte{0b10101010, 0b11001100, 0b11110000})

	val, ok := br.PopN(4)
	if !ok || val != 0b1010 {
		t.Errorf("Expected 0b1010, got %08b", val)
	}

	val, ok = br.PopN(4)
	if !ok || val != 0b1010 {
		t.Errorf("Expected 0b1010, got %08b", val)
	}

	val, ok = br.PopN(4)
	if !ok || val != 0b1100 {
		t.Errorf("Expected 0b1100, got %08b", val)
	}

	val, ok = br.PopN(4)
	if !ok || val != 0b1100 {
		t.Errorf("Expected 0b1100, got %08b", val)
	}

	val, ok = br.PopN(4)
	if !ok || val != 0b1111 {
		t.Errorf("Expected 0b1111, got %08b", val)
	}

	val, ok = br.PopN(4)
	if !ok || val != 0b0000 {
		t.Errorf("Expected 0b0000, got %08b", val)
	}

	// Test out of bounds
	_, ok = br.PopN(4)
	if ok {
		t.Errorf("Expected false when popping out of bounds")
	}
}
