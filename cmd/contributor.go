package cmd

import (
	"encoding/csv"
	"os"
	"strconv"

	generatedata "github.com/mhmdhaekal/dummydata/generate-data"
)

func GenerateContributor() {
	var records [][]string
	file, err := os.Create("out/contributors.csv")
	defer file.Close()

	if err != nil {
		panic(err)
	}
	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()
	for i := 0; i < 200; i++ {
		contributor, err := generatedata.GenerateContributors()
		if err != nil {
			panic(err)
		}
		row := []string{contributor.Id.String(), contributor.Nama, strconv.Itoa(contributor.JenisKelamin), contributor.Kewarganegaraan}
		records = append(records, row)
	}
	csvWriter.WriteAll(records)
}
