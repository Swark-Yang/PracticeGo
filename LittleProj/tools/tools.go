package tools

import (
	"fmt"
	"strconv"
)

func Say(worlds string) {
	fmt.Println(worlds)
}

type People struct {
	Name  string
	Age   int
	Phone string
}

type addressBook []People

func NewAddressBook(page int) addressBook {
	a := make(addressBook, 0, page)
	return a
}

func (book *addressBook) InsertPerson(p People) {
	*book = append(*book, p)
}

func (book *addressBook) InsertPersonBetter(p People) addressBook {
	if len(*book) < cap(*book) {
		(*book) = (*book)[:len(*book)+1]
		(*book)[len(*book)-1] = p
	} else {
		newBook := make(addressBook, len(*book)+1, len(*book)*3)
		copy(newBook, *book)
		book = nil
		book = &newBook
		(*book)[len(*book)-1] = p
	}
	return *book
}

func (book *addressBook) InsertPersonIdx(p People, i int) addressBook {
	if len(*book) < cap(*book) {
		(*book) = (*book)[:len(*book)+1]
		(*book)[i] = p
	} else {
		newBook := make(addressBook, len(*book)+1, (len(*book) + i))
		copy(newBook, *book)
		book = nil
		book = &newBook
		(*book)[i] = p
	}
	return *book
}

func printInt2String01(num int) string {
	return fmt.Sprintf("%d", num)
}

func printInt2String02(num int64) string {
	return strconv.FormatInt(num, 10)
}
func printInt2String03(num int) string {
	return strconv.Itoa(num)
}
