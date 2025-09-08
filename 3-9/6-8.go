package main // Циклы в массивах
import (
	"fmt"
)

// В цикле спрашиваем ввод транзакций: -10, 10, 40.5
// Добавить каждую в массив транзакций
// Вывести массив

func main() {
	transactions := []float64{} // объявляем пустой слайс без длины
	for {
		summ := scanTransaction()
		if summ != 0 {
			transactions = append(transactions, summ)
		} else {
			break
		}
	}
	fmt.Println(transactions) // это обычный вывод слайса
	//	for index, value := range transactions { // это вариант с индексами
	//		fmt.Println(index, value)
	//	}
	// for _, value := range transactions { // это вариант без индексов
	// 	fmt.Println(value)
	// }
	for index := range transactions { // это вариант только с индексами, указывать value не нужно в принципе
		fmt.Println(index)
	}
}

func scanTransaction() (summ float64) {
	fmt.Print("Введите сумму (n для выхода): ")
	fmt.Scan(&summ)
	return
}
