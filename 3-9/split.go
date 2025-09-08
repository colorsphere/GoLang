package main

import (
	"fmt"
	"strings"
)

func main() {
	// Исходная строка, которую мы будем разделять
	text := "яблоко,банан,вишня"

	// Используем функцию Split для разделения строки по запятой
	fruits := strings.Split(text, ",") // это слайс!!!

	// Выводим результат
	fmt.Println(fruits)
}
