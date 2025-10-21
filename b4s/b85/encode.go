package b85

import "borderlands_4_serials/lib/byte_mirror"

// Powers of 85
const (
	_85_1 = 85
	_85_2 = 85 * 85
	_85_3 = 85 * 85 * 85
	_85_4 = 85 * 85 * 85 * 85
)

func Encode(data []byte) string {
	// Make a copy to avoid modifying the original
	bytes := make([]byte, len(data), len(data))
	copy(bytes, data)

	// Mirror the bits in each byte
	// 76543210 -> 01234567
	for i := range bytes {
		// Using a lookup table for performance
		bytes[i] = byte_mirror.Uint8Mirror[bytes[i]]
	}

	var (
		result     = make([]byte, 0, (len(bytes)*5/4)+5) // Preallocate with some extra space
		idx        = 0
		length     = len(bytes)
		extraBytes = length % 4
		fullGroups = length / 4
	)

	// Process full 4-byte groups
	for range fullGroups {
		// Combine 4 bytes into 32-bit value (big-endian)
		v := uint32(bytes[idx])<<24 | uint32(bytes[idx+1])<<16 |
			uint32(bytes[idx+2])<<8 | uint32(bytes[idx+3])
		idx += 4

		// Divide by powers of 85 to extract Base85 digits

		// 85^4
		result = append(result, b85Charset[v/_85_4])
		v = v % _85_4

		// 85^3
		result = append(result, b85Charset[v/_85_3])
		v = v % _85_3

		// 85^2
		result = append(result, b85Charset[v/_85_2])
		v = v % _85_2

		// 85^1
		result = append(result, b85Charset[v/_85_1])

		// 85^0
		result = append(result, b85Charset[v%_85_1])
	}

	// Handle remaining bytes (1-3)
	if extraBytes != 0 {
		// Combine remaining bytes into 32-bit value
		v := uint32(bytes[idx])
		if extraBytes >= 2 {
			v = (v << 8) | uint32(bytes[idx+1])
		}
		if extraBytes == 3 {
			v = (v << 8) | uint32(bytes[idx+2])
		}

		// Shift to appropriate position
		if extraBytes == 3 {
			v = v << 8
		} else if extraBytes == 2 {
			v = v << 16
		} else {
			v = v << 24
		}

		// Encode partial group
		result = append(result, b85Charset[v/_85_4])
		v = v % _85_4
		result = append(result, b85Charset[v/_85_3])

		if extraBytes >= 2 {
			v = v % _85_3
			result = append(result, b85Charset[v/_85_2])

			if extraBytes == 3 {
				v = v % _85_2
				result = append(result, b85Charset[v/_85_1])
			}
		}
	}

	return "@U" + string(result)
}
