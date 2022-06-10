package tools

import (
	"testing"
)

var TestBook = NewAddressBook(10)

func TestInsertPerson(t *testing.T) {
	TestBook = TestBook.InsertPersonBetter(People{"Swark", 18, "123"})
	if len(TestBook) != 1 {
		t.Errorf("insert error len=%d", len(TestBook))
	}

	if TestBook[0].Name != "Swark" {
		t.Errorf("Swark not insert at book, name=%s", TestBook[0].Name)
	}
}

const test_time = 50000

/*
var TestBook2 = NewAddressBook(10)

func BenchmarkInsertPerson(b *testing.B) {
	for i := 0; i < test_time; i++ {
		TestBook2.InsertPerson(People{"SWARK", 18, "123"})
	}
	//b.Logf("BenchmarkInsertPerson inserted len=%d cap=%d", len(TestBook2), cap(TestBook2))
}

var TestBook3 = NewAddressBook(10)

func BenchmarkInsertPersonBetter(b *testing.B) {
	for i := 0; i < test_time; i++ {
		TestBook3 = TestBook3.InsertPersonBetter(People{"SWARK", 18, "123"})
	}
	//b.Logf("BenchmarkInsertPersonBetter inserted len=%d cap=%d", len(TestBook3), cap(TestBook3))
}

var TestBook5 = NewAddressBook(10)

func BenchmarkInsertPersonIdx(b *testing.B) {
	for i := 0; i < test_time; i++ {
		TestBook5 = TestBook5.InsertPersonIdx(People{"SWARK", 18, "123"}, i)
	}
	//b.Logf("after %d inserted len=%d cap=%d", b.N, len(TestBook), cap(TestBook))
}

var TestBook4 = NewAddressBook(10)

func BenchmarkInsertPersonBetter_randomtimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestBook4 = TestBook4.InsertPersonBetter(People{"SWARK", 18, "123"})
	}
	//b.Logf("BenchmarkInsertPersonBetter after %d inserted len=%d cap=%d", b.N, len(TestBook), cap(TestBook))
}*/

func BenchmarkPrintInt2String01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printInt2String01(100)
	}
}

func BenchmarkPrintInt2String02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printInt2String02(int64(100))
	}
}

func BenchmarkPrintInt2String03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printInt2String03(100)
	}
}
