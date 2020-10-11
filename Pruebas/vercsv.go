package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {
	fmt.Println("1) Ver pymes.csv\n2) Ver retail.csv\n3) Ver results.csv")
	var opcion string
	fmt.Scanf("%s", &opcion)
	if opcion == "1" {
		records := readCsvFile("../archivos/pymes.csv")
		for i, producto := range records {
			fmt.Println(i, producto)
		}
	} else if opcion == "2" {
		records := readCsvFile("../archivos/retail.csv")
		for i, producto := range records {
			fmt.Println(i, producto)
		}

	} else if opcion == "3" {
		records := readCsvFile("../archivos/results.csv")
		for i, producto := range records {
			fmt.Println(i, producto)
		}
	}
}
