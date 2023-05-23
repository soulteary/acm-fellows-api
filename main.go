package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Fellow struct {
	Name   string
	Year   string
	URL    string
	Region string
}

func GetListFromPage(link string) (fellows []Fellow, error error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("table tbody tr[role=row]").Each(func(_ int, row *goquery.Selection) {
		cells := row.Find("td")
		if cells.Length() != 5 {
			return
		}

		fellow := Fellow{
			Name:   strings.TrimSpace(cells.Eq(0).Find("a").Text()),
			URL:    cells.Eq(0).Find("a").AttrOr("href", ""),
			Year:   strings.TrimSpace(cells.Eq(2).Text()),
			Region: strings.TrimSpace(cells.Eq(3).Text()),
		}

		if fellow.Year == "" {
			fellow.Year = "0"
		}

		fellows = append(fellows, fellow)
	})

	return fellows, nil
}

func main() {
	// origin page addr: https://awards.acm.org/fellows/award-winners
	fellows, err := GetListFromPage("https://awards.acm.org/fellows/award-recipients")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total fellows:", len(fellows))
	for _, fellow := range fellows {
		fmt.Println(fellow)
	}
}
