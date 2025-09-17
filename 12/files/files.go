package files

import (
	"11/output"
	"fmt"
	"os"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename) // читаем список байтов
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError(err)
	}
	_, err = file.Write(content) // первая переменная - длина в байтах
	defer file.Close()           // Операция, помеченная defer, помещается в конец стека вызовов. Таким образом, она выполняется последней, когда функция завершает свою работу.
	if err != nil {
		file.Close()
		output.PrintError(err)
		return
	}
	fmt.Println("Запись успешна")
}
