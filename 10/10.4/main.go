package main // Пакеты (импорт/экспорт) внешний!
import (
	"fmt"
	//	"account" так работать не будет
	"lesson10/account" // 10 - название модуля из go.mod, в моем случае это lesson10
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
	fmt.Println(*myAcount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
