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

	outputPassord(login, password, url)

}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassord(login, password, url string) {
	fmt.Println(login, password, url)
}
