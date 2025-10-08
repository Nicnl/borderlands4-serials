package b85

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		serial string
		hex    string
	}{
		{"@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00",
			"21070601906270443339b05391542b85764567854f857785430567054e85638400",
		},
	}

	for _, tt := range tests {
		t.Run(tt.serial, func(t *testing.T) {
			decoded, err := Decode(tt.serial)
			if err != nil {
				t.Fatalf("Decode() error = %v", err)
			}
			if fmt.Sprintf("%x", decoded) != tt.hex {
				t.Errorf("Decode() = %x, want %s", decoded, tt.hex)
			}
		})
	}
}
