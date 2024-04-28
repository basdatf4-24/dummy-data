package cmd

import (
	"encoding/csv"
	"os"

	generatedata "github.com/mhmdhaekal/dummydata/generate-data"
)

func GenerateUser() {
	var records [][]string
	file, err := os.Create("out/user.csv")
	defer file.Close()

	if err != nil {
		panic(err)
	}
	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()
	for i := 0; i < 10; i++ {
		user, err := generatedata.GenerateUser()
		if err != nil {
			panic(err)
		}
		row := []string{user.Username, user.Password}
		records = append(records, row)
	}
	csvWriter.WriteAll(records)
}
