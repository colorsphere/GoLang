package account // Интерфейс встроенный

import (
	"11/output"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db Db
}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		// isMatched := strings.Contains(account.Url, url)
		isMatched := checker(account, str)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

// func (vault *VaultWithDb) FindAccountsByLogin(login string) []Account { // нет смысла дублировать функцию
// 	var accounts []Account                                                 // strings.Contains(account.Login, login) - функция,
// 	for _, account := range vault.Accounts {                               // которую нужно подставить вместо strings.Contains(account.Url, url)
// 		isMatched := strings.Contains(account.Login, login)
// 		if isMatched {
// 			accounts = append(accounts, account)
// 		}
// 	}
// 	return accounts
// }

func (vault *VaultWithDb) DeleteAccountsByUrl(url string) bool {
	var accounts []Account
	isDeleted := false

	for _, account := range vault.Accounts { // изменение тут, работаем с индексом
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			//			vault.Accounts = append(vault.Accounts[:i], vault.Accounts[i+1:]...) // так можно сделать, но не нужно - БРЕД
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.save()
	return isDeleted
}

func NewVault(db Db) *VaultWithDb {
	// db := files.NewJsonDb("data.json") // это зависимость - убираем ее отсюда
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Println(err)
		output.PrintError("Не удалось разобрать файл data.json")
		//		color.Red("Не удалось разобрать файл data.json")
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
	}
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte, error) { // методы прописываются сразу после объявления структуры
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		output.PrintError("Не удалось преобразовать")
	}
	// db := files.NewJsonDb("data.json") // меняем на правильное
	vault.db.Write(data)
	// db.Write(data) // убираем лишнее

}
