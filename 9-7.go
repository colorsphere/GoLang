package main // rand

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string // это список полей, название поля + тип поля
	password string
	url      string
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-!*")

func main() {
	n := 0
	//	fmt.Println(rand.IntN(10)) // пример генерации случайного числа
	fmt.Print("Введите длину пароля: ")
	fmt.Scan(&n)
	fmt.Println(generatePassword(n))
}

func generatePassword(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(res)
}
