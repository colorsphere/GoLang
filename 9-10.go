package main // constructor и валидация

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string // это список полей, название поля + тип поля
	password string
	url      string
}

func (acc *account) outputPassword() { // методы прописываются сразу после объявления структуры
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccount(login, password, url string) *account {
	return &account{ // объявляем аккаунт в функции, валидацию можно производить тут же до внесения данных
		url:      url,
		login:    login,
		password: password,
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-!*")

func main() {
	n := 12
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAcount := *newAccount(login, password, url)
	myAcount.generatePassword(n)
	myAcount.outputPassword() // так вызывается метод, передавать ничего не нужно
	fmt.Println(myAcount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}
