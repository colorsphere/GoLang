package main // type alias

import "fmt"

type bookmarkMap = map[string]string

func main() {
	bookmarks := bookmarkMap{"a": "1", "b": "2"}
	for {
		i := true
		switch menuChoice() {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			key, value := addBookmark()
			bookmarks[key] = value
		case 3:
			delete(bookmarks, deleteBookmark())
		case 4:
			i = false
		default:
			continue
		}
		if i == false {
			break
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

func addBookmark() (key, value string) {
	fmt.Print("Введите ключ: ")
	fmt.Scan(&key)
	fmt.Print("Введите значение: ")
	fmt.Scan(&value)
	return

}

func deleteBookmark() (key string) {
	fmt.Print("Введите ключ: ")
	fmt.Scan(&key)
	return

}
