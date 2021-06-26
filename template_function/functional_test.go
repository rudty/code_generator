package templatefunction

import (
	"fmt"
	"reflect"
	"testing"
)

type T1 struct {
	A int
	B int
}

func TestSelectStruct(t *testing.T) {
	s := T1{
		A: 3,
		B: 2,
	}
	a := Select("A", s)
	if a != 3 {
		t.Error(fmt.Sprint("T1/", a))
	}
}

func TestSelectMap(t *testing.T) {
	s := map[string]interface{}{
		"A": "3",
		"B": "2",
	}
	a := Select("A", s)
	if a != "3" {
		t.Error(fmt.Sprint("T1/", a))
	}
}

func TestSelectSlice(t *testing.T) {
	s := []T1{
		{A: 1, B: 2},
		{A: 3, B: 4},
	}
	a := Select("A", s)
	r := reflect.ValueOf(a)
	if r.Len() != 2 {
		t.Error("len 2")
	}

	if r.Index(0).Interface().(int) != 1 {
		t.Error("a[0] = 1")
	}

	if r.Index(1).Interface().(int) != 3 {
		t.Error("a[1] = 3")
	}
}
