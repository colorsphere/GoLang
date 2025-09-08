package main // Пакет files для работы с файлами
import (
	"fmt"
	//	"account" так работать не будет
	"lesson10/account" // название модуля из go.mod, в моем случае это lesson10
	"lesson10/files"
)

func main() {
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
	files.WriteFile()
	fmt.Println(*myAcount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
