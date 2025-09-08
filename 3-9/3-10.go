package main

import "fmt"

func main() {
	const usdToEur = 0.8541
	const usdToRub = 80.02
	eurToRub := usdToRub / usdToEur
	fmt.Println(eurToRub)
}
