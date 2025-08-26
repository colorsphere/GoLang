package main // Циклы
import (
	"fmt"
	"math"
)

const IMTPower = 2 //глобальная константа

func main() {
	fmt.Println("___Калькулятор индекса массы тела___")
	for check := 1; check != 0; {
		check = 1
		userHeight, userKg := getUserInput()
		IMT := calculateIMT(userKg, userHeight)
		outputResult(IMT)
		switch {
		case IMT < 16:
			fmt.Println("У вас сильный дефицит веса")
		case IMT < 18.5:
			fmt.Println("У вас дефицит веса")
		case IMT < 25:
			fmt.Println("Норма")
		case IMT < 30:
			fmt.Println("У вас избыточная масса")
		default:
			fmt.Println("У вас ожирение")
		}
		fmt.Print("Будем еще считать? (0 - выход)")
		fmt.Scan(&check)
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.2f", imt)
	fmt.Println(result)
}

func calculateIMT(userKg, userHeight float64) (IMT float64) { //альтернативный возврат функции
	IMT = userKg / math.Pow(userHeight/100, IMTPower) //не создание, а присвоение
	return
}

func getUserInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}
