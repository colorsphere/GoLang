package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight, userKg float64
	fmt.Print(`___Калькулятор индекса массы тела___
Введите свой рост в сантиметрах: `) // пример мультистрока
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	//	fmt.Printf("Ваш индекс массы тела: %v", IMT)
	//	fmt.Print("Ваш индекс массы тела: ", IMT)
	//	fmt.Println(IMT)
	result := fmt.Sprintf(`Ваш индекс массы тела: %.2f`, IMT)
	fmt.Print(result)
}
