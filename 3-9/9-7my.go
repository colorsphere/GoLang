package main // rand

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string // это список полей, название поля + тип поля
	password string
	url      string
}

func main() {
	n := 0
	fmt.Println(rand.IntN(10))

	// login := promptData("Введите логин: ")
	// password := promptData("Введите пароль: ")
	// url := promptData("Введите URL: ")

	// myAcount := account{
	// 	login:    login,    // данные из переменных заносятся в поля структуры
	// 	password: password, //
	// 	url:      url,
	// }
	// outputPassword(&myAcount)
	fmt.Print("Введите длину пароля: ")
	fmt.Scan(&n)
	fmt.Println(generatePassword(n))
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println(acc) // выводится указатель
	//	acc.login = "!!"
	fmt.Println(acc.login, acc.password, acc.url) // аналогичная запись выглядит fmt.Println((*acc).login, (*acc).password, (*acc).url) // выводятся значения
}

func generatePassword(n int) string {
	chars := []rune("abcdefghijklmn")
	password := ""
	for i := 0; i < n; i++ {
		password = password + string(chars[rand.IntN(len(chars))])
	}
	return password
}
