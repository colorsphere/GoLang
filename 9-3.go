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

	// account1 := account{
	// 	login, // порядок должен быть как изначально объявлено
	// 	password,
	// 	url, // после последнего поля ОБЯЗАТЕЛЬНО запятая
	// }

	// account1 := account{} // можно создать пустую структуру, в ней будут все поля нулевые

	account1 := account{ // альтернативный синтаксис
		login:    login,    // а здесь порядок может быть любой, т.к. принудительно прописано в какое поле что заносится
		password: password, // так же можно убрать любое поле, в отличие от первого способа
		url:      url,
	}
	outputPassword(login, password, url)

}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(login, password, url string) {
	fmt.Println(login, password, url)
}
