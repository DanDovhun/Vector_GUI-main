package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func popup(appl fyne.App, message string) {
	pop := appl.NewWindow("Error")

	pop.SetContent(container.NewVBox(
		widget.NewLabel(message),
		widget.NewButton("Ok", func() {
			pop.Hide()
		}),
	))

	pop.Show()
}

func ResultsTwoDimWindow(appl fyne.App, home fyne.Window, vecA TwoDimVec, vecB TwoDimVec) {
	results := appl.NewWindow("2D Results")

	vecAMag := vecA.Magnitude()
	vecBMag := vecB.Magnitude()

	vecATwoDimAngle := vecA.TwoDimAngle()
	vecBTwoDimAngle := vecB.TwoDimAngle()

	addition := vecA.Add(vecB)
	additionMag := addition.Magnitude()
	additionTwoDimAngle := addition.TwoDimAngle()

	subtraction := vecA.Sub(vecB)
	subtractionMag := subtraction.Magnitude()
	subtractionTwoDimAngle := subtraction.TwoDimAngle()

	results.SetContent(container.NewVBox(
		widget.NewLabel(
			fmt.Sprintf("Vector %v\n", vecA.name)+
				fmt.Sprintf("Coordinates: %v\n", vecA.Coordinates())+
				fmt.Sprintf("Magnitude: %v\n", round(vecAMag, 10))+
				fmt.Sprintf("Angle: %v radians; %v degrees", round(vecATwoDimAngle.radians, 5), round(vecATwoDimAngle.degrees, 5)),
		),

		widget.NewLabel(
			fmt.Sprintf("Vector %v\n", vecB.name)+
				fmt.Sprintf("Coordinates: %v\n", vecB.Coordinates())+
				fmt.Sprintf("Magnitude: %v\n", round(vecBMag, 5))+
				fmt.Sprintf("Angle: %v radians; %v degrees", round(vecBTwoDimAngle.radians, 5), round(vecBTwoDimAngle.degrees, 5)),
		),

		widget.NewLabel(
			fmt.Sprintf("Vector %v\n", addition.name)+
				fmt.Sprintf("Coordinates: %v\n", addition.Coordinates())+
				fmt.Sprintf("Magnitude: %v\n", round(additionMag, 5))+
				fmt.Sprintf("Angle: %v radians; %v degrees", round(additionTwoDimAngle.radians, 5), round(additionTwoDimAngle.degrees, 5)),
		),

		widget.NewLabel(
			fmt.Sprintf("Vector %v\n", subtraction.name)+
				fmt.Sprintf("Coordinates: %v\n", subtraction.Coordinates())+
				fmt.Sprintf("Magnitude: %v\n", round(subtractionMag, 5))+
				fmt.Sprintf("Angle: %v radians; %v degrees", round(subtractionTwoDimAngle.radians, 5), round(subtractionTwoDimAngle.degrees, 5)),
		),

		widget.NewButton("OK", func() {
			home.Show()
			home.SetMaster()
			results.Hide()
		}),
	))

	results.Show()
}

func TwoDimVecWindow(appl fyne.App, home fyne.Window) {
	home.Hide()
	window := appl.NewWindow("2D Vectors")
	window.SetMaster()

	vecALabel := widget.NewLabel("Vector A")
	vecAEntryX := widget.NewEntry()
	vecAEntryY := widget.NewEntry()

	vecBLabel := widget.NewLabel("Vector B")
	vecBEntryX := widget.NewEntry()
	vecBEntryY := widget.NewEntry()

	window.SetContent(container.NewVBox(
		vecALabel,
		vecAEntryX,
		vecAEntryY,

		vecBLabel,
		vecBEntryX,
		vecBEntryY,

		widget.NewButton("Calculate", func() {
			xA, err1 := strconv.ParseFloat(vecAEntryX.Text, 64)
			yA, err2 := strconv.ParseFloat(vecAEntryY.Text, 64)
			xB, err3 := strconv.ParseFloat(vecBEntryX.Text, 64)
			yB, err4 := strconv.ParseFloat(vecBEntryY.Text, 64)

			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				popup(appl, "Please only enter numbers")
				return
			}

			vecA := TwoDimVec{"A", xA, yA}
			vecB := TwoDimVec{"B", xB, yB}

			ResultsTwoDimWindow(appl, window, vecA, vecB)
		}),

		widget.NewButton("Back", func() {
			home.Show()
			home.SetMaster()
			window.Hide()
		}),
	))

	window.Resize(fyne.NewSize(250, 250))
	window.Show()
}

func ThreeDimVecWindow(appl fyne.App, home fyne.Window) {

}

func main() {
	myApp := app.New()
	home := myApp.NewWindow("Vector Caluclator")
	home.SetMaster()

	home.SetContent(container.NewVBox(
		widget.NewButton("2D Vectors", func() {
			TwoDimVecWindow(myApp, home)
		}),

		widget.NewButton("3D Vectors", func() {
			ThreeDimVecWindow(myApp, home)
		}),

		widget.NewButton("Exit", func() {
			home.Close()
		}),
	))

	home.Resize(fyne.NewSize(300, 50))
	home.ShowAndRun()
}
