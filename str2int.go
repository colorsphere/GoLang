package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	number, err := strconv.Atoi(str) // Atoi преобразует строку в int
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(number) // Вывод будет: 123
	}
}
