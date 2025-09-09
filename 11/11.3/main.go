package main // Stack frame, defer. 11.1, 11.2, 11.3, 11.4

import (
	"fmt"
	//	"account" так работать не будет
	// название модуля из go.mod
	"11/files"
)

func main() {
	files.WriteFile("Привет!", "file.txt") // передаем что положить в какой файл
	files.ReadFile()
	// login := promptData("Введите логин: ")
	// password := promptData("Введите пароль: ")
	// url := promptData("Введите URL: ")
	// myAcount, err := account.NewAccountWithTimeStamp(login, password, url) // не будет работать, пока название функции не будет написано с БОЛЬШОЙ буквы
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// //	a := account.Account{} // аналогично передается структура
	// myAcount.OutputPassword() // здесь аналогично

	// fmt.Println(*myAcount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
