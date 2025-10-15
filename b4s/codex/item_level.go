package codex

import (
	"borderlands_4_serials/b4s/serial_tokenizer"
)

func (i *Item) FindIntAtPos(pos int) (uint32, bool) {
	for _, b := range i.Serial {
		if b.Token == serial_tokenizer.TOK_VARINT || b.Token == serial_tokenizer.TOK_VARBIT {
			if pos == 0 {
				return b.Value, true
			}

			pos--
		}
	}

	return 0, false
}

// FindLevel finds the level in the serial.
// The level is stored as a pair of varints: (marker, level).
// Level seems to be after the marker "1".
func (i *Item) FindLevel() (uint32, bool) {
	pos := 2
	for {
		value, found := i.FindIntAtPos(pos)
		if !found {
			return 0, false
		}

		if value == 1 {
			pos++
			level, found := i.FindIntAtPos(pos)
			if found {
				return level, true
			} else {
				return 0, false
			}
		}

		pos += 2
	}
}
