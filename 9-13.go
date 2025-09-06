package main // Композиция

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string // это список полей, название поля + тип поля
	password string
	url      string
}

type accountWithTimeStamp struct {
	account
	createdAt time.Time
	updatedAt time.Time
}

// это пример того что мы хотим расширить стректуру, но создаем новую на самом деле
// type accountWithTimeStamp struct {
// 	login     string
// 	password  string
// 	url       string
// 	createdAt time.Time
// 	updatedAt time.Time
// }

func (acc *account) outputPassword() { // методы прописываются сразу после объявления структуры
	fmt.Println(*acc)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("Invalid_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString) // без http, https будет ошибка
	if err != nil {
		return nil, errors.New("Invalid_URL")
	}
	newAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}
	if password == "" {
		newAcc.generatePassword(12)
		//		newAcc.account.generatePassword(12) // запись аналогичная предыдущей
	}
	return newAcc, nil
}

// func newAccount(login, password, urlString string) (*account, error) {
// 	if login == "" {
// 		return nil, errors.New("Invalid_LOGIN")
// 	}
// 	_, err := url.ParseRequestURI(urlString) // без http, https будет ошибка
// 	if err != nil {
// 		return nil, errors.New("Invalid_URL")
// 	}
// 	newAcc := &account{
// 		login:    login,
// 		password: password,
// 		url:      urlString,
// 	}
// 	if password == "" {
// 		newAcc.generatePassword(12)
// 	}
// 	return newAcc, nil
// }

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-!*")

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")
	myAcount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAcount.outputPassword() // так вызывается метод, передавать ничего не нужно
	fmt.Println(*myAcount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
