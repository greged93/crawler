package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	fileName = ""
)

func main() {
	flag.StringVar(&fileName, "filename", "", "ingredient json file")
	flag.Parse()
	if fileName == "" {
		log.Fatal("missing filename")
		return
	}
	//Crawl(fileName)
	var ingredients = make(map[string]int)
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(&ingredients)
	if err != nil {
		log.Fatalf("error decoding file: %v", err)
		return
	}
	fmt.Printf("found %d ingredients \n", len(ingredients))
	var frequentlyUsed []string
	for ing, rep := range ingredients {
		if rep > 3 {
			frequentlyUsed = append(frequentlyUsed, ing)
		}
	}
	fmt.Println("got", len(frequentlyUsed), "ingredients")
}
