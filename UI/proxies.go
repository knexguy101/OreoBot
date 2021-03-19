package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func addProxies(w *fyne.Window) *widget.Button {

	win := *w //yay

	accountsLabel := widget.NewLabel("Proxies")
	proxiesSelect := widget.NewSelectEntry([]string{ "Test Proxies"})
	loadProxies := widget.NewButton("Load", func() {

	})
	deleteProxies := widget.NewButton("Delete", func(){

	})
	fSeparator := widget.NewSeparator()
	listName := widget.NewEntry()
	accounts := widget.NewMultiLineEntry()
	saveButton := widget.NewButton("Save", func() {

	})
	accountsItem := widget.NewFormItem("Proxies", accounts)
	separatorLabel := widget.NewFormItem("List Name", listName)
	proxiesForm := widget.NewForm(separatorLabel, accountsItem)

	//add pop up
	pidContainer := container.NewVBox(accountsLabel, proxiesSelect, loadProxies, deleteProxies, fSeparator, proxiesForm, saveButton)
	addPopup := widget.NewPopUp(pidContainer, win.Canvas())
	addButton := widget.NewButton("Proxies", func() {
		addPopup.Show()
		addPopup.Resize(fyne.Size{
			Width: 400,
			Height: 150,
		})
	})
	return addButton
}