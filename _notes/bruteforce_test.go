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

func TestFileToSlots1111(t *testing.T) {
	rawSerials := _loadLines("101_bruteforce_1-to-16-bits")
	_serialsToYaml(rawSerials, "101_bruteforce_1-to-16-bits_rawslots")
}

func TestFileToSlots2(t *testing.T) {
	rawSerials := _loadLines("101_bruteforce_steps/AA_half2")
	_serialsToYaml(rawSerials, "101_bruteforce_steps/AA_half2_slots")
}

func TestFileExcludeCollapses(t *testing.T) {
	rawSerials := _loadLines("101_bruteforce_steps/AA_half1")
	collapses := _loadLines("101_bruteforce_steps/AA_half1_collapse")
	excludeSerials := _loadLines("101_bruteforce_steps/list_of_collapses")

	if len(rawSerials) != len(collapses) {
		fmt.Println(len(rawSerials), len(collapses))
		panic("rawSerials and collapses must have the same length")
	}

	excludeMap := map[string]bool{}
	for _, line := range excludeSerials {
		excludeMap[line] = true
	}

	workableSerials := []string{}
	for i := range rawSerials {
		originalLine := rawSerials[i]
		collapseLine := collapses[i]
		if _, found := excludeMap[collapseLine]; !found && collapseLine != "" {
			if len(collapseLine) < 45 {
				continue
			}
			workableSerials = append(workableSerials, originalLine+" => "+collapseLine)
		}
	}

	_serialsToTxt(workableSerials, "101_bruteforce_steps/BB_half1_no_collapses")
}

func TestFileExcludeCollapses2(t *testing.T) {
	rawSerials := _loadLines("101_bruteforce_steps/AA_half2")
	collapses := _loadLines("101_bruteforce_steps/AA_half2_collapse")
	excludeSerials := _loadLines("101_bruteforce_steps/list_of_collapses")

	if len(rawSerials) != len(collapses) {
		fmt.Println(len(rawSerials), len(collapses))
		panic("rawSerials and collapses must have the same length")
	}

	excludeMap := map[string]bool{}
	for _, line := range excludeSerials {
		excludeMap[line] = true
	}

	workableSerials := []string{}
	for i := range rawSerials {
		originalLine := rawSerials[i]
		collapseLine := collapses[i]
		if _, found := excludeMap[collapseLine]; !found && collapseLine != "" {
			if len(collapseLine) <= 45 {
				continue
			}
			workableSerials = append(workableSerials, originalLine+" => "+collapseLine)
		}
	}

	_serialsToTxt(workableSerials, "101_bruteforce_steps/BB_half2_no_collapses")
}
