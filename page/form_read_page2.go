package page

import (
	"fmt"

	"github.com/rivo/tview"
	"go-tview-filedb/fileio"
)

func ReadApplication(app *tview.Application, pages *tview.Pages, pageNow *int, fileDir string) *tview.Pages {
	output := tview.NewTextView().SetDynamicColors(true)
	buttonBack := tview.NewButton("Back page").SetSelectedFunc(func() {
		*pageNow--
		pages.SwitchToPage("Add data")
	})

	pages.AddPage("Read page", tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(buttonBack, 1, 1, true).
		AddItem(output, 0, 1, false), true, *pageNow == 1)

	pages.SetChangedFunc(func() {
		if page, _ := pages.GetFrontPage(); page == "Read page" {
			resData := fileio.ReadData(fileDir)
			output.SetText(fmt.Sprintf("All data from the file:\n%s", resData))
		}
	})

	return pages
}
