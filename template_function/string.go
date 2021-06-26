package templatefunction

import (
	"fmt"
	"reflect"
	"strings"
)

// Repeat 인자로 받은 각체 만큼 반복합니다
// 숫자가 들어왔을때는 숫자를 해당 숫자만큼 반복하고
// 컬렉션이 들어왔을때는 컬렉션의 길이만큼 반복합니다
func Repeat(v interface{}, s string) string {
	var count int
	var t = reflect.ValueOf(v)

	switch t.Kind() {
	case reflect.Array,
		reflect.Map,
		reflect.Chan,
		reflect.Slice,
		reflect.String:
		count = t.Len()
	case reflect.Float32,
		reflect.Float64:
		count = int(t.Float())
	default:
		count = int(t.Int())
	}

	return strings.Repeat(s, count)
}

// Join 구분 문자열과 값들을 입력 받고
// 값 문자열사이에 구분 문자열을 넣어 새로운 문자열을 만듭니다
func Join(sep string, v interface{}) string {
	b := strings.Builder{}
	r := reflect.ValueOf(v)

	switch r.Kind() {
	case reflect.Array,
		reflect.Slice:
		l := r.Len()
		for i := 0; i < l; i++ {
			e := r.Index(i).Interface()
			b.WriteString(fmt.Sprint(e))
			b.WriteString(sep)
		}
	case reflect.Map:
		iter := r.MapRange()
		for iter.Next() {
			mk := iter.Key().Interface()
			mv := iter.Value().Interface()
			b.WriteString(fmt.Sprint(mk))
			b.WriteString(sep)
			b.WriteString(fmt.Sprint(mv))
			b.WriteString(sep)
		}
	}
	return b.String()
}

// RemoveLast 마지막 원소를 제거하고 값을 문자열로 만들어서 반환합니다
func RemoveLast(v interface{}) string {
	r := reflect.ValueOf(v)
	l := r.Len()
	if l == 0 {
		return ""
	}

	switch r.Kind() {
	case reflect.String:
		s := r.String()
		return s[:l-1]
	case reflect.Array,
		reflect.Slice:
		b := strings.Builder{}
		for i := 0; i < l-1; i++ {
			e := r.Index(i).Interface()
			b.WriteString(fmt.Sprint(e))
		}
		return b.String()
	}
	return ""
}

// RemoveLast 마지막 원소를 제거하고 값을 문자열로 만들어서 반환합니다
func RemoveFirst(v interface{}) string {
	r := reflect.ValueOf(v)
	l := r.Len()
	if l == 0 {
		return ""
	}

	switch r.Kind() {
	case reflect.String:
		s := r.String()
		return s[1:]
	case reflect.Array,
		reflect.Slice:
		b := strings.Builder{}
		for i := 1; i < l; i++ {
			e := r.Index(i).Interface()
			b.WriteString(fmt.Sprint(e))
		}
		return b.String()
	}
	return ""
}
