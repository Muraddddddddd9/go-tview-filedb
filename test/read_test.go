package test

import (
	"os"
	"testing"

	"go-tview-filedb/fileio"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	testData := "Name: John, Lastname: Doe, Zodiac: Aries, Birthday: 01/01/2000"
	err := os.WriteFile("test_data.txt", []byte(testData), 0644)
	assert.NoError(t, err)

	fileDir := "test_data.txt"
	text := fileio.ReadData(fileDir)

	assert.Equal(t, testData, text, "Содержимое файла не совпадает")

	_ = os.Remove("test_data.txt")
}
