package main

import (
	"fmt"
	"work/data"
	"work/files"
)

func main() {
	variable := data.NewJsonWork()
	fmt.Println(variable)
	fileName := files.NewBase("test.json")
	fmt.Println(fileName)
	// jsonBytes, err := files.ToByte(variable)
	// if err == nil {
	// 	fmt.Println(jsonBytes)
	// } else {
	// 	fmt.Println(err)
	// }
	// jsonList, err := files.ToList(&jsonBytes)
	// if err == nil {
	// 	fmt.Println(jsonList)
	// } else {
	// 	fmt.Println(err)
	// }
}
