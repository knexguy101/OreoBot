package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func addAccounts(w *fyne.Window) *widget.Button {

	win := *w //yay

	accountsLabel := widget.NewLabel("Accounts (User:Pass)")
	accounts := widget.NewMultiLineEntry()
	saveButton := widget.NewButton("Save", func() {

	})

	//add pop up
	pidContainer := container.NewVBox(accountsLabel, accounts, saveButton)
	addPopup := widget.NewPopUp(pidContainer, win.Canvas())
	addButton := widget.NewButton("Accounts", func() {
		addPopup.Show()
		addPopup.Resize(fyne.Size{
			Width: 400,
			Height: 150,
		})
	})
	return addButton
}