package main // unpack
import (
	"fmt"
	"slices"
)

// При использовании пакета slices задача сильно упрощается

func main() {
	tr1 := []int{1, 2, 3}
	tr2 := []int{4, 5, 6}

	tr3 := append(tr1, tr2...) // ... это распаковка слайса
	tr4 := slices.Concat(tr1, tr2, tr3)
	fmt.Println(tr3)
	fmt.Println(tr4)

}
