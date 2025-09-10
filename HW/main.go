package main

import (
	"11HW/bins"
	"11HW/files"
	"fmt"
)

func main() {
	fmt.Println("1. Создаем структуру в bin и сохраняем её в формате json")
	fmt.Println("2. Читаем файл, проверяем что это json файл")
	data := bins.Tojson{
		{Name: "Mike", CerName: "Morozov"},
		{Name: "Ivan", CerName: "Ivanov"},
	}
	fmt.Println(data) // потом удалить
	files.ExistsFile("test.test")
}
