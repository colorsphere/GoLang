package files

import (
	"errors"
	"fmt"
	"os"
)

func ReadFile() {

}

func ExistsFile(fileName string) bool {
	_, err := os.Stat(fileName)
	if err == nil {
		fmt.Println("Работаем с существующим файлом")
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Создан ноывый пустой файл")
		os.Create(fileName)
		return false
	}
	return false
}

// func WriteFile(content, name string) {
// 	file, err := os.Create(name)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	_, err = file.WriteString(content) // первая переменная - длина в байтах
// 	if err != nil {
// 		file.Close()
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Запись успешна")
// 	file.Close()
// }
