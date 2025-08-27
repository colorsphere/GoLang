package main // динамические массивы (слайсы)
import (
	"fmt"
	"reflect"
	"slices"
)

func main() {
	transactions := []int{40, 20, 35} // объявляем массив-слайс без длины

	//	slice := transactions[:2]
	//	slice[0] = 100
	//	transactions = append(transactions, 100) //  так добавляется элемент в динамический слайс
	//	fmt.Println(transactions)
	//	fmt.Println(newTransactions)
	//	fmt.Println(slice)

	temp := transactions //  создается НОВЫЙ слайс, а не ссылка на старый
	transactions = append(transactions, 100)
	temp = append(temp, 90)
	fmt.Println(transactions)
	fmt.Println(temp)

	fmt.Println(reflect.TypeOf(transactions).Kind())
	fmt.Println(reflect.TypeOf(temp).Kind())

	slices.Sort(transactions)
	fmt.Println(transactions)
}
