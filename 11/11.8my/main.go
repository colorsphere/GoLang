package main // 11.6 JSON

import (
	"fmt"
	//	"account" так работать не будет
	// название модуля из go.mod
	"11/account"
	"11/files"
)

func main() {
	var choice string
Menu:
	for {
		fmt.Println("1. Создать аккаунт")
		fmt.Println("2. Найти аккаунт")
		fmt.Println("3. Удалить аккаунт")
		fmt.Println("4. Выход")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			createAccount()
		case "2":
			findAccount()
		case "3":
			deleteAccount()
		case "4":
			break Menu
		}
	}
}

func findAccount() {

}

func deleteAccount() {

}

func createAccount() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")
	myAcount, err := account.NewAccount(login, password, url) // не будет работать, пока название функции не будет написано с БОЛЬШОЙ буквы
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*myAcount)
	file, err := myAcount.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать даные в JSON")
		return
	}
	files.WriteFile(file, "data.json")
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
