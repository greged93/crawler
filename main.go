package main

import (
	"flag"
	"log"
)

var (
	fileName = ""
	parse    = false
	clean    = false
)

func main() {
	flag.StringVar(&fileName, "filename", "", "ingredient json file")
	flag.BoolVar(&parse, "parse", false, "parse results if true, otherwise continue")
	flag.BoolVar(&clean, "clean", false, "clean results if true, otherwise continue")
	flag.Parse()
	if fileName == "" {
		log.Fatal("missing filename")
		return
	}
	if parse {
		Parse()
	} else if clean {
		Clean()
	} else {
		Crawl(fileName)
	}
}
