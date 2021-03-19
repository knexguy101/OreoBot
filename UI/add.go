package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func addPopUp(w *fyne.Window) *widget.Button {

	win := *w //yay

	//add pop up
	pidLabel := widget.NewLabel("Pid")
	pid := widget.NewEntry()
	profileLabel := widget.NewLabel("Profiles")
	profiles := widget.NewSelectEntry([]string{
		"Profile 1",
		"Profile 2",
	})
	accountLabel := widget.NewLabel("Account")
	accounts := widget.NewSelectEntry([]string{
		"Account 1",
		"Account 2",
	})
	proxiesLabel := widget.NewLabel("Proxies")
	proxies := widget.NewSelectEntry([]string{
		"Proxies 1",
		"Proxies 2",
	})
	addTaskButton := widget.NewButton("Add Task", func() {

	})
	pidContainer := container.NewVBox(pidLabel, pid, profileLabel, profiles, accountLabel, accounts, proxiesLabel, proxies, addTaskButton)
	addPopup := widget.NewPopUp(pidContainer, win.Canvas())
	addButton := widget.NewButton("Add Task", func() {
		addPopup.Show()
		addPopup.Resize(fyne.Size{
			Width: 200,
			Height: 350,
		})
	})
	return addButton
}
