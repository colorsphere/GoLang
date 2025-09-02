package main // reverse меняем порядок в массиве на обратный
import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	reverse(&a)
}

func reverse(a *[5]int) {
	s := *a
	for i := 0; i < len(s); i++ {
		(*a)[i] = (s)[len(s)-i-1]
	}
	fmt.Println(*a)
}
