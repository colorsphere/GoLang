package main // 11.6 JSON

import (
	"fmt"
	//	"account" так работать не будет
	// название модуля из go.mod
	"11/account"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("__ _ Менеджер паролей _ __")
	vault := account.NewVault()
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() (variant int) {
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scanln(&variant)
	return
}

func findAccount(vault *account.Vault) {
	url := promptData("Введите URL для поиска: ")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount() {

}

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")
	myAcount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault.AddAccount(*myAcount)

}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
