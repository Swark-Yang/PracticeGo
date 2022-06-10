package lib

import (
	"fmt"
	"testing"
)

var lst List[int]

func BenchmarkPush(b *testing.B) {
	lst.Push(15)
	a := lst.GetAll()
	fmt.Println(len(a))
	if len(a) != 1 {
		b.Error("push error")
	}
	fmt.Println(a[0])
	if a[0] != 5 {
		b.Error("push 222")
	}

}
