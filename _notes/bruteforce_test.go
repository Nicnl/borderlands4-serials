package _notes

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func _loadLines(filePath string) []string {
	rawContent, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// split by lines
	lines := strings.Split(string(rawContent), "\n")
	for i := range lines {
		lines[i] = strings.Trim(lines[i], "\r")
		lines[i] = strings.Trim(lines[i], "\n")
	}

	return lines
}

func _serialsToYaml(serials []string, filePath string) {
	output := strings.Builder{}
	slotCounter := 0
	for _, serial := range serials {
		if strings.HasPrefix(serial, "@U") {
			serial = strings.Trim(serial, "\r")
			serial = strings.Trim(serial, "\n")
			output.WriteString("        slot_" + fmt.Sprintf("%d", slotCounter) + ":\n")
			output.WriteString("          serial: '" + serial + "'\n")
			slotCounter++
		}
	}

	err := os.WriteFile(filePath, []byte(output.String()), 0644)
	if err != nil {
		panic(err)
	}
}

func _serialsToTxt(serials []string, filePath string) {
	output := strings.Builder{}
	for _, serial := range serials {
		if strings.HasPrefix(serial, "@U") {
			serial = strings.Trim(serial, "\r")
			serial = strings.Trim(serial, "\n")
			output.WriteString(serial + "\n")
		}
	}

	err := os.WriteFile(filePath, []byte(output.String()), 0644)
	if err != nil {
		panic(err)
	}
}

func TestFileToSlots(t *testing.T) {
	rawSerials := _loadLines("101_bruteforce_1-to-16-bits")
	_serialsToYaml(rawSerials, "101_bruteforce_1-to-16-bits_slots")
}

func TestLinesExcludingOthers(t *testing.T) {
	rawSerials := _loadLines("101_bruteforce_1-to-16-bits")
	excludeSerials := _loadLines("101_bruteforce_unknown")

	excludeMap := map[string]bool{}
	for _, line := range excludeSerials {
		excludeMap[line] = true
	}

	workableSerials := []string{}
	for _, line := range rawSerials {
		if _, found := excludeMap[line]; !found && line != "" {
			workableSerials = append(workableSerials, line)
		}
	}

	_serialsToTxt(workableSerials, "101_bruteforce_steps/AA_half1")
}

func TestFileToSlots2(t *testing.T) {
	rawSerials := _loadLines("101_bruteforce_steps/AA_half1")
	_serialsToYaml(rawSerials, "101_bruteforce_steps/AA_half1_slots")
}
