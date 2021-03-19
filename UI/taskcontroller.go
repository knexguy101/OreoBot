package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TaskController struct {
	Label *widget.Label
	PidLabel *widget.Label
	StartButton *widget.Button
	EditButton *widget.Button
	DeleteButton *widget.Button
	StopButton *widget.Button
	Scroll *container.Scroll
	Container *fyne.Container
}

func (tc *TaskController) Init() {

	//main window
	tc.StartButton = widget.NewButton("Start", func() {
	})
	tc.EditButton = widget.NewButton("Edit", func() {
	})
	tc.StopButton = widget.NewButton("Stop", func() {
	})
	tc.DeleteButton = widget.NewButton("Delete", func() {
	})
	tc.Label = widget.NewLabel("Task Status")
	tc.PidLabel = widget.NewLabel("123456789")
	tc.Container = container.NewHBox(tc.StartButton, tc.StopButton, tc.EditButton, tc.DeleteButton, tc.PidLabel, tc.Label)
	tc.Scroll = container.NewHScroll(tc.Container)
}

func (tc *TaskController) GetControl() *container.Scroll {
	return tc.Scroll
}

func (tc *TaskController) SetStatus(status string) {
	tc.Label.SetText(status)
}
