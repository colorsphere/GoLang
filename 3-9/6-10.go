package main // make
import (
	"fmt"
)

// В цикле спрашиваем ввод транзакций: -10, 10, 40.5
// Добавить каждую в массив транзакций
// Вывести массив

func main() {
	// tr := make([]string, 2) // создаётся пустое значениt - 2шт., т.к. при объявлении строковой переменной будут пустые значения
	// tr = append(tr, "1")
	// tr = append(tr, "2")
	// fmt.Println(tr)

	// tr := make([]string, 2)
	// tr[0] = "1"
	// tr[1] = "2"
	// fmt.Println(tr)

	tr := make([]string, 0, 2) // пустой элемент не создается !!!
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "1")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "2")
	fmt.Println(cap(tr))
	tr = append(tr, "3")
	fmt.Println(cap(tr))
	tr = append(tr, "4")
	fmt.Println(cap(tr))
	tr = append(tr, "5")
	fmt.Println(cap(tr))
	fmt.Println(tr)

	transactions := []float64{} // объявляем пустой слайс без длины
	for {
		summ := scanTransaction()
		if summ != 0 {
			transactions = append(transactions, summ)
		} else {
			break
		}
	}
	fmt.Printf("Ваш баланс: %.2f", calculateBalance(transactions))

}

func scanTransaction() (summ float64) {
	fmt.Print("Введите сумму (n для выхода): ")
	fmt.Scan(&summ)
	return
}

func calculateBalance(transactions []float64) (balance float64) {
	for _, value := range transactions {
		balance += value
	}
	return
}
