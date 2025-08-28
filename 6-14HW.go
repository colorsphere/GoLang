package main

import (
	"fmt"
	"slices"
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
	switch operation {
	case "SUM", "sum":
		for _, value := range numbers {
			result += value
		}
		fmt.Printf("Сумма чисел равна: %v", result)
	case "AVG", "avg":
		for _, value := range numbers {
			result += value
		}
		result = result / float64(len(numbers))
		fmt.Printf("Среднее значение равно: %v", result)
	case "MED", "med":
		slices.Sort(numbers)
		if len(numbers)%2 == 1 {
			numbers = numbers[len(numbers)/2 : len(numbers)/2+1]
			result = float64(numbers[0])
		} else {
			numbers = numbers[len(numbers)/2-1 : len(numbers)/2+1]
			result = float64((numbers[0] + numbers[1]) / 2)
		}
		fmt.Printf("Медиана равна: %v", result)
	}
	return
}
