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

type Building struct {
	UUID        string
	State       string
	Type        string
	SubType     string
	Standard    string
	Description string
	Cost        float64
}

func main() {
	csvFile, err := os.Open("raw_cub.csv")
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

	statesUUIDMap := make(map[string]string)

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

		if val, ok := statesUUIDMap[mergedState]; ok {
			building.UUID = val
		} else {
			building.UUID = uuid.New().String()
			statesUUIDMap[mergedState] = building.UUID
		}

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

	fmt.Println(string(jsonData))

	jsonFile, err := os.Create("raw_data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}
