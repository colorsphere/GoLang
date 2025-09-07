package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "1,3,5,6,8,42"

	strs := strings.Split(str, ",") // это СЛАЙС
	var ints []int
	for _, s := range strs {
		num, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, num)
		}
	}

	fmt.Println(ints)
}
