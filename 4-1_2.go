package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight, userKg float64
	fmt.Println("___Калькулятор индекса массы тела___")
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Print("Ваш индекс массы тела: ")
	fmt.Println(IMT)
}
