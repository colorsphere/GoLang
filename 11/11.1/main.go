package main // Пакеты (импорт/экспорт)
import (
	"fmt"
	//	"account" так работать не будет
	"11.1/account" // название модуля из go.mod
	"11.1/files"
)

func main() {
	files.WriteFile("Привет!!!", "file.txt")
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
