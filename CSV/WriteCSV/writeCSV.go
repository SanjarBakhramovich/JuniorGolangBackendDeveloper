package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// os.O_WRONLY|os.O_APPEND это флаг Write Only Append only,
	// 0666 это доступ для всех в Операционке
	file, err := os.OpenFile("example.csv", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// что бы закрыть файл после операций
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	newData := []string{"Somebodiya", "Sombodieva", "00"}
	err = writer.Write(newData)
	if err != nil {
		log.Fatal(err)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}

}
