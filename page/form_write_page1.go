package page

import (
	"github.com/rivo/tview"
	"go-tview-filedb/fileio"
)

func FormApplication(app *tview.Application, pages *tview.Pages, pageNow *int, zodiacArr []string, fileDir string) *tview.Pages {
	inputFirstName := tview.NewInputField().SetLabel("Name: ").SetFieldWidth(20)
	inputLastName := tview.NewInputField().SetLabel("Lastname: ").SetFieldWidth(20)
	selectZodiacSign := tview.NewDropDown().SetLabel("Zodiac:").SetFieldWidth(20).SetOptions(zodiacArr, nil)
	inputBirthdayDay := tview.NewInputField().SetLabel("Birthday day: ").SetFieldWidth(20)
	inputBirthdayMonth := tview.NewInputField().SetLabel("Birthday month: ").SetFieldWidth(20)
	inputBirthdayYear := tview.NewInputField().SetLabel("Birthday year: ").SetFieldWidth(20)

	output := tview.NewTextView().SetDynamicColors(true).SetText("Welcome! Enter the details and click Submit.\n")

	form := tview.NewForm().
		AddFormItem(inputFirstName).
		AddFormItem(inputLastName).
		AddFormItem(selectZodiacSign).
		AddFormItem(inputBirthdayDay).
		AddFormItem(inputBirthdayMonth).
		AddFormItem(inputBirthdayYear).
		AddButton("Submit", func() {
			firstName := inputFirstName.GetText()
			lastName := inputLastName.GetText()
			zodiacSignIndex, _ := selectZodiacSign.GetCurrentOption()

			if zodiacSignIndex == -1 {
				output.SetText("You have not entered all the data.")
				return
			}

			birthdayDay := inputBirthdayDay.GetText()
			birthdayMonth := inputBirthdayMonth.GetText()
			birthdayYear := inputBirthdayYear.GetText()

			birthday := []string{birthdayDay, birthdayMonth, birthdayYear}
			res := fileio.AddData(firstName, lastName, zodiacArr[zodiacSignIndex], birthday, fileDir)
			output.SetText(res)
		}).
		AddButton("Next page", func() {
			*pageNow++
			pages.SwitchToPage("Read page")
		}).
		AddButton("Quit", func() {
			app.Stop()
		})

	return pages.AddPage("Add data", tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true).
		AddItem(output, 1, 1, false), true, *pageNow == 0)
}
