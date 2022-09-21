package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Parse() {
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
