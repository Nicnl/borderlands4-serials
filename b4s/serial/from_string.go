package serial

import (
	"borderlands_4_serials/b4s/serial_datatypes/part"
	"borderlands_4_serials/b4s/serial_datatypes/varbit"
	"borderlands_4_serials/b4s/serial_datatypes/varint"
	"borderlands_4_serials/b4s/serial_tokenizer"
	"borderlands_4_serials/lib/bit"
	"fmt"
	"strconv"
	"strings"
)

func isNumbers(str string) (uint32, bool) {
	str = strings.TrimSpace(str)

	for _, r := range str {
		if r < '0' || r > '9' {
			return 0, false
		}
	}

	v, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, false
	}

	return uint32(v), true
}

func isPartSimple(str string) (uint32, bool) {
	str = strings.TrimSpace(str)
	if str[0] != '{' || str[len(str)-1] != '}' {
		return 0, false
	}
	str = strings.TrimSpace(str[1 : len(str)-1])

	// Check if middle is numbers
	return isNumbers(str)
}

func isPartSubtypeInt(str string) (uint32, uint32, bool) {
	str = strings.TrimSpace(str)

	if str[0] != '{' || str[len(str)-1] != '}' {
		return 0, 0, false
	}

	middle := strings.TrimSpace(str[1 : len(str)-1])

	splitted := strings.Split(middle, ":")
	if len(splitted) != 2 {
		return 0, 0, false
	}

	index, ok := isNumbers(splitted[0])
	if !ok {
		return 0, 0, false
	}

	value, ok := isNumbers(splitted[1])
	if !ok {
		return 0, 0, false
	}

	return index, value, true
}

func isPartSubtypeList(str string) (uint32, []uint32, bool) {
	str = strings.TrimSpace(str)

	if str[0] != '{' || str[len(str)-1] != '}' {
		return 0, nil, false
	}

	middle := strings.TrimSpace(str[1 : len(str)-1])
	splitted := strings.Split(middle, ":")
	if len(splitted) < 2 {
		return 0, nil, false
	}

	index, ok := isNumbers(strings.TrimSpace(splitted[0]))
	if !ok {
		return 0, nil, false
	}

	listStr := strings.TrimSpace(splitted[1])
	if listStr[0] != '[' || listStr[len(listStr)-1] != ']' {
		return 0, nil, false
	}
	listStr = strings.TrimSpace(listStr[1 : len(listStr)-1])
	//fmt.Println("listStr =", listStr)

	list := make([]uint32, 0, len(splitted)-1)
	splittedNumbers := strings.Split(listStr, " ")
	for _, numStr := range splittedNumbers {
		numStr = strings.TrimSpace(numStr)

		if numStr == "" {
			continue
		}

		v, ok := isNumbers(numStr)
		if !ok {
			return 0, nil, false
		}
		list = append(list, v)
	}

	return index, list, true
}

func bestTypeForValue(v uint32) serial_tokenizer.Token {
	bw1Varint := bit.NewWriter()
	varint.Write(bw1Varint, v)

	bwVarbit := bit.NewWriter()
	varbit.Write(bwVarbit, v)

	if bw1Varint.Pos() > bwVarbit.Pos() {
		return serial_tokenizer.TOK_VARBIT
	} else {
		return serial_tokenizer.TOK_VARINT
	}
}

func (s *Serial) FromString(str string) error {
	for {
		before := str

		str = strings.ReplaceAll(str, "} ", "}")
		str = strings.ReplaceAll(str, " {", "{")
		str = strings.ReplaceAll(str, ", ", ",")
		str = strings.ReplaceAll(str, "| ", "|")
		str = strings.ReplaceAll(str, " |", "|")

		if before == str {
			break
		}
	}
	// str = "24,0,1,50|2,3379||{76}{2}{3}{75}{57}{60}{59}{16}{25}{44}{49:4}{54:[1 2 3]}|"

	var (
		from   = 0
		to     = 1
		blocks = make([]Block, 0, 50)
	)
	for {
		if to >= len(str) {
			break
		}

		buffer := str[from:to]
		char := str[to]
		//fmt.Printf("buffer='%s'   //   char='%s'\n", buffer, string(char))

		switch char {
		case ',', '|', '{':
			from = to
			if buffer == "" {
				// Nothing to do
			} else if v, ok := isNumbers(buffer); ok {
				blocks = append(blocks, Block{
					Token: bestTypeForValue(v),
					Value: v,
				})
				//fmt.Println("    Add block:", blocks[len(blocks)-1])
			} else if v, ok := isPartSimple(buffer); ok {
				blocks = append(blocks, Block{
					Token: serial_tokenizer.TOK_PART,
					Part: part.Part{
						Index:   v,
						SubType: part.SUBTYPE_NONE,
					},
				})
				//fmt.Println("    Add block:", blocks[len(blocks)-1])
			} else if index, v, ok := isPartSubtypeInt(buffer); ok {
				blocks = append(blocks, Block{
					Token: serial_tokenizer.TOK_PART,
					Part: part.Part{
						Index:   index,
						SubType: part.SUBTYPE_INT,
						Value:   v,
					},
				})
				//fmt.Println("    Add block:", blocks[len(blocks)-1])
			} else if index, list, ok := isPartSubtypeList(buffer); ok {
				blocks = append(blocks, Block{
					Token: serial_tokenizer.TOK_PART,
					Part: part.Part{
						Index:   index,
						SubType: part.SUBTYPE_LIST,
						Values:  list,
					},
				})
				//fmt.Println("    Add block:", blocks[len(blocks)-1])
			} else {
				return fmt.Errorf("invalid integer: '%s'", buffer)
			}
		}

		switch char {
		case '|':
			from++
			blocks = append(blocks, Block{Token: serial_tokenizer.TOK_SEP1})
			//fmt.Println("    Add block:", blocks[len(blocks)-1])
		case ',':
			from++
			blocks = append(blocks, Block{Token: serial_tokenizer.TOK_SEP2})
			//fmt.Println("    Add block:", blocks[len(blocks)-1])
		}

		if char >= '0' && char <= '9' {
			// Just a number, continue
		} else if char == '{' || char == '}' {
			// Just started a part, continue
		} else if char == ' ' {
			// Spaces are okay
		} else if char == ':' || char == '[' || char == ']' || char == ',' || char == '|' {
			// Also okay inside parts
		} else {
			return fmt.Errorf("invalid character: '%s' at pos %d", string(char), to)
		}

		to++
	}

	//fmt.Println("FINAL:", s.String())
	*s = blocks

	return nil
}
