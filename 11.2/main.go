package main // Stack frame, defer. 11.1, 11.2, 11.3

import (
	"fmt"
	//	"account" так работать не будет
	"11.2/account" // название модуля из go.mod
	"11.2/files"
)

func main() {
	files.WriteFile("Привет!!!", "file.txt") // передаем что положить в какой файл
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")
	myAcount, err := account.NewAccountWithTimeStamp(login, password, url) // не будет работать, пока название функции не будет написано с БОЛЬШОЙ буквы
	if err != nil {
		fmt.Println(err)
		return
	}
	//	a := account.Account{} // аналогично передается структура
	myAcount.OutputPassword() // здесь аналогично

	fmt.Println(*myAcount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
