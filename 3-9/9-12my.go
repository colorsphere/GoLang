package main // упражнение

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct {
	login    string // это список полей, название поля + тип поля
	password string
	url      string
}

func (acc *account) outputPassword() { // методы прописываются сразу после объявления структуры
	fmt.Println(acc.login, acc.password, acc.url)
}

func generatePassword(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(res)
}

// 1. Если логина нет - ошибка
// 2. Если нет пароля - генерим

func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("Пустой логин")
	}
	if password == "" {
		password = generatePassword(10)
	}
	_, err := url.ParseRequestURI(urlString) // без http, https будет ошибка
	if err != nil {
		return nil, errors.New("Invalid_URL")
	}
	return &account{ // объявляем аккаунт в функции, валидацию можно производить тут же до внесения данных
		url:      urlString,
		login:    login,
		password: password,
	}, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-!*")

func main() {
	//	n := 12
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")
	myAcount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//	myAcount.generatePassword(n)
	//	myAcount.outputPassword() // так вызывается метод, передавать ничего не нужно
	fmt.Println(*myAcount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
