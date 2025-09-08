package main // struct пример объявления

import "fmt"

type account struct {
	login    string // это список полей, название поля + тип поля
	password string
	url      string
}

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAcount := account{ //
		login:    login,    // данные из переменных заносятся в поля структуры
		password: password, //
		url:      url,
	}
	outputPassword(myAcount)

}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc account) {
	fmt.Println(acc)
	fmt.Println(acc.login, acc.password, acc.url)
}
