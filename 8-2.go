package main // указатели
import "fmt"

type bookmarkMap = map[string]string

func add(b *string) { // так передается указатель в функцию
	*b = "3"
}

func main() {

	a := "1"
	aPointer := &a // создание указателя
	add(aPointer)  // передача в функцию
	fmt.Println(a)

	bookmarks := bookmarkMap{"a": "1", "b": "2"}
Menu:
	for {
		switch menuChoice() {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			delete(bookmarks, deleteBookmark()) // можно переделать по аналогии с addBookmark
		case 4:
			break Menu
		default:
			continue
		}
	}
}

func printBookmarks(bookmarks bookmarkMap) {
	for key, value := range bookmarks {
		fmt.Println(key, value)
	}
}

func menuChoice() (choice int) {
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&choice)
	return
}

func addBookmark(bookmarks bookmarkMap) { // нет return, т.к. передаем переменную ссылочного типа
	var key string              // и все изменения происходят сразу в ней, возвращать их
	var value string            // не нужно, но тогда приходится заводить доп. переменные,
	fmt.Print("Введите ключ: ") // а не использовать краткую записть функции
	fmt.Scan(&key)
	fmt.Print("Введите значение: ")
	fmt.Scan(&value)
	bookmarks[key] = value

}

func deleteBookmark() (key string) {
	fmt.Print("Введите ключ: ")
	fmt.Scan(&key)
	return

}
