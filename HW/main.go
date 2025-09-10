package main // 11.6 JSON

import (
	"fmt"
	//	"account" так работать не будет
	// название модуля из go.mod
	bin "11/storage"
)

func main() {
	vault := bin.NewVault()
	createBin(vault)
}

func createBin(vault *bin.Vault) {
	login := promptData("Введите имя: ")
	password := promptData("Введите фамилию: ")
	url := promptData("Введите дату: ")
	myAcount, err := bin.NewBin(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault.AddBin(*myAcount)

}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
