package account // JSON
import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"reflect"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	//	login    string `demo:"abc"` // так описываются теги, ключ и значение
	login string `json:"login" xml:"test"` // тегов может быть несколько
	//	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	Account
	createdAt time.Time
	updatedAt time.Time
}

func (acc *Account) OutputPassword() { // методы прописываются сразу после объявления структуры
	color.Cyan(acc.login, acc.password, acc.url)
	//	fmt.Println(*acc)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("Invalid_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString) // без http, https будет ошибка
	if err != nil {
		return nil, errors.New("Invalid_URL")
	}
	newAcc := &AccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}
	field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login") // читаем поле login
	fmt.Println(string(field.Tag))                                 //  выводим tag в виде строки
	if password == "" {
		newAcc.generatePassword(12)
		//		newAcc.account.generatePassword(12) // запись аналогичная предыдущей
	}
	return newAcc, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-!*")
