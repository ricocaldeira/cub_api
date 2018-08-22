package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// Building represents a strcuture who can be built by a Civil Engineer
type Building struct {
	UUID        string  `json:"uuid"`
	State       string  `json:"state"`
	Type        string  `json:"type"`
	SubType     string  `json:"sub_type"`
	Standard    string  `json:"standard"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}

func main() {
	csvFile, err := os.Open("cub.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	reader.Comma = ';'

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	var building Building
	var buildings []Building

	var mergedState string
	var mergedType string
	var mergedSubType string

	for _, each := range csvData {
		if len(each[0]) > 0 {
			mergedState = each[0]
		}
		if len(each[1]) > 0 {
			mergedType = each[1]
		}
		if len(each[2]) > 0 {
			mergedSubType = each[2]
		}

		building.UUID = uuid.New().String()
		building.State = mergedState
		building.Type = mergedType
		building.SubType = mergedSubType
		building.Standard = each[3]
		building.Description = each[4]
		parsedCost := strings.Replace(each[5], ",", ".", -1)
		building.Cost, _ = strconv.ParseFloat(parsedCost, 64)
		buildings = append(buildings, building)
	}

	jsonData, err := json.Marshal(buildings)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonFile, err := os.Create("cub.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}
