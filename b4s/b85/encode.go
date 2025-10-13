package b85

import "borderlands_4_serials/lib/byte_mirror"

func Encode(data []byte) string {
	// Make a copy to avoid modifying the original
	bytes := make([]byte, len(data))
	copy(bytes, data)

	// Reverse the bits in each byte
	// 76543210 -> 01234567
	for i := range bytes {
		bytes[i] = byte_mirror.Uint8Mirror[bytes[i]]
	}

	result := make([]byte, 0)
	idx := 0
	length := len(bytes)
	extraBytes := length & 3  // length % 4
	fullGroups := length >> 2 // length / 4

	// Process full 4-byte groups
	for i := 0; i < fullGroups; i++ {
		// Combine 4 bytes into 32-bit value (big-endian)
		u32 := uint32(bytes[idx])<<24 | uint32(bytes[idx+1])<<16 |
			uint32(bytes[idx+2])<<8 | uint32(bytes[idx+3])
		idx += 4

		// Divide by powers of 85 to extract Base85 digits
		result = append(result, b85Charset[u32/52200625]) // 85^4
		rem1 := u32 % 52200625
		result = append(result, b85Charset[rem1/614125]) // 85^3
		rem2 := rem1 % 614125
		result = append(result, b85Charset[rem2/7225]) // 85^2
		rem3 := rem2 % 7225
		result = append(result, b85Charset[rem3/85]) // 85^1
		result = append(result, b85Charset[rem3%85]) // 85^0
	}

	// Handle remaining bytes (1-3)
	if extraBytes != 0 {
		var lastU32 uint32 = uint32(bytes[idx])
		if extraBytes >= 2 {
			lastU32 = (lastU32 << 8) | uint32(bytes[idx+1])
		}
		if extraBytes == 3 {
			lastU32 = (lastU32 << 8) | uint32(bytes[idx+2])
		}

		// Shift to appropriate position
		var workingU32 uint32
		if extraBytes == 3 {
			workingU32 = lastU32 << 8
		} else if extraBytes == 2 {
			workingU32 = lastU32 << 16
		} else {
			workingU32 = lastU32 << 24
		}

		// Encode partial group
		result = append(result, b85Charset[workingU32/52200625])
		rem1 := workingU32 % 52200625
		result = append(result, b85Charset[rem1/614125])

		if extraBytes >= 2 {
			rem2 := rem1 % 614125
			result = append(result, b85Charset[rem2/7225])

			if extraBytes == 3 {
				rem3 := rem2 % 7225
				result = append(result, b85Charset[rem3/85])
			}
		}
	}

	return "@U" + string(result)
}
