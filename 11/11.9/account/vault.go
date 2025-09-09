package account // здесо создается слайс аккаунтов

import (
	"encoding/json"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	return &Vault{
		Accounts:  []Account{},
		UpdatedAt: time.Now(),
	}
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
}

func (vault *Vault) ToBytes() ([]byte, error) { // методы прописываются сразу после объявления структуры
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}
