package main

import "fmt"

func main() {
	var a = 5
	var p = &a
	*p = 10
	fmt.Println(a)
}
