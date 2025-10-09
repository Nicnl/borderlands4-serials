package tools

import (
	"fmt"
	"testing"
)

func serialsToYaml(serials []string) {
	for i, serial := range serials {
		fmt.Printf("        slot_%d:\n", i)
		fmt.Printf("          serial: '%s'\n", serial)
	}
}

func TestShortenSerial(t *testing.T) {
	serialsToYaml(ShortenSerial("@Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00"))
}
