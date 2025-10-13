package part

import (
	"borderlands_4_serials/b4s/serial_datatypes/varbit"
	"borderlands_4_serials/b4s/serial_datatypes/varint"
	"borderlands_4_serials/b4s/serial_tokenizer"
	"fmt"
)

func Read(t *serial_tokenizer.Tokenizer) (Part, error) {
	br := t.BitReader()

	p := Part{
		SubType: SUBTYPE_NONE,
	}

	// First, read the index
	{
		index, err := varint.Read(br)
		if err != nil {
			return Part{}, err
		}
		p.Index = index
	}

	// Next flag partially determines the type of part
	flagType1, ok := br.Read()
	if !ok {
		return Part{}, fmt.Errorf("unexpected end of data while reading part flag type 1")
	}

	if flagType1 == 1 {
		p.SubType = SUBTYPE_INT

		// Subtype is just a varint
		value, err := varint.Read(br)
		if err != nil {
			return Part{}, fmt.Errorf("error reading part int value: %w", err)
		}
		p.Value = value

		// Should end with "000"
		err = t.Expect("type part, subpart of type int, expect 0x000 as terminator", 0, 0, 0)
		if err != nil {
			return Part{}, err
		}

		return p, nil
	}

	// If we are here, we're at 0x0
	// The rest of the decoding depends on the next two bits
	flagType2, ok := br.ReadN(2)
	if !ok {
		return Part{}, fmt.Errorf("unexpected end of data while reading part flag type 2")
	}

	//fmt.Printf("flagType2 = %02b\n", flagType2)
	switch flagType2 {
	case 0b10:
		// No data, end of part
		return p, nil
	case 0b01:
		// List of varints
		p.SubType = SUBTYPE_LIST

		// Beginning token should be a 01
		token, err := t.NextToken()
		if err != nil {
			return Part{}, fmt.Errorf("error reading part list beginning token: %w", err)
		}

		if token != serial_tokenizer.TOK_SEP2 {
			return Part{}, fmt.Errorf("expected part list beginning token to be TOK_SEP2 (%d), got %d", serial_tokenizer.TOK_SEP2, token)
		}

		for {
			token, err = t.NextToken()
			if err != nil {
				return Part{}, fmt.Errorf("error reading part list item token: %w", err)
			}

			switch token {
			case serial_tokenizer.TOK_SEP1:
				// End of list, there's a 0:00 ahead
				//err = t.expect("type part, subpart of type list, expect 0:00 as terminator", 0, 0)
				//if err != nil {
				//	return Part{}, err
				//}
				return p, nil

			case serial_tokenizer.TOK_VARINT:
				value, err := varint.Read(br)
				if err != nil {
					return Part{}, fmt.Errorf("error reading part list item value: %w", err)
				}
				p.Values = append(p.Values, value)

			case serial_tokenizer.TOK_VARBIT:
				value, err := varbit.Read(br)
				if err != nil {
					return Part{}, fmt.Errorf("error reading part list item value (varbit): %w", err)
				}
				p.Values = append(p.Values, value)

			default:
				return Part{}, fmt.Errorf("unexpected token %d while reading part list item", token)
			}
		}
	}

	return Part{}, fmt.Errorf("ERROR: unknown part flagType2 %02b", flagType2)
}
