package output // any type switch

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(value any) {
	//	fmt.Println(value)
	switch t := value.(type) {
	case string:
		fmt.Println("string", t)
	case int:
		fmt.Println("int", t)
	case error:
		color.Red(t.Error())
	default:
		fmt.Println("any", t)
	}
}
