package _notes

import (
	"borderlands_4_serials/lib/b85"
	"borderlands_4_serials/lib/bit_reader"
	"borderlands_4_serials/lib/serial_tokenizer"
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

type workableSerial struct {
	original  string
	collapsed string
}

var globalWorkableSerials []workableSerial

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
			if len(collapseLine) < 23 {
				continue
			}
			if strings.HasPrefix(collapseLine, "@Ugy3L+2@aC}/NsC0/Nnm") {
				continue
			}
			workableSerials = append(workableSerials, originalLine+" => "+collapseLine)
			globalWorkableSerials = append(globalWorkableSerials, workableSerial{original: originalLine, collapsed: collapseLine})
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
			if len(collapseLine) <= 23 {
				continue
			}
			if strings.HasPrefix(collapseLine, "@Ugy3L+2@aC}/NsC0/Nnm") {
				continue
			}
			workableSerials = append(workableSerials, originalLine+" => "+collapseLine)
			globalWorkableSerials = append(globalWorkableSerials, workableSerial{original: originalLine, collapsed: collapseLine})
		}
	}

	_serialsToTxt(workableSerials, "101_bruteforce_steps/BB_half2_no_collapses")
}

func TestFilterExcludeCollapsesPerLevel(t *testing.T) {
	t.Run("TestFileExcludeCollapses", TestFileExcludeCollapses)
	t.Run("TestFileExcludeCollapses2", TestFileExcludeCollapses2)

	mapSerialsPerLevel := make(map[int][]workableSerial)

	for _, serial := range globalWorkableSerials {
		data, err := b85.Decode(serial.collapsed)
		if err != nil {
			panic(err)
		}
		tokenizer := serial_tokenizer.NewTokenizer(data)
		itemLevel, _, _ := tokenizer.Parse()

		fmt.Println(itemLevel, serial.collapsed)

		if itemLevel >= 1 && itemLevel <= 50 {
			mapSerialsPerLevel[itemLevel] = append(mapSerialsPerLevel[itemLevel], serial)
		}
	}

	output := ""
	for level := 1; level <= 50; level++ {
		didStuff := false

		serials, found := mapSerialsPerLevel[level]
		if !found {
			continue
		}
		output += fmt.Sprintln("Level:", level)
		for _, serial := range serials {
			//fmt.Sprintln("  ", serial.original, "=>", serial.collapsed)
			data, err := b85.Decode(serial.original)
			if err != nil {
				panic(err)
			}

			bitstreamPrefix := "001000010000011100000110000000011001000001"
			bitstreamMarker := "001000100001100110011100110110000010100111001000101010100001010111000010101110110010001010110011110000101010011111000010101110111100001010100001100000101011001110000010101001110100001010110001110"

			rawBitstream := bit_reader.NewBitReader(data).StringAfter()
			if strings.HasPrefix(rawBitstream, bitstreamPrefix) {
				markerpos := strings.Index(rawBitstream, bitstreamMarker)
				if markerpos != -1 {
					// Strip the prefix
					didStuff = true
					bitstream := rawBitstream[len(bitstreamPrefix):markerpos]

					output += fmt.Sprintln("  ", bitstream[0:3], bitstream[3:])
				}
			}
		}
		if didStuff {
			output += fmt.Sprintln()
		}
	}

	err := os.WriteFile("101_bruteforce_steps/DD_final_output_bitstream", []byte(output), 0644)
	if err != nil {
		panic(err)
	}
}
