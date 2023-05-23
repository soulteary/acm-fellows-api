package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/serenize/snaker"
)

type Fellow struct {
	Name string
	Year string
}

func main() {
	// origin page addr: https://awards.acm.org/fellows/award-winners
	quotePage := "https://awards.acm.org/fellows/award-recipients"
	resp, err := http.Get(quotePage)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	names := make(map[string]string)

	doc.Find("tr[role=row]").Each(func(i int, row *goquery.Selection) {
		name := row.Find("td:nth-child(1) a").Text()
		year := row.Find("td[role=rowheader]").Text()
		if year == "" {
			year = "0"
		}
		name = processName(name)
		names[name] = year
	})

	err = writeCSV(names)
	if err != nil {
		log.Fatal(err)
	}
}

func processName(name string) string {
	name = strings.TrimSpace(name)
	name = snaker.SnakeToCamel(name)
	return name
}

func writeCSV(names map[string]string) error {
	file, err := os.Create("acm-fellows.csv")
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
