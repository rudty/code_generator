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

func TestSelectSliceElemPointer(t *testing.T) {
	s := []*T1{
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

func TestRemoveLast(t *testing.T) {
	a := RemoveLast("12345")
	if !reflect.DeepEqual(a, "1234") {
		t.Error("100200300400/" + fmt.Sprint(a))
	}
}

func TestRemoveLastSlice(t *testing.T) {
	a := RemoveLast([]int{100, 200, 300, 400, 500})

	if !reflect.DeepEqual(a, []interface{}{100, 200, 300, 400}) {
		t.Error("100200300400/" + fmt.Sprint(a))
	}
}

func TestRemoveFirst(t *testing.T) {
	a := RemoveFirst("12345")
	if !reflect.DeepEqual(a, "1234") {
		t.Error("1234/" + fmt.Sprint(a))
	}
}

func TestRemoveFirstSlice(t *testing.T) {
	a := RemoveFirst([]int{100, 200, 300, 400, 500})
	if !reflect.DeepEqual(a, []interface{}{200, 300, 400, 500}) {
		t.Error("200300400500/" + fmt.Sprint(a))
	}
}

func TestMapArrayArg1(t *testing.T) {
	a := Map(fmt.Sprint, []int{1, 2, 3})
	v := reflect.ValueOf(a)
	if v.Len() != 3 {
		t.Error("length must 3")
	}
	if v.Index(0).Interface().(string) != "1" {
		t.Error("1")
	}
	if v.Index(1).Interface().(string) != "2" {
		t.Error("2")
	}
	if v.Index(2).Interface().(string) != "3" {
		t.Error("3")
	}
}
func TestMapArrayArg2(t *testing.T) {
	a := Map(fmt.Sprintf, "%d", []int{1, 2, 3})
	v := reflect.ValueOf(a)
	if v.Len() != 3 {
		t.Error("length must 3")
	}
	if v.Index(0).Interface().(string) != "1" {
		t.Error("1")
	}
	if v.Index(1).Interface().(string) != "2" {
		t.Error("2")
	}
	if v.Index(2).Interface().(string) != "3" {
		t.Error("3")
	}
}

func TestMapMapArg0Return1(t *testing.T) {
	f := func(a, b int) int {
		return a + b
	}

	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	a := Map(f, m)

	v := reflect.ValueOf(a)
	if v.Len() != 3 {
		t.Error("length must 3")
	}
	if v.Index(0).Interface().(int) != 2 {
		t.Error("2")
	}
	if v.Index(1).Interface().(int) != 4 {
		t.Error("4")
	}
	if v.Index(2).Interface().(int) != 6 {
		t.Error("6")
	}
}

func TestMapMapArg0Return2(t *testing.T) {
	f := func(a, b int) (int, int) {
		return b, a
	}
	m := map[int]int{
		1: 4,
		2: 5,
		3: 6,
	}
	a := Map(f, m)

	v := reflect.ValueOf(a)
	if v.Len() != 3 {
		t.Error("length must 3")
	}

	e1 := v.MapIndex(reflect.ValueOf(4))
	if e1.Interface().(int) != 1 {
		t.Error("1")
	}
	e2 := v.MapIndex(reflect.ValueOf(5))
	if e2.Interface().(int) != 2 {
		t.Error("2")
	}
	e3 := v.MapIndex(reflect.ValueOf(6))
	if e3.Interface().(int) != 3 {
		t.Error("3")
	}
}
