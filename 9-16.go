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

type BinList struct {
	list Bin
}

func newBin() *Bin {
	newBin := &Bin{
		id:        "1",
		private:   true,
		createdAt: time.Now(),
		name:      "John",
	}
	return newBin
}
func newBinList() *BinList {
	newBinList := &BinList{
		list: Bin{
			id:        "2",
			private:   false,
			createdAt: time.Now(),
			name:      "Ivan",
		},
	}
	return newBinList
}

func main() {
	fmt.Println(newBin())
	fmt.Println(newBinList())
}
