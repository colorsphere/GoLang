package main // слайсы
import "fmt"

func main() {
	//	var transactions [3]int
	//	transactions = [3]int{5, 10, -7}
	transactions := [5]int{1, 2, 3, 4, 5}
	banks := [2]string{}

	fmt.Println(transactions[1])
	banks[0] = "A"
	fmt.Println(banks)
	partial := transactions[1:]
	fmt.Println(partial)
}
