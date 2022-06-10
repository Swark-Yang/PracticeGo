package main

import (
	"LittleProj/tools"
	"fmt"
)

func main() {
	tools.Say("Hi tools")
	var TestBook = tools.NewAddressBook(3)
	TestBook = TestBook.InsertPersonIdx(tools.People{Name: "JASON", Age: 18, Phone: "123"}, 0)
	TestBook = TestBook.InsertPersonIdx(tools.People{Name: "JASON", Age: 18, Phone: "123"}, 1)
	//fmt.Printf("after inserted len=%d cap=%d\n", len(TestBook), cap(TestBook))
	for i := 2; i < 18; i++ {
		TestBook = TestBook.InsertPersonIdx(tools.People{Name: "SWARK", Age: i, Phone: "123"}, i)
	}
	//fmt.Printf("after inserted len=%d cap=%d\n", len(TestBook), cap(TestBook))
	fmt.Println(TestBook)
}
