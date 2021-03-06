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
	fmt.Println("1) Ver pymes.csv\n2) Ver retail.csv\n3) Ver registro.csv\n 4) Ver registroRetail1.csv\n 5) Ver registroRetail2.csv\n 6) ver registroNormal.csv")
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
		records := readCsvFile("../archivos/registro.csv")
		for i, producto := range records {
			fmt.Println(i, producto)
		}
	} else if opcion == "4" {
		records := readCsvFile("../archivos/registroRetail1.csv")
		for i, producto := range records {
			fmt.Println(i, producto)
		}
	} else if opcion == "5" {
		records := readCsvFile("../archivos/registroRetail2.csv")
		for i, producto := range records {
			fmt.Println(i, producto)
		}
	} else if opcion == "6" {
		records := readCsvFile("../archivos/registroNormal.csv")
		for i, producto := range records {
			fmt.Println(i, producto)
		}
	}
}
