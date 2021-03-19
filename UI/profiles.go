package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func addProfiles(w *fyne.Window) *widget.Button {

	win := *w //yay

	profileSelect := widget.NewSelectEntry([]string{ "Profile 1" })
	loadProfile := widget.NewButton("Load", func() {

	})
	deleteProfile := widget.NewButton("Delete", func(){

	})
	fSeparator := widget.NewSeparator()

	title := widget.NewEntry()
	titleItem := widget.NewFormItem("Title", title)
	firstname := widget.NewEntry()
	firstnameItem := widget.NewFormItem("First Name", firstname)
	lastname := widget.NewEntry()
	lastnameItem := widget.NewFormItem("Last Name", lastname)
	address := widget.NewEntry()
	addressItem := widget.NewFormItem("Address", address)
	apt := widget.NewEntry()
	aptItem := widget.NewFormItem("Apt", apt)
	city := widget.NewEntry()
	cityItem := widget.NewFormItem("City", city)
	zip := widget.NewEntry()
	zipItem := widget.NewFormItem("Zip", zip)
	state := widget.NewEntry()
	stateItem := widget.NewFormItem("State", state)
	email := widget.NewEntry()
	emailItem := widget.NewFormItem("Email", email)
	phone := widget.NewEntry()
	phoneItem := widget.NewFormItem("Phone", phone)
	cc := widget.NewEntry()
	ccItem := widget.NewFormItem("CC Number", cc)
	month := widget.NewSelectEntry([]string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"11",
		"12",
	})
	//monthItem := widget.NewFormItem("Month", month)
	year := widget.NewSelectEntry([]string{
		"2021",
		"2022",
		"2023",
		"2024",
		"2025",
		"2026",
		"2027",
		"2028",
		"2029",
		"2030",
	})
	//yearItem := widget.NewFormItem("Year", year)
	mycontainer := container.NewVBox(month, year)
	containerItem := widget.NewFormItem("Month/Year", mycontainer)
	cvv := widget.NewEntry()
	cvvItem := widget.NewFormItem("CVV", cvv)
	saveButton := widget.NewButton("Save", func() {

	})

	profileForm := widget.NewForm(
		titleItem,
		firstnameItem,
		lastnameItem,
		addressItem,
		aptItem,
		cityItem,
		zipItem,
		stateItem,
		emailItem,
		phoneItem,
		ccItem,
		containerItem,
		cvvItem,
	)

	//add pop up
	pidContainer := container.NewVBox(profileSelect,
		loadProfile,
		deleteProfile,
		fSeparator,
		profileForm,
		saveButton,
	)
	profilePopup := widget.NewPopUp(pidContainer, win.Canvas())
	addButton := widget.NewButton("Profiles", func() {
		profilePopup.Show()
		profilePopup.Resize(fyne.Size{
			Width: 400,
			Height: 550,
		})
	})
	return addButton
}
