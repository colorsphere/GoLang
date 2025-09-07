package main // MAP задачка
import "fmt"

func main() {
	var menu int
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for menu != 4 {
		fmt.Println("1 - Показать закладки")
		fmt.Println("2 - Добавить закладку")
		fmt.Println("3 - Удалить закладку")
		fmt.Println("4 - Выход")
		fmt.Scan(&menu)
		switch menu {
		case 1:
			viewMap(m)
		case 2:
			key, value := addMap()
			m[key] = value
			viewMap(m)
		case 3:
			delete(m, deleteMap())
			viewMap(m)
		default:
			continue
		}
	}

}

func viewMap(m map[string]int) {
	if len(m) == 0 {
		fmt.Println("Закладок нет")
		return
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
	return
}

func addMap() (newKey string, newValue int) {
	fmt.Print("Введите ключ: ")
	fmt.Scan(&newKey)
	fmt.Print("Введите значение: ")
	fmt.Scan(&newValue)
	return
}

func deleteMap() (key string) {
	fmt.Print("Введите ключ: ")
	fmt.Scan(&key)
	return
}
