package fileio

import (
	"os"
)

func ReadData(fileDir string) string {
	data, err := os.ReadFile(fileDir)
	if err != nil {
		return "Error reading file"
	}

	return string(data)
}
