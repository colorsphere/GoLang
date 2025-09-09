package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name) // читаем список байтов
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content) // первая переменная - длина в байтах
	defer file.Close()           // Операция, помеченная defer, помещается в конец стека вызовов. Таким образом, она выполняется последней, когда функция завершает свою работу.
	if err != nil {
		file.Close()
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}
