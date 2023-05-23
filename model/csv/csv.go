package csv

import (
	"encoding/csv"
	"os"
)

func Save(filename string, names map[string]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"name", "year"})
	if err != nil {
		return err
	}

	for name, year := range names {
		err := writer.Write([]string{name, year})
		if err != nil {
			return err
		}
	}
	return nil
}
