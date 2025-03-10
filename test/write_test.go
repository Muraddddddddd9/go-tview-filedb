package test

import (
	"os"
	"testing"

	"go-tview-filedb/fileio"
	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	fileDir := "test_data.txt"

	_ = os.Remove(fileDir)

	result := fileio.AddData("John", "Doe", "Aries", []string{"01", "01", "2000"}, fileDir)
	expected := "Added: John Doe, Zodiac: Aries, Birthday: 01/01/2000."
	assert.Equal(t, expected, result, "Ожидаемый и фактический результаты не совпадают")

	data, err := os.ReadFile("test_data.txt")
	assert.NoError(t, err, "Ошибка чтения файла")
	assert.Contains(t, string(data), "Name: John, Lastname: Doe, Zodiac: Aries, Birthday: 01/01/2000", "Данные в файле не совпадают")

	assert.Equal(t, "You have not entered all the data.", fileio.AddData("", "Doe", "Aries", []string{"01", "01", "2000"}, fileDir))
	assert.Equal(t, "Incorrect input", fileio.AddData("John", "Doe", "Aries", []string{"01", "xx", "2000"}, fileDir))
	assert.Equal(t, "Incorrect date entry", fileio.AddData("John", "Doe", "Aries", []string{"32", "01", "2000"}, fileDir))
	assert.Equal(t, "Incorrect month entry", fileio.AddData("John", "Doe", "Aries", []string{"01", "13", "2000"}, fileDir))

	_ = os.Remove("test_data.txt")
}
