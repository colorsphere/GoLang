package account // здесь создается слайс аккаунтов

import (
	"11/files"
	"encoding/json"

	"github.com/fatih/color"
)

type Vault struct {
	Bins []Bin `json:"bins"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Bins: []Bin{},
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &Vault{
			Bins: []Bin{},
		}
	}
	return &vault
}

func (vault *Vault) AddBin(acc Bin) {
	vault.Bins = append(vault.Bins, acc)
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	files.WriteFile(data, "data.json")
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}
