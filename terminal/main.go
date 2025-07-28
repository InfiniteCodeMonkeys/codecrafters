package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("germ")

	ui := widget.NewTextGrid()       // Create a new TextGrid
	ui.SetText("I'm on a terminal!") // Set text to display

	w.SetContent(ui)
	w.Resize(fyne.NewSize(900, 700))
	w.Show()

	a.Run()

	w.ShowAndRun()

}
