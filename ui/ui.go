package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type Todo struct {
	Description string
	Done 		bool
}

func NewTodo(description string) Todo {
	return Todo{description, false}
}

func (t Todo) String() string {
	return fmt.Sprintf("%s - %t", t.Description, t.Done)
}

func Run() {
	app := app.New()
	window := app.NewWindow("TODO App")
	window.Resize(fyne.NewSize(300, 400))


	newtodoDescTxt := widget.NewEntry()
	newtodoDescTxt.PlaceHolder = "New Todo Description..."
	addBtn := widget.NewButton("Add", func() {fmt.Println("Add was clicked")})
	addBtn.Disable()

	newtodoDescTxt.OnChanged = func(s string) {
		addBtn.Disable()

		if len(s) >= 3 {
			addBtn.Enable()
		}
	}

	data := []Todo{
        NewTodo("Some stuff"),
        NewTodo("Some more stuff"),
        NewTodo("Some other things"),
    }

	todos := binding.NewUntypedList()
	for _, todo := range data {
		todos.Append(todo)
	}

	window.SetContent(
		container.NewBorder(
			nil,
			container.NewBorder(
				nil,
				nil,
				nil,
				addBtn,
				newtodoDescTxt,
			),
			nil,
			nil,
			widget.NewListWithData(
				todos,
				func() fyne.CanvasObject {
					return container.NewBorder(
						nil, nil, nil,
						widget.NewCheck("", func(b bool) {}),
						widget.NewLabel(""),
					)
				},
				func(di binding.DataItem, o fyne.CanvasObject) {
					ctr, _ := o.(*fyne.Container)
					label := ctr.Objects[0].(*widget.Label)
					check := ctr.Objects[1].(*widget.Check)
					diu, _ := di.(binding.Untyped).Get()
					todo := diu.(Todo)

					label.SetText(todo.Description)
					check.SetChecked(todo.Done)
				},
			),
		),
	)

	window.ShowAndRun()
}