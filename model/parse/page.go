package parse

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/soulteary/acm-fellows-api/model/define"
)

func GetListFromPage(pageData []byte) (fellows []define.Fellow, error error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(pageData))
	if err != nil {
		return nil, err
	}

	doc.Find("table tbody tr[role=row]").Each(func(_ int, row *goquery.Selection) {
		cells := row.Find("td")
		if cells.Length() != 5 {
			return
		}

		fellow := define.Fellow{
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
