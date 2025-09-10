package files

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	if strings.Contains(name, "json") {
		data, err := os.ReadFile(name) // читаем список байтов
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return data, nil
	} else {
		return nil, errors.New("Это не json")
	}

}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		file.Close()
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}
