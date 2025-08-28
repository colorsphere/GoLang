package main // массивы
import "fmt"

func main() {
	//	var transactions [3]int
	//	transactions = [3]int{5, 10, -7}
	transactions := [3]int{1, 2, 3}
	banks := [2]string{}

	fmt.Println(transactions[1])
	banks[0] = "A"
	fmt.Println(banks)
}
