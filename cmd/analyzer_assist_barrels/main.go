//go:build ignore

package main

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/codex"
	"borderlands_4_serials/b4s/serial"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gopkg.in/yaml.v3"
)

var (
	MODE_SINGLE_PART = false

	loadedItems []codex.LoadedItem

	app   = tview.NewApplication()
	pages = tview.NewPages()

	mainMenuList  = tview.NewList()
	mainMenuFrame = tview.NewFrame(mainMenuList)

	formDataFill  = tview.NewForm()
	dataFillFrame = tview.NewFrame(formDataFill)

	selectedItem         codex.LoadedItem
	selectedSerial       = "<NONE>"
	serialData           = "<NONE>"
	selectedUnknownParts uint32
	serialsToTry         []string

	slot0Serial        = "<NONE>"
	slot0SerialData    = "<NONE>"
	slot0SerialDecoded serial.Serial

	jsonFileName  string
	jsonPartInfos string

	extractSaves = false
)

func updateMainMenuInfos() {
	mainMenuFrame.Clear().
		AddText(fmt.Sprintf("Total loaded items:  %d", len(loadedItems)), true, tview.AlignLeft, tcell.ColorWhite).
		AddText("", true, tview.AlignLeft, tcell.ColorWhite).
		AddText("Serial: "+selectedSerial, true, tview.AlignLeft, tcell.ColorWhite).
		AddText("Data:  "+serialData, true, tview.AlignLeft, tcell.ColorWhite).
		AddText("Unkown parts:  "+fmt.Sprint(selectedUnknownParts), true, tview.AlignLeft, tcell.ColorWhite).
		AddText("Serials to try:  "+fmt.Sprint(len(serialsToTry)), true, tview.AlignLeft, tcell.ColorWhite)
}

func updateDataFillInfos() {
	dataFillFrame.Clear().
		AddText(fmt.Sprintf("Slot0: %s", slot0Serial), true, tview.AlignLeft, tcell.ColorWhite).
		AddText(fmt.Sprintf("Data:  %s", slot0SerialData), true, tview.AlignLeft, tcell.ColorWhite)
}

func initializeCodexInfos() {
	codex.SkipFailedItems = true
	var err error
	loadedItems, _, err = codex.Codex.Load(os.Getenv("CODEX_JSON_RAW_ITEMS"))
	if err != nil {
		panic(err)
	}
}

type _partsInfo struct {
	PartStr string `json:"part_str"`
	Serial  string `json:"serial"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Desc    string `json:"description"`
	Stats   string `json:"stats"`
}

func loadSlot0Serial(foundSlot0Serial string) {
	jsonFileName = ""
	jsonPartInfos = ""

	formDataFill.GetFormItem(0).SetDisabled(true).(*tview.InputField).SetText("barrel")
	formDataFill.GetFormItem(1).SetDisabled(true).(*tview.InputField).SetText("")
	formDataFill.GetFormItem(2).SetDisabled(true).(*tview.InputField).SetText("")
	formDataFill.GetFormItem(3).SetDisabled(true).(*tview.InputField).SetText("")

	slot0Serial = foundSlot0Serial

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

	if slot0SerialData == "" {
		app.QueueUpdateDraw(func() {
			updateDataFillInfos()
		})
		return
	}

	jsonFileName = slot0SerialDecoded.FindPartAtPos(1, false).String()
	jsonFileName = fmt.Sprintf("%32x", md5.Sum([]byte(slot0Serial))) + "_" + strings.ReplaceAll(jsonFileName, ":", "_")
	jsonPartInfos = filepath.Join(os.Getenv("CODEX_JSON_DATABASE_DIR"), jsonFileName+".json")

	formDataFill.GetFormItem(0).SetDisabled(false)
	formDataFill.GetFormItem(1).SetDisabled(false)
	formDataFill.GetFormItem(2).SetDisabled(false)
	formDataFill.GetFormItem(3).SetDisabled(false)
	app.SetFocus(formDataFill.GetFormItem(0))

	// Check if a file exists
	if _, err := os.Stat(jsonPartInfos); err == nil {
		// File exists
		rawData, err := os.ReadFile(jsonPartInfos)
		if err != nil {
			//panic(err)
		} else {
			var partInfos _partsInfo
			err = json.Unmarshal(rawData, &partInfos)
			if err == nil {
				formDataFill.GetFormItem(0).(*tview.InputField).SetText(partInfos.Type)
				formDataFill.GetFormItem(1).(*tview.InputField).SetText(partInfos.Name)
				formDataFill.GetFormItem(2).(*tview.InputField).SetText(partInfos.Desc)
				formDataFill.GetFormItem(3).(*tview.InputField).SetText(partInfos.Stats)
			}
		}

	}

	app.QueueUpdateDraw(func() {
		updateDataFillInfos()
	})
}

//
// INJECTION_YAML_DEST
// INJECTION_SAV

func main() {
	// Load conf from config.json if exists
	{
		pwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		configPath := filepath.Join(pwd, "config.json")
		if _, err := os.Stat(configPath); err == nil {
			rawData, err := os.ReadFile(configPath)
			if err != nil {
				panic(err)
			}

			var conf struct {
				CODEX_JSON_DATABASE_DIR string `json:"CODEX_JSON_DATABASE_DIR"`
				CODEX_JSON_RAW_ITEMS    string `json:"CODEX_JSON_RAW_ITEMS"`
				INJECTION_SAV           string `json:"INJECTION_SAV"`
				INJECTION_USERID        string `json:"INJECTION_USERID"`
				INJECTION_YAML_DEST     string `json:"INJECTION_YAML_DEST"`
				INJECTION_YAML_MODEL    string `json:"INJECTION_YAML_MODEL"`
				MODE_SINGLE_PART        bool   `json:"MODE_SINGLE_PART"`
			}
			err = json.Unmarshal(rawData, &conf)
			if err != nil {
				panic(err)
			}

			if os.Getenv("CODEX_JSON_DATABASE_DIR") == "" {
				os.Setenv("CODEX_JSON_DATABASE_DIR", conf.CODEX_JSON_DATABASE_DIR)
			}
			if os.Getenv("CODEX_JSON_RAW_ITEMS") == "" {
				os.Setenv("CODEX_JSON_RAW_ITEMS", conf.CODEX_JSON_RAW_ITEMS)
			}
			if os.Getenv("INJECTION_SAV") == "" {
				os.Setenv("INJECTION_SAV", conf.INJECTION_SAV)
			}
			if os.Getenv("INJECTION_USERID") == "" {
				os.Setenv("INJECTION_USERID", conf.INJECTION_USERID)
			}
			if os.Getenv("INJECTION_YAML_DEST") == "" {
				os.Setenv("INJECTION_YAML_DEST", conf.INJECTION_YAML_DEST)
			}
			if os.Getenv("INJECTION_YAML_MODEL") == "" {
				os.Setenv("INJECTION_YAML_MODEL", conf.INJECTION_YAML_MODEL)
			}
			MODE_SINGLE_PART = conf.MODE_SINGLE_PART
		}
	}

	initializeCodexInfos()
	updateMainMenuInfos()
	updateDataFillInfos()

	// Save inject loop
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			if extractSaves {
				pwd, err := os.Getwd()
				if err != nil {
					panic(err)
				}

				err = exec.Command(filepath.Join(pwd, "bl4-crypt-cli-38-1-0-7-1758195212.exe"),
					"decrypt",
					"--in",
					os.Getenv("INJECTION_SAV"),
					"--out",
					os.Getenv("INJECTION_YAML_DEST"),
					"--userid",
					os.Getenv("INJECTION_USERID"),
				).Run()
				if err != nil {
					//panic(cmd)
					continue
				} else {
					rawData, err := os.ReadFile(os.Getenv("INJECTION_YAML_DEST"))
					if err != nil {
						//panic(err)
					} else {

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
							//panic(err)
							continue
						} else {
							if len(yamlData.State.Inventory.EquippedInventory.Equipped.Slot0) > 0 {
								if yamlData.State.Inventory.EquippedInventory.Equipped.Slot0[0].Serial != "" {
									foundSerialSlot0 := yamlData.State.Inventory.EquippedInventory.Equipped.Slot0[0].Serial
									if slot0Serial != foundSerialSlot0 {
										loadSlot0Serial(foundSerialSlot0)
									}
								}
							}
						}
					}

				}

			}

		}
	}()

	//pages.SwitchToPage(fmt.Sprintf("page-%d", (page+1)%pageCount))

	mainMenuList.
		AddItem("Comparison mode", "Enter comparison mode", '1', func() {
			jsonFileName = ""
			jsonPartInfos = ""
			slot0Serial = ""
			slot0SerialData = ""
			extractSaves = true
			pages.SwitchToPage("data-fill")
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	funcAutosaveJson := func(_ string) {
		if jsonPartInfos != "" {
			partInfos := _partsInfo{
				PartStr: slot0SerialDecoded.FindPartAtPos(1, false).String(),
				Serial:  slot0Serial,
				Type:    formDataFill.GetFormItem(0).(*tview.InputField).GetText(),
				Name:    formDataFill.GetFormItem(1).(*tview.InputField).GetText(),
				Desc:    formDataFill.GetFormItem(2).(*tview.InputField).GetText(),
				Stats:   formDataFill.GetFormItem(3).(*tview.InputField).GetText(),
			}

			rawData, err := json.MarshalIndent(partInfos, "", "  ")
			if err != nil {
				panic(err)
			}
			err = os.WriteFile(jsonPartInfos, rawData, 0644)
			if err != nil {
				panic(err)
			}

			updateMainMenuInfos()
		}
	}

	formDataFill.
		AddInputField("Type", "", 30, nil, funcAutosaveJson).
		AddInputField("Name", "", 30, nil, funcAutosaveJson).
		AddInputField("Description", "", 60, nil, funcAutosaveJson).
		AddInputField("Stats", "", 30, nil, funcAutosaveJson).
		AddButton("Exit", func() {
			initializeCodexInfos()
			updateMainMenuInfos()
			updateDataFillInfos()
			extractSaves = false
			jsonFileName = ""
			jsonPartInfos = ""
			slot0Serial = ""
			slot0SerialData = ""
			pages.SwitchToPage("menu")
		}).
		SetBorder(true).SetTitle("Fill the data").SetTitleAlign(tview.AlignLeft)

	//formDataFill.GetFormItem(0).SetDisabled(true)
	//formDataFill.GetFormItem(1).SetDisabled(true)
	//formDataFill.GetFormItem(2).SetDisabled(true)
	//formDataFill.GetFormItem(3).SetDisabled(true)

	pages.AddPage("menu", mainMenuFrame, true, true)
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
