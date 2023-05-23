package csv

import (
	"encoding/csv"
	"os"
	"sort"
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

	var keys []string
	for k := range names {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, name := range keys {
		year := names[name]
		err := writer.Write([]string{name, year})
		if err != nil {
			return err
		}
	}
	return nil
}
