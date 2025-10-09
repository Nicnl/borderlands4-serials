package b85

import (
	"borderlands_4_serials/lib/byte_mirror"
	"fmt"
)

const b85Charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{/}~"

var reverseLookup [256]byte

func init() {
	// Initialize reverse lookup table
	for i := range reverseLookup {
		reverseLookup[i] = 0xFF
	}
	for i := 0; i < len(b85Charset); i++ {
		reverseLookup[b85Charset[i]] = byte(i)
	}
}

func Decode(serial string) ([]byte, error) {
	if len(serial) < 2 || serial[0] != '@' || serial[1] != 'U' {
		return nil, fmt.Errorf("not a valid borderlands 4 item serial")
	}
	serial = serial[2:]

	result := make([]byte, 0)
	idx := 0
	size := len(serial)

	for idx < size {
		var workingU32 uint32 = 0
		charCount := 0

		// Collect up to 5 valid Base85 characters
		for idx < size && charCount < 5 {
			charCode := serial[idx]
			idx++

			if reverseLookup[charCode] < 0x56 {
				workingU32 = workingU32*85 + uint32(reverseLookup[charCode])
				charCount++
			}
		}

		if charCount == 0 {
			break
		}

		// Handle padding for incomplete groups
		if charCount < 5 {
			padding := 5 - charCount
			for i := 0; i < padding; i++ {
				workingU32 = workingU32*85 + 0x7e // '~' value
			}
		}

		// Extract bytes - same for both full and partial groups
		byteCount := 4
		if charCount < 5 {
			byteCount = charCount - 1
		}

		if byteCount >= 1 {
			result = append(result, byte((workingU32>>24)&0xFF))
		}
		if byteCount >= 2 {
			result = append(result, byte((workingU32>>16)&0xFF))
		}
		if byteCount >= 3 {
			result = append(result, byte((workingU32>>8)&0xFF))
		}
		if byteCount >= 4 {
			result = append(result, byte((workingU32>>0)&0xFF))
		}
	}

	// Reverse the bits in each byte
	// 76543210 -> 01234567
	for i := range result {
		// Using a lookup table for performance
		result[i] = byte_mirror.Uint8Mirror[result[i]]
	}

	return result, nil
}
