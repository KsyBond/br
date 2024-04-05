package databus

import (
	"fmt"
	"testing"
)

type Foo struct {
	Data string
}

func TestAdd(T *testing.T) {
	d := New(Foo{})
	d.Set("t0", Foo{Data: "somedata"})
	d.Set("t1", Foo{Data: "some data1"})
	fmt.Println(d.Get("t0"))
	fmt.Println(d.Get("t1"))
}
