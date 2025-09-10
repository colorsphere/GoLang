package storage // JSON (маппинг полей)
import (
	"errors"
)

type Bin struct {
	Name     string `json:"name"`     // для экспорта полей их нужно назвать с большой буквы
	LastName string `json:"lastname"` // описываем теги так, как они должны быть в JSON - по правилам с маленькой буквы
	Date     string `json:"date"`
}

func (acc *Bin) Output() { // методы прописываются сразу после объявления структуры
	//	fmt.Println(*acc)
}

func NewBin(name, lastName, date string) (*Bin, error) {
	if name == "" {
		return nil, errors.New("Invalid_LOGIN")
	}
	newBin := &Bin{
		Name:     name,
		LastName: lastName,
		Date:     date,
	}
	return newBin, nil
}
