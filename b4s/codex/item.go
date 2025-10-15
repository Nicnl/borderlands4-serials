package codex

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/serial"
)

type Item struct {
	B85    string
	Bits   string
	Serial serial.Serial
}

func Deserialize(base85 string) (*Item, error) {
	data, err := b85.Decode(base85)
	if err != nil {
		return nil, err
	}

	blocks, bits, err := serial.Deserialize(data)
	if err != nil {
		return nil, err
	}

	return &Item{
		B85:    base85,
		Bits:   bits,
		Serial: blocks,
	}, nil
}
