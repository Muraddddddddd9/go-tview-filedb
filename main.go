package main

import (
	"github.com/rivo/tview"
	"go-tview-filedb/page"
)

var (
	zodiacArr = []string{
		"Aries", "Taurus", "Gemini", "Cancer", "Leo", "Virgo",
		"Libra", "Scorpio", "Sagittarius", "Capricorn", "Aquarius", "Pisces",
	}
	fileDir = "data.txt"
)

func main() {
	app := tview.NewApplication()
	pages := tview.NewPages()
	pageNow := 0

	_ = page.FormApplication(app, pages, &pageNow, zodiacArr, fileDir)
	_ = page.ReadApplication(app, pages, &pageNow, fileDir)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
