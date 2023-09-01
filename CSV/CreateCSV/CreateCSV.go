package main

import (
	"encoding/csv"
	"os"
)

func main() {
	// os.Create создаёт файл
	file, err := os.Create("example.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush() // Закрытие файла будет выполнено при выходе из функции

	header := []string{"Name", "Surname", "Age"}
	err = writer.Write(header)
	if err != nil {
		panic(err)
	}

	data1 := []string{"Sanjar", "Usmonov", "29"}
	err = writer.Write(data1)
	if err != nil {
		panic(err)
	}

	data2 := []string{"Somebody", "Somebodiev", "00"}
	err = writer.Write(data2)
	if err != nil {
		panic(err)
	}

	writer.Flush() // Сброс буфера и запись данных при выходе из функции

	if err := writer.Error(); err != nil {
		panic(err)
	}
}

/*
+----------------------------------------------------+
|           example.csv                             |
+------------+--------------+-----------------------+
|    Name    |   Surname    |         Age           |
+------------+--------------+-----------------------+
|   Sanjar   |   Usmonov    |          29           |
|  Somebody  |  Somebodiev  |          00           |
+------------+--------------+-----------------------+

*/
