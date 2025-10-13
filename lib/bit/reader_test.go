package bit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type _testSizedInt struct {
	value uint32
	size  int
}

var testsBits = []struct {
	name  string
	bytes []byte
	bits  []byte
}{
	{
		name:  "1",
		bytes: []byte{0b00000001},
		bits: []byte{
			0, 0, 0, 0, 0, 0, 0, 1,
		},
	},
	{
		name:  "128",
		bytes: []byte{0b10000000},
		bits: []byte{
			1, 0, 0, 0, 0, 0, 0, 0,
		},
	},
	{
		name:  "multiblocks",
		bytes: []byte{0b10101010, 0b11001100, 0b11110000},
		bits: []byte{
			1, 0, 1, 0, 1, 0, 1, 0,
			1, 1, 0, 0, 1, 1, 0, 0,
			1, 1, 1, 1, 0, 0, 0, 0,
		},
	},
}

var testsInts = []struct {
	name  string
	bytes []byte
	ints  []_testSizedInt
}{
	{
		name:  "5x4-bits",
		bytes: []byte{0b10101010, 0b11001100, 0b11110000},
		ints: []_testSizedInt{
			{0b1010, 4},
			{0b1010, 4},
			{0b1100, 4},
			{0b1100, 4},
			{0b1111, 4},
		},
	},
	{
		name:  "6x4-bits",
		bytes: []byte{0b10101010, 0b11001100, 0b11110001},
		ints: []_testSizedInt{
			{0b1010, 4},
			{0b1010, 4},
			{0b1100, 4},
			{0b1100, 4},
			{0b1111, 4},
			{0b0001, 4},
		},
	},
	{
		name:  "4x6-bits (overlap)",
		bytes: []byte{0b11000101, 0b01100110, 0b01111000},
		ints: []_testSizedInt{
			{0b110001, 6},
			{0b010110, 6},
			{0b011001, 6},
			{0b111000, 6},
		},
	},
}

func bit(br *Reader, expected byte, t *testing.T) {
	bit, ok := br.Read()
	if !ok || bit != expected {
		t.Errorf("Expected %d, got %d", expected, bit)
	}
}

func TestBitReaderReadBits(t *testing.T) {
	for _, test := range testsBits {
		t.Run(test.name, func(t *testing.T) {
			br := NewReader(test.bytes)
			fmt.Printf("Data:")
			for _, b := range test.bytes {
				fmt.Printf(" %08b", b)
			}
			fmt.Println()

			for _, expected := range test.bits {
				fmt.Println("  Expect", expected)
				bit(br, expected, t)
			}
		})
	}
}

func TestBitReaderReadInts(t *testing.T) {
	for _, test := range testsInts {
		t.Run(test.name, func(t *testing.T) {
			br := NewReader(test.bytes)
			for _, block := range test.ints {
				fmt.Println("Reading: ", block.size, "bits")
				value, ok := br.ReadN(block.size)
				fmt.Println("Got:     ", value, ok)
				fmt.Println("Expected:", block.value)
				assert.True(t, ok)
				assert.Equal(t, block.value, value)
				fmt.Println()
			}
		})
	}
}
