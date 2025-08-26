package main // Циклы
import (
	"fmt"
	"math"
)

const IMTPower = 2 //глобальная константа

func main() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Printf("%d\n", i)
	//	}
	i := 0
	for i < 10 {
		fmt.Printf("%d\n", i)
		i++
	}

	fmt.Println("___Калькулятор индекса массы тела___")
	userHeight, userKg := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
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
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf(`Ваш индекс массы тела: %.2f`, imt)
	fmt.Print(result)
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
