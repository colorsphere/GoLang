package main

import (
	"errors"
	"fmt"
)

func main() {
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
	fmt.Printf("%.3f", calculation(summ, currencyIn, currencyOut))
}

func readCurrency() (currency string) {
	for currency != "USD" && currency != "usd" && currency != "EUR" && currency != "eur" && currency != "RUB" && currency != "rub" {
		fmt.Scan(&currency)
		if currency != "USD" && currency != "usd" && currency != "EUR" && currency != "eur" && currency != "RUB" && currency != "rub" {
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

func calculation(summ float64, currency1 string, currency2 string) (result float64) {
	usd := 80.68
	eur := 94.46
	if currency1 == "rub" || currency1 == "RUB" {
		if currency2 == "usd" {
			result = summ / usd
		} else {
			result = summ / eur
		}
	}
	if currency1 == "eur" || currency1 == "EUR" {
		if currency2 == "rub" {
			result = eur * summ
		} else {
			result = eur / usd * summ
		}
	}
	if currency1 == "usd" || currency1 == "USD" {
		if currency2 == "rub" {
			result = usd * summ
		} else {
			result = usd / eur * summ
		}
	}
	return
}
