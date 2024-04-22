package ui

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/EnesDemirtas/steely-will/core"
	"github.com/EnesDemirtas/steely-will/storage"
)

func NewHabitFromDataItem(item binding.DataItem) core.Habit {
	v, _ := item.(binding.Untyped).Get()
	return v.(core.Habit)
}

func Run() error {
	storage := storage.Storage{
		Habits: []core.Habit{},
	}

	coreHabits, err := storage.LoadAll()
	if err != nil {
		return err
	}

	habits := binding.NewUntypedList()
	for _, habit := range coreHabits {
		habits.Append(habit)
	}

	app := app.New()
	window := app.NewWindow("Steely Will")
	window.Resize(fyne.NewSize(650, 500))

	newHabitName := widget.NewEntry()
	newHabitName.PlaceHolder = "Habit name..."
	createBtn := widget.NewButton("Create", func() {
		habit := core.New(
			newHabitName.Text,
			time.Now(),
		)

		if err := storage.Store(*habit); err != nil {
			fmt.Println(err)
		}

		habits.Append(*habit)

		newHabitName.Text = ""
	})
	createBtn.Disable()

	newHabitName.OnChanged = func(s string) {
		createBtn.Disable()

		if len(s) >= 3 {
			createBtn.Enable()
		}
	}


	window.SetContent(
		container.NewBorder(
			nil,
			container.NewBorder(
				nil,
				container.NewVBox(
					newHabitName,
					createBtn,
				),
				nil, nil,
			),
			nil,
			nil,
			container.NewGridWrap(
				fyne.NewSize(650, 500),
				widget.NewListWithData(
					habits,
					func() fyne.CanvasObject {
						return container.NewVBox(
								widget.NewLabel(""),
								widget.NewLabel(""),
						)
					},
					func(di binding.DataItem, o fyne.CanvasObject) {
						ctr, _ := o.(*fyne.Container)
						nameLabel := ctr.Objects[0].(*widget.Label)
						failStackLabel := ctr.Objects[1].(*widget.Label)
						habit := NewHabitFromDataItem(di)
						nameLabel.SetText(habit.Name)
						failStackLabel.SetText(strconv.FormatInt(int64(habit.FailStack), 10))
					},
				),
			),
		),
	)

	window.ShowAndRun()
	return nil
}