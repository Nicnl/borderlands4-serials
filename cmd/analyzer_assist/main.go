package main

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/codex_loader"
	"borderlands_4_serials/b4s/serial"
	"borderlands_4_serials/b4s/serial_datatypes/part"
	"borderlands_4_serials/b4s/serial_tokenizer"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gopkg.in/yaml.v3"
)

const pageCount = 5

func mainsdqsd() {
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}

func currentSaveInjectionsString() string {
	if injectSaves {
		return fmt.Sprintf("Save rewrite: %d", saveInjectionCounter)
	} else if extractSaves {
		return fmt.Sprintf("Save extract: %d", saveInjectionCounter)
	} else {
		return "Idle"
	}
}

var (
	loadedItems []codex.LoadedItem

	app   = tview.NewApplication()
	pages = tview.NewPages()

	mainMenuList  = tview.NewList()
	mainMenuFrame = tview.NewFrame(mainMenuList)

	formDataFill  = tview.NewForm()
	dataFillFrame = tview.NewFrame(formDataFill)

	selectedItem         codex_loader.LoadedItem
	selectedSerial       = "<NONE>"
	serialData           = "<NONE>"
	selectedUnknownParts uint32
	serialsToTry         []string

	slot0Serial        = "<NONE>"
	slot0SerialData    = "<NONE>"
	slot0SerialDecoded serial.Serial

	slot1Serial        = "<NONE>"
	slot1SerialData    = "<NONE>"
	slot1SerialDecoded serial.Serial

	modeDiffParts []part.Part

	allParts     = make(map[string]part.Part)
	partToItems  = make(map[string][]*codex_loader.LoadedItem)
	unknownParts = []part.Part{}

	injectSaves          = false
	extractSaves         = false
	saveInjectionCounter = 0
	saveInjectionModal   = tview.NewModal().
				SetText(currentSaveInjectionsString()).
				AddButtons([]string{"Stop"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			injectSaves = false
			pages.SwitchToPage("menu")
		})
)

func updateUnknownParts() {
	// Count unknown parts
	unknownParts = []part.Part{}
	for partStr := range allParts {
		_, isKnown := codex_loader.Parts[partStr]
		if !isKnown {
			unknownParts = append(unknownParts, allParts[partStr])
		}
	}
}

func updateMainMenuInfos() {
	mainMenuFrame.Clear().
		AddText(fmt.Sprintf("Total loaded items:  %d", len(loadedItems)), true, tview.AlignLeft, tcell.ColorWhite).
		AddText(fmt.Sprintf("Total unique parts:  %d", len(allParts)), true, tview.AlignLeft, tcell.ColorWhite).
		AddText(fmt.Sprintf("Total unknown parts: %d", len(unknownParts)), true, tview.AlignLeft, tcell.ColorWhite).
		AddText("", true, tview.AlignLeft, tcell.ColorWhite).
		AddText("Serial: "+selectedSerial, true, tview.AlignLeft, tcell.ColorWhite).
		AddText("Data:  "+serialData, true, tview.AlignLeft, tcell.ColorWhite).
		AddText("Unkown parts:  "+fmt.Sprint(selectedUnknownParts), true, tview.AlignLeft, tcell.ColorWhite).
		AddText("Serials to try:  "+fmt.Sprint(len(serialsToTry)), true, tview.AlignLeft, tcell.ColorWhite)
}

func updateDataFillInfos() {
	diffTxt := ""
	for i, p := range modeDiffParts {
		if i != 0 {
			diffTxt += " "
		}
		diffTxt += p.String()
	}

	dataFillFrame.Clear().
		AddText(fmt.Sprintf("Slot0: %s", slot0Serial), true, tview.AlignLeft, tcell.ColorWhite).
		AddText(fmt.Sprintf("Data:  %s", slot0SerialData), true, tview.AlignLeft, tcell.ColorWhite).
		AddText(fmt.Sprintf("Slot1: %s", slot1Serial), true, tview.AlignLeft, tcell.ColorWhite).
		AddText(fmt.Sprintf("Data:  %s", slot1SerialData), true, tview.AlignLeft, tcell.ColorWhite).
		AddText(fmt.Sprint("Diff:  ", diffTxt), true, tview.AlignLeft, tcell.ColorWhite)
}

func initializeCodexInfos() {
	codex_loader.SkipFailedItems = true
	var err error
	loadedItems, _, err = codex_loader.Codex.Load(os.Getenv("CODEX_JSON_RAW_ITEMS"))
	if err != nil {
		panic(err)
	}

	// Extract all unique parts
	for _, item := range loadedItems {
		for _, block := range item.Parsed.Blocks {
			if block.Token == serial_tokenizer.TOK_PART {
				switch block.Part.SubType {
				case part.SUBTYPE_NONE, part.SUBTYPE_INT:
					allParts[block.Part.String()] = block.Part
					partToItems[block.Part.String()] = append(partToItems[block.Part.String()], &item)
				case part.SUBTYPE_LIST:
					// We explode list parts into multiple parts
					for _, subpartValue := range block.Part.Values {
						p := part.Part{
							Index:   block.Part.Index,
							SubType: part.SUBTYPE_INT,
							Values:  []uint32{subpartValue},
						}
						allParts[p.String()] = p
						partToItems[p.String()] = append(partToItems[p.String()], &item)
					}
				}
			}
		}
	}

	updateUnknownParts()
}

func pickItem() {
	// Pick a random part from the unknown parts
	if len(unknownParts) == 0 {
		return
	}

	pickedPart := unknownParts[rand.Intn(len(unknownParts))]
	// Pick a random item that contains this part
	itemsWithPart := partToItems[pickedPart.String()]
	if len(itemsWithPart) == 0 {
		return
	}

	selectedItem = *itemsWithPart[rand.Intn(len(itemsWithPart))]
	selectedSerial = selectedItem.Serial
	serialData = selectedItem.Parsed.String()

	selectedUnknownParts = 0

	// Split the item blocks, including subpart of type list
	itemBlocks := make([]serial.Block, 0, len(selectedItem.Parsed.Blocks)*2)
	for _, block := range selectedItem.Parsed.Blocks {
		if block.Token == serial_tokenizer.TOK_PART {
			switch block.Part.SubType {
			case part.SUBTYPE_NONE, part.SUBTYPE_INT:
				if _, isKnown := codex_loader.Parts[block.Part.String()]; !isKnown {
					selectedUnknownParts++
				}
				itemBlocks = append(itemBlocks, block)
			case part.SUBTYPE_LIST:
				// Split list parts into multiple parts
				for _, subpartValue := range block.Part.Values {
					p := part.Part{
						Index:   block.Part.Index,
						SubType: part.SUBTYPE_INT,
						Values:  []uint32{subpartValue},
					}
					if _, isKnown := codex_loader.Parts[p.String()]; !isKnown {
						selectedUnknownParts++
					}
					itemBlocks = append(itemBlocks, serial.Block{Token: serial_tokenizer.TOK_PART, Part: p})
				}
			}
		} else {
			itemBlocks = append(itemBlocks, block)
		}
	}

	// We're making new items with one part missing each time
	serialsToTry = make([]string, 0)

	// The first is the original serial
	serialsToTry = append(serialsToTry, selectedSerial)

	for i := range itemBlocks {
		newBlocks := make([]serial.Block, len(itemBlocks))
		copy(newBlocks, itemBlocks)

		block := &itemBlocks[i]
		if block.Token != serial_tokenizer.TOK_PART {
			continue
		}

		// is it unknown?
		if _, isKnown := codex.Parts[block.Part.String()]; isKnown {
			continue
		}

		// Remove this part
		newBlocks = append(newBlocks[:i], newBlocks[i+1:]...)

		// Create a new serial
		newSerialData := serial.Serial{Blocks: newBlocks}
		newSerialBytes := serial.Serialize(newSerialData)
		newSerialB85 := b85.Encode(newSerialBytes)
		serialsToTry = append(serialsToTry, newSerialB85)
	}

	app.QueueUpdateDraw(func() {
		updateMainMenuInfos()
	})
}

func writeYamlFile() {
	fileData, err := os.ReadFile(os.Getenv("INJECTION_YAML_MODEL"))
	if err != nil {
		panic(err)
	}

	yamlContent := ""
	for i, s := range serialsToTry {
		yamlContent += fmt.Sprintf("        slot_%d:\n", i)
		yamlContent += fmt.Sprintf("          serial: '%s'\n", s)
	}
	yamlContent += "    equipped_inventory:\n"
	yamlContent += "      equipped:\n"
	yamlContent += "        slot_0:\n"
	yamlContent += fmt.Sprintf("        - serial: '%s'\n", selectedSerial)

	fileDataStr := string(fileData)
	fileDataStr = strings.Replace(fileDataStr, "%INJECT_ITEMS_HERE%", yamlContent, -1)

	err = os.WriteFile(os.Getenv("INJECTION_YAML_DEST"), []byte(fileDataStr), 0644)
	if err != nil {
		panic(err)
	}
}

func compareSlot01Serials(foundSlot0Serial string, foundSlot1Serial string) {
	slot0Serial = foundSlot0Serial
	slot1Serial = foundSlot1Serial

	// Decode 0
	{
		decodedBytes, err := b85.Decode(slot0Serial)
		if err != nil {
			slot0SerialData = "ERROR"
		} else {
			decodedSerial, err := serial.Deserialize(decodedBytes)
			if err != nil {
				slot0SerialData = "ERROR"
			} else {
				slot0SerialData = decodedSerial.String()
				slot0SerialDecoded = decodedSerial
			}
		}
	}

	// Decode 1
	{
		decodedBytes, err := b85.Decode(slot1Serial)
		if err != nil {
			slot1SerialData = "ERROR"
		} else {
			decodedSerial, err := serial.Deserialize(decodedBytes)
			if err != nil {
				slot1SerialData = "ERROR"
			} else {
				slot1SerialData = decodedSerial.String()
				slot1SerialDecoded = decodedSerial
			}
		}
	}

	// Find part differences
	allPartsIn1 := make(map[string]part.Part)
	for _, block := range slot1SerialDecoded.Blocks {
		if block.Token == serial_tokenizer.TOK_PART {
			switch block.Part.SubType {
			case part.SUBTYPE_NONE, part.SUBTYPE_INT:
				allPartsIn1[block.Part.String()] = block.Part
			case part.SUBTYPE_LIST:
				// We explode list parts into multiple parts
				for _, subpartValue := range block.Part.Values {
					p := part.Part{
						Index:   block.Part.Index,
						SubType: part.SUBTYPE_LIST,
						Values:  []uint32{subpartValue},
					}
					allPartsIn1[p.String()] = p
				}
			}
		}
	}

	foundNewParts := []part.Part{}
	for _, block := range slot0SerialDecoded.Blocks {
		if block.Token == serial_tokenizer.TOK_PART {
			switch block.Part.SubType {
			case part.SUBTYPE_NONE, part.SUBTYPE_INT:
				if _, exists := allPartsIn1[block.Part.String()]; !exists {
					foundNewParts = append(foundNewParts, block.Part)
				}
			case part.SUBTYPE_LIST:
				// We explode list parts into multiple parts
				for _, subpartValue := range block.Part.Values {
					p := part.Part{
						Index:   block.Part.Index,
						SubType: part.SUBTYPE_LIST,
						Values:  []uint32{subpartValue},
					}
					if _, exists := allPartsIn1[p.String()]; !exists {
						foundNewParts = append(foundNewParts, p)
					}
				}
			}
		}
	}
	modeDiffParts = foundNewParts

	app.QueueUpdateDraw(func() {
		updateDataFillInfos()
	})
}

//
// INJECTION_YAML_DEST
// INJECTION_SAV

func main() {
	initializeCodexInfos()
	updateMainMenuInfos()
	updateDataFillInfos()

	// Save inject loop
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			if injectSaves {
				saveInjectionCounter++
				// run  bl4-crypt-cli-38-1-0-7-1758195212.exe
				// get current directory
				pwd, err := os.Getwd()
				if err != nil {
					panic(err)
				}

				cmd := exec.Command(filepath.Join(pwd, "bl4-crypt-cli-38-1-0-7-1758195212.exe"),
					"encrypt",
					"--in",
					os.Getenv("INJECTION_YAML_DEST"),
					"--out",
					os.Getenv("INJECTION_SAV"),
					"--userid",
					os.Getenv("INJECTION_USERID"),
				).Run()
				if cmd != nil {
					panic(cmd)
				}

				app.QueueUpdateDraw(func() {
					saveInjectionModal.SetText(currentSaveInjectionsString())
				})

				continue
			} else if extractSaves {
				saveInjectionCounter++
				pwd, err := os.Getwd()
				if err != nil {
					panic(err)
				}

				cmd := exec.Command(filepath.Join(pwd, "bl4-crypt-cli-38-1-0-7-1758195212.exe"),
					"decrypt",
					"--in",
					os.Getenv("INJECTION_SAV"),
					"--out",
					os.Getenv("INJECTION_YAML_DEST"),
					"--userid",
					os.Getenv("INJECTION_USERID"),
				).Run()
				if cmd != nil {
					//panic(cmd)
					continue
				} else {
					rawData, err := os.ReadFile(os.Getenv("INJECTION_YAML_DEST"))
					if err != nil {
						panic(err)
					}

					var yamlData struct {
						State struct {
							Inventory struct {
								EquippedInventory struct {
									Equipped struct {
										Slot0 []struct {
											Serial string `yaml:"serial"`
										} `yaml:"slot_0"`
										Slot1 []struct {
											Serial string `yaml:"serial"`
										} `yaml:"slot_1"`
									} `yaml:"equipped"`
								} `yaml:"equipped_inventory"`
							} `yaml:"inventory"`
						} `yaml:"state"`
					}

					// Check if slot_1 is set
					err = yaml.Unmarshal(rawData, &yamlData)
					if err != nil {
						panic(err)
					}

					if len(yamlData.State.Inventory.EquippedInventory.Equipped.Slot0) > 0 {
						if yamlData.State.Inventory.EquippedInventory.Equipped.Slot0[0].Serial != "" {
							if len(yamlData.State.Inventory.EquippedInventory.Equipped.Slot1) > 0 {
								if yamlData.State.Inventory.EquippedInventory.Equipped.Slot1[0].Serial != "" {
									foundSerialSlot0 := yamlData.State.Inventory.EquippedInventory.Equipped.Slot0[0].Serial
									foundSerialSlot1 := yamlData.State.Inventory.EquippedInventory.Equipped.Slot1[0].Serial
									if slot1Serial != foundSerialSlot1 {
										compareSlot01Serials(foundSerialSlot0, foundSerialSlot1)
									}
								}
							}
						}
					}

					app.QueueUpdateDraw(func() {
						saveInjectionModal.SetText(currentSaveInjectionsString())
					})
				}

			}

		}
	}()

	//pages.SwitchToPage(fmt.Sprintf("page-%d", (page+1)%pageCount))

	mainMenuList.
		AddItem("Pick item", "Pick item to start a test round", '1', func() {
			go pickItem()
		}).
		AddItem("Inject save", "Inject the items in the save game", '2', func() {
			if selectedSerial == "<NONE>" {
				return
			}

			writeYamlFile()

			saveInjectionCounter = 0
			injectSaves = true
			pages.SwitchToPage("save-rewrite")
		}).
		AddItem("Comparison mode", "Enter comparison mode", '3', func() {
			extractSaves = true
			pages.SwitchToPage("data-fill")
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	formDataFill.
		AddInputField("Type", "", 30, nil, nil).
		AddInputField("Description", "", 60, nil, nil).
		AddInputField("Stats", "", 30, nil, nil).
		SetBorder(true).SetTitle("Fill the data").SetTitleAlign(tview.AlignLeft)

	pages.AddPage("menu", mainMenuFrame, true, true)
	pages.AddPage("save-rewrite", saveInjectionModal, true, false)
	pages.AddPage("data-fill", dataFillFrame, true, false)

	pages.AddPage("menuqsd2", tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
		SetBorders(2, 2, 2, 2, 4, 4).
		AddText("Header left", true, tview.AlignLeft, tcell.ColorWhite).
		AddText("Header left", true, tview.AlignLeft, tcell.ColorWhite).
		AddText("Header right", true, tview.AlignRight, tcell.ColorWhite).
		AddText("Header second middle", true, tview.AlignCenter, tcell.ColorRed).
		AddText("Footer middle", false, tview.AlignCenter, tcell.ColorGreen).
		AddText("Footer second middle", false, tview.AlignCenter, tcell.ColorGreen), true, false)

	//time.Sleep(5 * time.Second)
	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}
