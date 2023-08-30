package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Создаем читателя CSV
	// чтоб читал файл
	reader := csv.NewReader(file)

	// -------------------------------------------------
	// Читаем все строки
	// records, err := reader.ReadAll()
	// if err != nil {
	// 	panic(err)
	// }

	// Выводим считанные данные
	// for _, record := range records {
	// 	fmt.Println(record)
	// }
	// -------------------------------------------------

	for i := 0; i < 3; i++ {
		record, err := reader.Read()
		if err != nil {
			panic(err)
		}
		fmt.Println(record)
	}

}
