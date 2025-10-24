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
	if len(listStr) == 0 {
		return 0, nil, false
	}
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
	var blocks []Block
	i := 0

	for i < len(str) {
		char := str[i]

		switch {
		case char == ' ' || char == '\t' || char == '\n' || char == '\r':
			// Skip whitespace
			i++
		case char == '|':
			blocks = append(blocks, Block{Token: serial_tokenizer.TOK_SEP1})
			i++
		case char == ',':
			blocks = append(blocks, Block{Token: serial_tokenizer.TOK_SEP2})
			i++
		case char == '{':
			// Find the matching '}'
			end := strings.IndexRune(str[i:], '}')
			if end == -1 {
				return fmt.Errorf("unmatched '{' at position %d", i)
			}
			end += i // Adjust index to be relative to the start of the string

			partStr := str[i : end+1]
			i = end + 1

			if index, list, ok := isPartSubtypeList(partStr); ok {
				blocks = append(blocks, Block{
					Token: serial_tokenizer.TOK_PART,
					Part: part.Part{
						Index:   index,
						SubType: part.SUBTYPE_LIST,
						Values:  list,
					},
				})
			} else if index, v, ok := isPartSubtypeInt(partStr); ok {
				blocks = append(blocks, Block{
					Token: serial_tokenizer.TOK_PART,
					Part: part.Part{
						Index:   index,
						SubType: part.SUBTYPE_INT,
						Value:   v,
					},
				})
			} else if v, ok := isPartSimple(partStr); ok {
				blocks = append(blocks, Block{
					Token: serial_tokenizer.TOK_PART,
					Part: part.Part{
						Index:   v,
						SubType: part.SUBTYPE_NONE,
					},
				})
			} else {
				return fmt.Errorf("invalid part format: '%s'", partStr)
			}
		case char >= '0' && char <= '9':
			// Find the end of the number
			start := i
			for i < len(str) && str[i] >= '0' && str[i] <= '9' {
				i++
			}
			numStr := str[start:i]

			v, err := strconv.ParseUint(numStr, 10, 32)
			if err != nil {
				// This should not happen given the loop condition
				return fmt.Errorf("invalid number: '%s'", numStr)
			}
			blocks = append(blocks, Block{
				Token: bestTypeForValue(uint32(v)),
				Value: uint32(v),
			})

		case char == '"':
			// Find the closing quote
			end := i + 1
			for end < len(str) && str[end] != '"' {
				// Handle escaped quotes
				if str[end] == '\\' && end+1 < len(str) && str[end+1] == '"' {
					end += 2
					continue
				}
				end++
			}
			if end >= len(str) {
				return fmt.Errorf("unmatched '\"' at position %d", i)
			}
			strContent := str[i+1 : end]
			i = end + 1

			// Unescape quotes and backslashes
			strContent = strings.ReplaceAll(strContent, "\\\"", "\"")
			strContent = strings.ReplaceAll(strContent, "\\\\", "\\")

			blocks = append(blocks, Block{
				Token:    serial_tokenizer.TOK_STRING,
				ValueStr: strContent,
			})
		default:
			return fmt.Errorf("invalid character: '%c' at position %d", char, i)
		}
	}

	*s = blocks
	return nil
}
