package csv_loader

import (
	"encoding/csv"
	"os"
)

type CSV struct {
	Photo string 		  `json:"photo"`
	Producer string 	  `json:"producer"`
	Denomination string   `json:"denomination"`
	Price string 		  `json:"price"`
	Characteristic string `json:"characteristic"`
}

var Data []*CSV

func ParseCSV(filename string) ([]*CSV, error) {

	lines, err := readCSV(filename)
	if err != nil {
		return []*CSV{}, err
	}

	for _, line := range lines {
		Data = append(Data, &CSV{
			Photo:          line[0],
			Producer:       line[1],
			Denomination:   line[2],
			Price:          line[3],
			Characteristic: line[4],
		})
	}

	return Data, nil
}

func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil  {
		return [][]string{}, err
	}

	return lines, err
}