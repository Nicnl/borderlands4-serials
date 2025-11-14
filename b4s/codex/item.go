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

func Deserialize(base85 string, populateBitstream bool) (*Item, error) {
	data, err := b85.Decode(base85)
	if err != nil {
		return nil, err
	}

	var (
		blocks serial.Serial
		bits   string
	)

	if populateBitstream {
		blocks, bits, err = serial.Deserialize(data)
	} else {
		blocks, _, err = serial.DeserializeWithoutBitstream(data)
	}
	if err != nil {
		return &Item{
			B85:    base85,
			Bits:   bits,
			Serial: blocks,
		}, err
	}

	return &Item{
		B85:    base85,
		Bits:   bits,
		Serial: blocks,
	}, nil
}
