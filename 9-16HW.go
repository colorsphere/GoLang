package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func (bin *Bin) readBin() {

}

func main() {
	BinList := []Bin{}
	fmt.Println(BinList)
}
