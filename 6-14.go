package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	calculate(readString(), readOperation())
}

func readString() (numbers []float64) {
	var str string
	fmt.Print("Введите последовательность чисел через запятую: ")
	fmt.Scan(&str)
	strs := strings.Split(str, ",")
	for _, value := range strs {
		number, err := strconv.Atoi(value)
		if err == nil {
			numbers = append(numbers, float64(number))
		}
	}
	return
}

func readOperation() (operation string) {
	fmt.Print("Введите операцию (AVG, SUM, MED):")
	for operation != "AVG" && operation != "avg" && operation != "SUM" && operation != "sum" && operation != "MED" && operation != "med" {
		fmt.Scan(&operation)
		if operation != "AVG" && operation != "avg" && operation != "SUM" && operation != "sum" && operation != "MED" && operation != "med" {
			fmt.Print("Неверный ввод, повторите (AVG, SUM, MED):")
		}
	}
	return
}

func calculate(numbers []float64, operation string) (result float64) {
	fmt.Println(numbers)
	fmt.Println(operation) // удалить
	switch operation {
	case "SUM", "sum":
		fmt.Print("Сумма чисел равна: ")
	case "AVG", "avg":
		fmt.Print("Среднее значение равно: ")
	case "MED", "med":
		fmt.Print("Медиана равна: ")
		numbers := 
	}
	return
}
