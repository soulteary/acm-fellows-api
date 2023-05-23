package main

import (
	"fmt"
	"log"

	"github.com/soulteary/acm-fellows-api/model/humanname"
	"github.com/soulteary/acm-fellows-api/model/network"
	"github.com/soulteary/acm-fellows-api/model/parse"
)

func main() {
	// origin page addr: https://awards.acm.org/fellows/award-winners
	buf, err := network.GetRemotePage("https://awards.acm.org/fellows/award-recipients")
	if err != nil {
		log.Fatal(err)
	}

	fellows, err := parse.GetListFromPage(buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total fellows:", len(fellows))
	for _, fellow := range fellows {
		fmt.Println(fellow)
		fmt.Println(humanname.Parse(fellow.Name))
	}
}
