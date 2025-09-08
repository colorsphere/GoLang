package main //  Error
import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2 //глобальная константа

func main() {
	fmt.Println("___Калькулятор индекса массы тела___")
	for {
		userHeight, userKg := getUserInput()
		IMT, _ := calculateIMT(userKg, userHeight) //  Подчеркивание говорит о том что из ф-ии что-то передается, но это не нужно обрабатывать
		// IMT, err := calculateIMT(userKg, userHeight)
		// if err != nil {
		// 	fmt.Println("Неверно заданы параметры")
		// 	continue
		// }
		outputResult(IMT)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}

	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.2f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит веса")
	case imt < 18.5:
		fmt.Println("У вас дефицит веса")
	case imt < 25:
		fmt.Println("Норма")
	case imt < 30:
		fmt.Println("У вас избыточная масса")
	default:
		fmt.Println("У вас ожирение")
	}
}

func calculateIMT(userKg, userHeight float64) (float64, error) { //альтернативный возврат функции
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower) //не создание, а присвоение
	return IMT, nil                                    // специальное значение, обозначающее отсутствие чего-либо
}

func getUserInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Print("Будем еще считать? (y/n)")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}
