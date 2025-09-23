package main // 11.6 JSON

import (
	"fmt"
	//	"account" так работать не будет
	// название модуля из go.mod
	"11/account"
	"11/files"
	"11/output"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccount,
	"3": deleteAccount,
}

func main() {
	fmt.Println("__ _ Менеджер паролей _ __")
	vault := account.NewVault(files.NewJsonDb("data.json"))
Menu:
	for {
		variant := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})
		menuFunc := menu[variant] // здесь получаем ФУНКЦИЮ в виде переменной
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault) // здесь вызываем функцию
		// switch variant {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// default:
		// 	break Menu
		// }
	}
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для поиска"})
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для удаления"})
	isDeleted := vault.DeleteAccountsByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		//		color.Red("Не найдено") заменяем на вывод из output
		output.PrintError("Не найдено")
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или Логин")
		//		fmt.Println(err)
		return
	}
	vault.AddAccount(*myAccount)

}

// Функция должна принимать slice любого типа
// Выводит строкой каждый элемент, а последний - Printf добавляя :

func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	//	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
