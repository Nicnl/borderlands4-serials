package tools

import "borderlands_4_serials/lib/b85"

func ShortenSerial(serial string) []string {
	data, err := b85.Decode(serial)
	if err != nil {
		panic(err)
	}

	output := make([]string, 0)
	output = append(output, serial)

	for len(data) > 0 {
		lastByte := data[len(data)-1]

		// Changes bits from the end to the start to zero
		for i := 7; i >= 0; i-- {
			if (lastByte>>i)&1 == 1 {
				lastByte &^= (1 << i)
				data[len(data)-1] = lastByte
				output = append(output, b85.Encode(data))
			}
		}

		// Remove trailing zero bytes
		data = data[0 : len(data)-1]
		output = append(output, b85.Encode(data))
	}

	return output
}
