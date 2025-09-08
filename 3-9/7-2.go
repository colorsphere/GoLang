package main // MAP
import "fmt"

func main() {
	// ключ - значение
	m := map[string]string{
		"PurpleSchool": "https://purpleschool.ru",
	}
	fmt.Println(m)
	fmt.Println(m["PurpleSchool"])
	m["PurpleSchool"] = ".ru"
	fmt.Println(m["PurpleSchool"])
	m["Google"] = "google.com"
	m["Yandex"] = "yandex.ru"
	fmt.Println(m)
	delete(m, "Yandex")
	fmt.Println(m)
}
