package main // создание указателя

import "fmt"

func main() {
	a := 5
	pointerA := &a // создание указателя
	fmt.Println(pointerA)
	res := double(a)
	fmt.Println(res)
}

func double(num int) int {
	return num * 2
}
