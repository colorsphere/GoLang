package main // слайсы, массивы: длина и ёмкость
import (
	"fmt"
	"reflect"
)

func main() {
	// transactions := [5]int{1, 2, 3, 4, 5}
	// transactionsNew := transactions  // массив копируется
	// transactionsNew[0] = 30

	// fmt.Println(transactions)
	// fmt.Println(transactionsNew)

	// разные примеры !!!

	// transactions := [5]int{1, 2, 3, 4, 5}
	// transactionsPartial := transactions[1:] // слайс только ссылка на исходный массив !!!
	// transactionsNewPartial := transactionsPartial[:1]
	// transactionsNewPartial[0] = 30

	// fmt.Println(transactions)
	// fmt.Println(transactionsPartial)
	// fmt.Println(transactionsNewPartial)
	// fmt.Println(len(transactions))
	// fmt.Println(len(transactionsPartial), cap(transactionsPartial))
	// fmt.Println(len(transactionsNewPartial), cap(transactionsNewPartial))

	// разные примеры!!!

	transactions := [6]int{1, 2, 3, 4, 5, 6}
	transactionsPartial := transactions[1:5] // слайс только ссылка на исходный массив !!!
	transactionsNewPartial := transactionsPartial[:1]
	transactionsNewPartial[0] = 30

	transactionsNewPartial = transactionsNewPartial[0:4] // только со слайсами так можно

	fmt.Println(transactions)
	fmt.Println(reflect.TypeOf(transactions).Kind()) // выясняем тип - массив или слайс
	fmt.Println(transactionsPartial)
	fmt.Println(transactionsNewPartial)
	fmt.Println(len(transactions))
	fmt.Println(len(transactionsPartial), cap(transactionsPartial))
	fmt.Println(len(transactionsNewPartial), cap(transactionsNewPartial))
}
