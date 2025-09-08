package main

import (
	"errors"
	"fmt"
)

func main() {
	currencyMap := map[string]map[string]float64{
		"USD": map[string]float64{"RUB": 80.68, "EUR": 0.8541},
		"EUR": map[string]float64{"RUB": 94.46, "USD": 1.1682},
		"RUB": map[string]float64{"USD": 1 / 80.68, "EUR": 1 / 94.46}}
	currencyMapPointer := &currencyMap

	fmt.Println("___ __ _ Калькулятор валют _ __ ___")
	fmt.Print("Введите исходную валюту (USD/EUR/RUB): ")
	currencyIn := readCurrency()
	fmt.Print("Введите сумму: ")
	summ, err := readSumm()
	for err != nil {
		fmt.Print("Ошибка ввода суммы, повторите: ")
		summ, err = readSumm()
	}
	fmt.Println("Вы ввели сумму, равную", summ, currencyIn)
	fmt.Print("Введите целевую валюту (USD/EUR/RUB): ")
	currencyOut := readCurrency()
	for currencyOut == currencyIn {
		fmt.Print("Целевая валюта не должна совпадать с исходной (USD/EUR/RUB): ")
		currencyOut = readCurrency()
	}
	fmt.Printf("Переводим %v %v в %v: ", summ, currencyIn, currencyOut)
	fmt.Printf("%.3f", calculation(currencyMapPointer, summ, currencyIn, currencyOut))
}

func readCurrency() (currency string) {
	for currency != "USD" && currency != "EUR" && currency != "RUB" {
		fmt.Scan(&currency)
		if currency != "USD" && currency != "EUR" && currency != "RUB" {
			fmt.Print("Неверный ввод, повторите (USD/EUR/RUB):")
		}
	}
	return
}
func readSumm() (float64, error) {
	var summ float64
	fmt.Scan(&summ)
	if summ <= 0 {
		return 0, errors.New("Некорректно введена сумма")
	}
	return summ, nil
}

func calculation(calcMap *map[string]map[string]float64, summ float64, currency1 string, currency2 string) (result float64) {
	result = (*calcMap)[currency1][currency2] * summ
	return
}
