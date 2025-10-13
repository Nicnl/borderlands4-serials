package serial_tokenizer

import (
	"borderlands_4_serials/lib/serial_tokenizer/part"
	"fmt"
)

func (t *Tokenizer) readPart() (part.Part, error) {
	p := part.Part{
		SubType: part.SUBTYPE_NONE,
	}

	// First, read the index
	{
		index, err := t.readVarint()
		if err != nil {
			return part.Part{}, err
		}
		p.Index = index
	}

	// Next flag partially determines the type of part
	flagType1, ok := t.bs.Read()
	if !ok {
		return part.Part{}, fmt.Errorf("unexpected end of data while reading part flag type 1")
	}

	if flagType1 == 1 {
		p.SubType = part.SUBTYPE_INT

		// Subtype is just a varint
		value, err := t.readVarint()
		if err != nil {
			return part.Part{}, fmt.Errorf("error reading part int value: %w", err)
		}
		p.Value = value

		// Should end with "000"
		err = t.expect("type part, subpart of type int, expect 0x000 as terminator", 0, 0, 0)
		if err != nil {
			return part.Part{}, err
		}

		return p, nil
	}

	// If we are here, we're at 0x0
	// The rest of the decoding depends on the next two bits
	flagType2, ok := t.bs.ReadN(2)
	if !ok {
		return part.Part{}, fmt.Errorf("unexpected end of data while reading part flag type 2")
	}

	//fmt.Printf("flagType2 = %02b\n", flagType2)
	switch flagType2 {
	case 0b10:
		// No data, end of part
		return p, nil
	case 0b01:
		// List of varints
		p.SubType = part.SUBTYPE_LIST

		// Beginning token should be a 01
		token, err := t.nextToken()
		if err != nil {
			return part.Part{}, fmt.Errorf("error reading part list beginning token: %w", err)
		}

		if token != TOK_SEP2 {
			return part.Part{}, fmt.Errorf("expected part list beginning token to be TOK_SEP1 (%d), got %d", TOK_SEP1, token)
		}

		for {
			token, err = t.nextToken()
			if err != nil {
				return part.Part{}, fmt.Errorf("error reading part list item token: %w", err)
			}

			switch token {
			case TOK_SEP1:
				// End of list, there's a 0:00 ahead
				//err = t.expect("type part, subpart of type list, expect 0:00 as terminator", 0, 0)
				//if err != nil {
				//	return Part{}, err
				//}
				return p, nil

			case TOK_VARINT:
				value, err := t.readVarint()
				if err != nil {
					return part.Part{}, fmt.Errorf("error reading part list item value: %w", err)
				}
				p.Values = append(p.Values, value)

			case TOK_VARBIT:
				value, err := t.readVarBit()
				if err != nil {
					return part.Part{}, fmt.Errorf("error reading part list item value (varbit): %w", err)
				}
				p.Values = append(p.Values, value)

			default:
				return part.Part{}, fmt.Errorf("unexpected token %d while reading part list item", token)
			}
		}
	}

	return part.Part{}, fmt.Errorf("ERROR: unknown part flagType2 %02b", flagType2)
}
