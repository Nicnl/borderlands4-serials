package serial_tokenizer

import (
	"borderlands_4_serials/lib/byte_mirror"
	"fmt"
)

func (t *Tokenizer) readPart() (uint32, byte, uint32, error) {
	index, err := t.readVarint()
	if err != nil {
		return 0, 0, 0, err
	}

	flag, ok := t.bs.ReadN(3)
	if !ok {
		return 0, 0, 0, fmt.Errorf("unexpected end of data while reading part 101 flag")
	}
	//fmt.Printf("flag = %03b", flag)

	switch flag {
	//case 0b010, 0b001, 0b100, 0b000, 0b110, 0b101:
	case 0b010, 0b001:
		return index, byte(flag), 0, nil
	case 0b110:
		param, ok := t.bs.ReadN(6)
		if !ok {
			return 0, 0, 0, fmt.Errorf("unexpected end of data while reading part 101.110 extra 6 bits")
		}
		return index, byte(flag), param, nil
	case 0b101:
		// ...101=110000
		// ...101=01110000000

		// 101:10000.101:01000010100001001000101010001100100010101100110010001010111101001000101010000110000010101110111000001010101001010000101011000101000010101001110100001010101111010000100000

		// 101:01101-11110.101:01110000000  1010001111110.001  01  [10011100 10001100]  00000000
		// 1010110111110101011100000001010001111110001011001000110000100000011000000000000000

		// subtype .101 at 6 bits because it lines up good with these weapon parts:
		//     Item: Accelerated Converging Kickballer (weapon)
		//     Serial: @UgeU_{Fme!KC`?dlRG}I*bm&npQU6dOQDIPZP=8Q;P_a<A5C
		//     101100001011100001011110001010111101001000101011111111000010101000010010001010100001100000101010101110000010101111111000001010101111100000101010001101000010101110110100001000
		//     101:10000.101:110000  101:11100.010  101:11101-00100.010  101:11111-11000.010  101:00001-00100.010  101:00001-10000.010  101:01011-10000.010  101:11111-10000.010  101:01111-10000.010  101:00011-01000.010  101:11011-01000.010  00
		//        => Ends with a "00" separator, so the param is likely 6 bits long + all the other parts align with "101.....010"

		{
			param, ok := t.bs.ReadN(6)
			if !ok {
				return 0, 0, 0, fmt.Errorf("unexpected end of data while reading part 101.100 extra 6 bits")
			}

			fmt.Printf("%06b", param)

			switch param {
			case 0b110000, 0b010000:
				// known good
				param = byte_mirror.GenericMirror(param, 6)
				return index, byte(flag), param, nil
			default:
				t.bs.Rewind(6)
			}
		}

		{
			param, ok := t.bs.ReadN(11)
			if !ok {
				return 0, 0, 0, fmt.Errorf("unexpected end of data while reading part 101.111 extra 11 bits")
			}

			switch param {
			case 0b01110000000:
				param = byte_mirror.GenericMirror(param, 11)
				return index, byte(flag), param, nil
			default:
				t.bs.Rewind(11)
			}
		}

		return 0, 0, 0, fmt.Errorf("unknown part 101 subtype <:%03b> at position %d", flag, t.bs.Pos()-3)

	case 0b100:
		param, ok := t.bs.ReadN(6)
		if !ok {
			return 0, 0, 0, fmt.Errorf("unexpected end of data while reading part 101.100 extra 6 bits")
		}
		param = byte_mirror.GenericMirror(param, 6)
		return index, byte(flag), param, nil
	case 0b111:
		// subtype .111 at 11 bits because it lines up good with these weapon parts:
		//    Item: Rapid Buoy (gadget)
		//    Serial: @Uge8jxm/)}}!eC9HOciQ<Z$Ot?AI&8h00
		//    Bits: 0010000  11010010011010001  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100000110101100100  00  00  10100110010  10101000010  101101011111011110110000  00  01  01  11010010101101011  11100010110011101001001001100000000000
		//    101:10101-11110.11110110000000  101:11010.010  101:10101-11110.001 01 [1001110100100 10011000] 00000000
		//       => Ends with an integer array (100...001 01 [...]) which is

		{
			param, ok := t.bs.ReadN(6)
			if !ok {
				return 0, 0, 0, fmt.Errorf("unexpected end of data while reading part 101.100 extra 6 bits")
			}

			fmt.Printf("%06b", param)

			switch param {
			case 0b110000, 0b010000:
				// known good
				param = byte_mirror.GenericMirror(param, 6)
				return index, byte(flag), param, nil
			default:
				t.bs.Rewind(6)
			}
			return 0, 0, 0, fmt.Errorf("unknown part 111 subtype <:%03b> at position %d", flag, t.bs.Pos()-3)
		}
	default:
		return 0, 0, 0, fmt.Errorf("unknown part 101 flag <:%03b> at position %d", flag, t.bs.Pos()-3)
	}
}
