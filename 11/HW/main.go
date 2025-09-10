package main // 11.6 JSON

import (
	bin "11/storage"
	"fmt"
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
