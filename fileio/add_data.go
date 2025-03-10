package fileio

import (
	"fmt"
	"os"
	"strconv"
)

type ZNAK struct {
	FirstName  string
	LastName   string
	ZodiacSign string
	Birthday   []string
}

func AddData(firstName, lastName, zodiacSign string, birthday []string, fileDir string) string {
	if firstName == "" || lastName == "" || zodiacSign == "" || len(birthday) != 3 {
		return "You have not entered all the data."
	}

	for i, v := range birthday {
		data, err := strconv.Atoi(v)
		if err != nil {
			return "Incorrect input"
		}

		switch i {
		case 0:
			if data > 31 || data <= 0 {
				return "Incorrect date entry"
			}
		case 1:
			if data > 12 || data <= 0 {
				return "Incorrect month entry"
			}
		}
	}

	data := ZNAK{
		FirstName:  firstName,
		LastName:   lastName,
		ZodiacSign: zodiacSign,
		Birthday:   birthday,
	}

	file, err := os.OpenFile(fileDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return "Error opening file"
	}
	defer file.Close()

	msg := fmt.Sprintf("Name: %s, Lastname: %s, Zodiac: %s, Birthday: %s/%s/%s\n",
		data.FirstName, data.LastName, data.ZodiacSign, data.Birthday[0], data.Birthday[1], data.Birthday[2])

	_, err = file.WriteString(msg)
	if err != nil {
		return "Error writing to file"
	}

	return fmt.Sprintf("Added: %s %s, Zodiac: %s, Birthday: %s/%s/%s.", firstName, lastName, zodiacSign, birthday[0], birthday[1], birthday[2])
}
