package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LoadUI() {
	// init
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.Size{
		Width: 800,
		Height: 800,
	})
	w.SetFixedSize(true)
	seperator := widget.NewSeparator()

	//add
	addButton := addPopUp(&w)

	//profiles
	profilesButton := addProfiles(&w)

	//accounts
	accountsButton := addAccounts(&w)

	//proxies
	proxiesButton := addProxies(&w)

	var con *fyne.Container
	var scroll *container.Scroll
	topButtonContainer := container.NewHBox(addButton, profilesButton, accountsButton, proxiesButton)
	con = container.NewVBox(
		topButtonContainer,
		seperator,
		widget.NewButton("Add Task", func() {
			task := TaskController{}
			task.Init()
			con.Add(task.GetControl())
		}),
	)
	scroll = container.NewVScroll(con)
	w.SetContent(scroll)

	w.ShowAndRun()
}


