package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"os"
)

const (
	URL = "https://www.marmiton.org/recettes/recette-hasard.aspx?v=2"
)

func Crawl(fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("error creating file %s: %v", fileName, err)
		return
	}
	defer f.Close()
	var ingredients = make(map[string]int)

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.marmiton.org", "marmiton.org", "https://www.marmiton.org"),
	)
	c.AllowURLRevisit = true

	c.OnHTML(".RCP__sc-vgpd2s-0", func(element *colly.HTMLElement) {
		element.ForEach("img", func(_ int, element *colly.HTMLElement) {
			ingredients[element.Attr("alt")] += 1
		})
		if len(ingredients) > 800 {
			return
		}
		fmt.Println(len(ingredients))
		err := element.Request.Visit(URL)
		if err != nil {
			log.Fatalf("error visiting link: %v", err)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.Visit(URL)

	enc := json.NewEncoder(f)
	err = enc.Encode(ingredients)
	if err != nil {
		log.Fatalf("error encoding data: %v", err)
		return
	}
}
