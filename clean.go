package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var mapRunes = map[rune]rune{
	'à': 'a',
	'è': 'e',
	'é': 'e',
	'â': 'a',
	'ô': 'o',
	'î': 'i',
	'ï': 'i',
	'ê': 'e',
	'û': 'u',
	'ñ': 'n',
}

var mapDeleteRunes = map[rune]int{
	'©':      -1,
	'’':      -1,
	'®':      -1,
	'(':      -1,
	')':      -1,
	'"':      -1,
	',':      -1,
	rune(39): -1,
	rune(45): -1,
	rune(32): -1,
}

func Clean() {
	var arr, err = readToArray()
	if err != nil {
		log.Fatalf("error reading file to array: %v", err)
	}
	var cleanArr []string
	for _, s := range arr {
		sLower := strings.ToLower(s)
		sClean := strings.Map(cleanString, sLower)
		cleanArr = append(cleanArr, sClean)
	}
	sort.Slice(cleanArr, func(i, j int) bool {
		return cleanArr[i] < cleanArr[j]
	})
	err = writeToCSV(cleanArr)
	if err != nil {
		log.Fatalf("error writing array to file: %v", err)
	}
}

func cleanString(r rune) rune {
	if v, ok := mapRunes[r]; ok {
		return v
	}
	if v, ok := mapDeleteRunes[r]; ok {
		return rune(v)
	}
	return r
}

func readToArray() ([]string, error) {
	var ingredients = make(map[string]int)
	ext := filepath.Ext(fileName)
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	var ingredientsArray []string
	switch ext {
	case ".json":
		dec := json.NewDecoder(f)
		err = dec.Decode(&ingredients)
		if err != nil {
			return nil, fmt.Errorf("error decoding file: %v", err)
		}
		for k := range ingredients {
			ingredientsArray = append(ingredientsArray, k)
		}
	case ".csv":
		var reader = csv.NewReader(f)
		data, err := reader.ReadAll()
		if err != nil {
			return nil, fmt.Errorf("error opening file: %v", err)
		}
		for _, line := range data {
			ingredientsArray = append(ingredientsArray, line...)
		}
	}
	return ingredientsArray, nil
}

func writeToCSV(arr []string) error {
	outputFile, err := os.Create("data/clean_ingredients.csv")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"ingredient"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range arr {
		if err := writer.Write([]string{r}); err != nil {
			return err
		}
	}
	return nil
}

func writeToJson(arr []string) error {
	b, err := json.Marshal(arr)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("data/clean_ingredients.json", b, 0644)
}
