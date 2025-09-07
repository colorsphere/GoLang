package main // MAP
import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for key, value := range m {
		fmt.Println(key, value)
	}
	return
}
