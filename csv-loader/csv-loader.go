package csv_loader

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
)

type CSV struct {
	Photo string 		  `json:"photo"`
	Producer string 	  `json:"producer"`
	Denomination string   `json:"denomination"`
	Price string 		  `json:"price"`
	Characteristic string `json:"characteristic"`
}

var csvFile []*CSV

func ParseCSV(reader *csv.Reader) []byte {
	for {
		line , err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Error:", err)
		}

		csvFile = append(csvFile, &CSV{
			Photo:          line[0],
			Producer:       line[1],
			Denomination:   line[2],
			Price:          line[3],
			Characteristic: line[4],
		})
	}

	csvFileToJson, err := json.Marshal(csvFile)
	if err != nil {
		log.Fatal("Error:", err)
	}

	return csvFileToJson
}