package main

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/serial"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	var (
		currentBase85 = ""
		currentParts  = ""
		currentStatus = "Ready"
		currentSerial serial.Serial
	)

	a := app.New()
	w := a.NewWindow("Borderlands 4 Serial Converter")

	labelBase85 := widget.NewLabel("Base85:")
	inputBase85 := widget.NewEntry()
	inputBase85.SetPlaceHolder("Enter serial here:   @Ugy3L+2}TYg...")

	labelParts := widget.NewLabel("Parts:")
	inputParts := widget.NewEntry()
	inputParts.SetPlaceHolder("Enter parts here:   24, 0, 1, 50| 2, 3379...")

	labelStatus := widget.NewLabel("Status: " + currentStatus)

	deferUpdateData := func(updateBase85 bool, updateParts bool) {
		if updateBase85 {
			inputBase85.SetText(currentBase85)
		}
		if updateParts {
			inputParts.SetText(currentParts)
		}

		if currentBase85 != "" && currentParts != "" {
			additionalData := ""
			if level, found := currentSerial.FindLevel(); found {
				additionalData += "L" + fmt.Sprint(level)
			}
			if additionalData != "" {
				additionalData += " "
			}

			if itemType, found := currentSerial.FindItemType(); found {
				additionalData += itemType.Type + ":" + itemType.Manufacturer
			} else {
				additionalData += "?"
			}
			if additionalData != "" {
				currentStatus += " (" + additionalData + ")"
			}
		}

		labelStatus.SetText("Status: " + currentStatus)
	}

	inputBase85.OnChanged = func(str string) {
		if str == currentBase85 {
			return
		}

		defer deferUpdateData(false, true)

		data, err := b85.Decode(str)
		if err != nil {
			currentBase85 = ""
			currentParts = ""
			currentStatus = "ERROR: " + err.Error()
			currentSerial = serial.Serial{}
			return
		}

		deserialized, err := serial.Deserialize(data)
		if err != nil {
			currentBase85 = ""
			currentParts = ""
			currentStatus = "ERROR: " + err.Error()
			currentSerial = serial.Serial{}
			return
		}

		currentBase85 = str
		currentParts = deserialized.String()
		currentStatus = "OK"
		currentSerial = deserialized
	}

	inputParts.OnChanged = func(str string) {
		if str == currentParts {
			return
		}

		defer deferUpdateData(true, false)

		var s serial.Serial
		err := s.FromString(str)
		if err != nil {
			currentBase85 = ""
			currentParts = ""
			currentStatus = "ERROR: " + err.Error()
			currentSerial = serial.Serial{}
			return
		}

		serialized := serial.Serialize(s)
		encoded := b85.Encode(serialized)

		currentBase85 = encoded
		currentParts = str
		currentStatus = "OK"
		currentSerial = s
	}

	w.SetContent(container.NewVBox(
		labelBase85,
		inputBase85,
		labelParts,
		inputParts,
		labelStatus,
	))
	w.Resize(fyne.NewSize(700, 1))
	w.ShowAndRun()
}
