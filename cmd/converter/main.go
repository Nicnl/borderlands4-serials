package main

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/codex"
	"borderlands_4_serials/b4s/serial"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	var (
		currentBase85      = ""
		currentParts       = ""
		currentStatus      = "Ready"
		currentStatusError = false
		currentItem        codex.Item
	)

	a := app.New()
	w := a.NewWindow("Borderlands 4 Deserializer v1.0 / By @Nicnl and @InflamedSebi")

	labelBase85 := widget.NewLabel("Base85:")
	inputBase85 := widget.NewEntry()
	inputBase85.SetPlaceHolder("Enter serial here:   @Ugy3L+2}TYg...")
	//inputBase85.TextStyle = fyne.TextStyle{Monospace: true}

	labelParts := widget.NewLabel("Parts:")
	inputParts := widget.NewEntry()
	inputParts.SetPlaceHolder("Enter parts here:   24, 0, 1, 50| 2, 3379...")
	inputParts.TextStyle = fyne.TextStyle{Monospace: true}

	labelStatus := widget.NewLabel("Status: " + currentStatus)

	deferUpdateData := func(updateBase85 bool, updateParts bool) {
		if updateBase85 {
			inputBase85.SetText(currentBase85)
		}
		if updateParts {
			inputParts.SetText(currentParts)
		}

		if currentBase85 != "" && currentParts != "" {
			additionalDataArr := []string{}

			if level, found := currentItem.Level(); found {
				additionalDataArr = append(additionalDataArr, "L"+fmt.Sprint(level))
			}

			if itemType, found := currentItem.Type(); found {
				additionalDataArr = append(additionalDataArr, itemType.Type+":"+itemType.Manufacturer)
			} else {
				additionalDataArr = append(additionalDataArr, "?")
			}

			if baseType, found := currentItem.BaseBarrel(); found {
				additionalDataArr = append(additionalDataArr, baseType.Name)
			}

			if len(additionalDataArr) > 0 {
				currentStatus += " (" + strings.Join(additionalDataArr, " ") + ")"
			}
		}

		labelStatus.SetText("Status: " + currentStatus)
		if currentStatusError {
		} else {
		}
	}

	inputBase85.OnChanged = func(serialB85 string) {
		if serialB85 == currentBase85 {
			return
		}

		defer deferUpdateData(false, true)

		item, err := codex.Deserialize(serialB85)
		if err != nil {
			currentBase85 = ""
			currentParts = ""
			//currentStatus = "ERROR: " + err.Error()
			currentStatus = "ERROR: Invalid serial"
			currentItem = codex.Item{}
			currentStatusError = true
			return
		}

		currentBase85 = serialB85
		currentParts = item.Serial.String()
		currentStatus = "OK"
		currentItem = *item
		currentStatusError = false
	}

	inputParts.OnChanged = func(str string) {
		if str == currentParts {
			return
		}

		// Add a terminator if missing
		if !strings.HasSuffix(str, "|") {
			str = str + "|"
		}

		// If more than two terminators, keep only the first two
		for len(str) >= 2 && str[len(str)-1] == '|' && str[len(str)-2] == '|' {
			str = str[:len(str)-1]
		}

		defer deferUpdateData(true, false)

		var s serial.Serial
		err := s.FromString(str)
		if err != nil {
			currentBase85 = ""
			currentParts = ""
			//currentStatus = "ERROR: " + err.Error()
			currentStatus = "ERROR: Invalid parts"
			currentItem = codex.Item{}
			currentStatusError = true
			return
		}

		serialized := serial.Serialize(s)
		encoded := b85.Encode(serialized)

		currentBase85 = encoded
		currentParts = str
		currentStatus = "OK"
		currentItem = codex.Item{
			B85:    currentBase85,
			Serial: s,
		}
		currentStatusError = false
	}

	w.SetContent(container.NewVBox(
		labelBase85,
		inputBase85,
		labelParts,
		inputParts,
		labelStatus,
	))
	w.Resize(fyne.NewSize(850, 1))
	w.SetFixedSize(true)
	w.ShowAndRun()
}
