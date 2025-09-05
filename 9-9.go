package main // методы
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
	//	fmt.Println(acc)
	//	acc.login = "!!"
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-!*")

func main() {
	n := 12
	login := promptData("Введите логин: ")
	//	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAcount := account{
		login: login,
		url:   url,
	}
	myAcount.generatePassword(n)
	myAcount.outputPassword() // так вызывается метод, передавать ничего не нужно
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}
