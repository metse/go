package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Record struct {
	merchantID string
	name       string
	liveAppID  string
	testAppID  string
}

func main() {
	file, err := os.Open("file.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var stores []Record

	for _, line := range lines {
		row := Record{
			merchantID: line[0],
			name:       line[1],
			liveAppID:  line[2],
			testAppID:  line[3],
		}

		var jsonData []byte
		jsonData, err := json.Marshal(row)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(jsonData))
		stores = append(stores, row)
	}

	// formatted, err := json.MarshalIndent(stores, "", "  ")
	// fmt.Println(formatted)
}
