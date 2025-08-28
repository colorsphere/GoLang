package main

import (
	"fmt"
	"strings"
)

func main() {
	calculate(readOperation())

}

func readString() (numbers []float64) {
	var str string
	fmt.Print("Введите последовательность чисел через запятую: ")
	fmt.Scan(&str)
	strings.Split(str, ",")

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

func calculate(operation string) (result float64) {
	fmt.Print(operation)
	return
}
