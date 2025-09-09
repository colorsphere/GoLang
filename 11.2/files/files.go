package files

import (
	"fmt"
	"os"
)

func ReadFile() {

}

func WriteFile(content, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(content) // первая переменная - длина в байтах
	defer file.Close()                 // Операция, помеченная defer, помещается в конец стека вызовов. Таким образом, она выполняется последней, когда функция завершает свою работу.
	if err != nil {
		file.Close()
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}
