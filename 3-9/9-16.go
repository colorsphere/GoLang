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
	list []Bin
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
	return &BinList{
		list: []Bin{
			{
				id:        "2",
				private:   false,
				createdAt: time.Now(),
				name:      "Ivan",
			},
		},
	}
}

func main() {
	fmt.Println(newBin())
	fmt.Println(newBinList())
}
