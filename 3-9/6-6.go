package main // Упражнение
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
	fmt.Println(transactions)
}

func scanTransaction() (summ float64) {
	fmt.Print("Введите сумму (n для выхода): ")
	fmt.Scan(&summ)
	return
}
