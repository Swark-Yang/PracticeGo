package main

import (
	"fmt"
	"linklist/lib"
)

func main() {
	lst := lib.List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	lst.Push(16)
	lst.Push(17)
	fmt.Println("list:", lst.GetAll())
}
