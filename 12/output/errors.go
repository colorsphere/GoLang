package output // дженерики
import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(value any) {
	//	fmt.Println(value)
	intval, ok := value.(int)
	if ok {
		fmt.Println("int", intval, ok)
	}
	strval, ok := value.(string)
	if ok {
		fmt.Println("string", strval, ok)
	}
	errval, ok := value.(error)
	if ok {
		fmt.Println("error", errval, ok)
	}
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

func sum[T int | float32 | float64 | int16 | int32 | string, V int](a, b T, c V) T { // так описывается дженерик
	return a + b
}
