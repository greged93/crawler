package main

import (
	"flag"
	"log"
)

var (
	fileName = ""
	parse    = true
)

func main() {
	flag.StringVar(&fileName, "filename", "", "ingredient json file")
	flag.BoolVar(&parse, "parse", true, "parse results if true, otherwise crawl")
	flag.Parse()
	if fileName == "" {
		log.Fatal("missing filename")
		return
	}
	if parse {
		Parse()
	} else {
		Crawl(fileName)
	}
}
