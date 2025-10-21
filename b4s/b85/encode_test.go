package b85

import (
	"encoding/hex"
	"testing"
)

func TestEncode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.hex, func(t *testing.T) {
			// Convert hex string to bytes
			data, err := hex.DecodeString(tt.hex)
			if err != nil {
				t.Fatalf("hex.DecodeString() error = %v", err)
			}

			// Encode the bytes
			encoded := Encode(data)

			if encoded != tt.serial {
				t.Errorf("Encode() = %s, want %s", encoded, tt.serial)
			}
		})
	}
}
