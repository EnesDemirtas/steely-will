package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	a := app.New()
	w := a.NewWindow("Steely Will")

	w.Resize(fyne.NewSize(650, 500))

	hello := widget.NewLabel("Hello Fyne!")

	clock := widget.NewLabel("")
	updateTime(clock)

	w.SetContent(clock)

	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
		clock,
	))


	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.ShowAndRun()
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}