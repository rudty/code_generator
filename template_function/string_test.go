package templatefunction

import (
	"testing"
)

func TestRepeatInt(t *testing.T) {
	a := Repeat(3, "a")
	if a != "aaa" {
		t.Error("aaa/" + a)
	}
}

func TestRepeatFloat(t *testing.T) {
	a := Repeat(3.8, "a")
	if a != "aaa" {
		t.Error("aaa/" + a)
	}
}

func TestRepeatArray(t *testing.T) {
	var r = [5]int{1, 2, 3, 4, 5}
	a := Repeat(r, "a")
	if a != "aaaaa" {
		t.Error("aaaaa/" + a)
	}
}

func TestRepeatSlice(t *testing.T) {
	var r = []int{1, 2, 3, 4, 5}
	a := Repeat(r, "a")
	if a != "aaaaa" {
		t.Error("aaaaa/" + a)
	}
}

func TestRepeatMap(t *testing.T) {
	var r = map[string]interface{}{
		"hello":  "world",
		"string": "value",
		"repeat": "test",
	}
	a := Repeat(r, "a")
	if a != "aaa" {
		t.Error("aaa/" + a)
	}
}

func TestJoinSlice(t *testing.T) {
	a := Join(".", []int{1, 2, 3, 4, 5})
	if a != "1.2.3.4.5." {
		t.Error("1.2.3.4.5./" + a)
	}
}

func TestJoinMap(t *testing.T) {
	a := Join(".", map[string]interface{}{
		"hello": "world",
	})
	if a != "hello.world." {
		t.Error("hello.world./" + a)
	}
}

func TestRemoveLast(t *testing.T) {
	a := RemoveLast("12345")
	if a != "1234" {
		t.Error("1234/" + a)
	}
}

func TestRemoveLastSlice(t *testing.T) {
	a := RemoveLast([]int{100, 200, 300, 400, 500})
	if a != "100200300400" {
		t.Error("100200300400/" + a)
	}
}

func TestRemoveFirst(t *testing.T) {
	a := RemoveFirst("12345")
	if a != "2345" {
		t.Error("2345/" + a)
	}
}

func TestRemoveFirstSlice(t *testing.T) {
	a := RemoveFirst([]int{100, 200, 300, 400, 500})
	if a != "200300400500" {
		t.Error("200300400500/" + a)
	}
}
