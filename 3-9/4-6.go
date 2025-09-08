package main //функции

import (
	"fmt"
	"math"
)

func main() {
	var userHeight, userKg float64
	fmt.Print(`___Калькулятор индекса массы тела___
Введите свой рост в сантиметрах: `) // пример мультистрока
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)

	//	fmt.Printf("Ваш индекс массы тела: %v", IMT)
	//	fmt.Print("Ваш индекс массы тела: ", IMT)
	//	fmt.Println(IMT)
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf(`Ваш индекс массы тела: %.2f`, imt)
	fmt.Print(result)
}

func calculateIMT(userKg, userHeight float64) float64 {
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}
