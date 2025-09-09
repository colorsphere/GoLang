package account // JSON (маппинг полей)
import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-!*")

type Account struct {
	Login     string    `json:"login"`    // для экспорта полей их нужно назвать с большой буквы
	Password  string    `json:"password"` // описываем теги так, как они должны быть в JSON - по правилам с маленькой буквы
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) OutputPassword() { // методы прописываются сразу после объявления структуры
	// color.Cyan(acc.Login)
	fmt.Println(*acc)
}

func (acc *Account) ToBytes() ([]byte, error) { // методы прописываются сразу после объявления структуры
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("Invalid_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString) // без http, https будет ошибка
	if err != nil {
		return nil, errors.New("Invalid_URL")
	}
	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     login,
		Password:  password,
		Url:       urlString,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
