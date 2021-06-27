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
			if i > 0 {
				b.WriteString(sep)
			}
			e := r.Index(i).Interface()
			b.WriteString(fmt.Sprint(e))
		}
	case reflect.Map:
		iter := r.MapRange()
		for i := 0; iter.Next(); i++ {
			mk := iter.Key().Interface()
			mv := iter.Value().Interface()
			if i > 0 {
				b.WriteString(sep)
			}
			b.WriteString(fmt.Sprint(mk))
			b.WriteString(sep)
			b.WriteString(fmt.Sprint(mv))
		}
	case reflect.String:
		return v.(string)
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

// Concat 인자로 들어온 값들을 모두 더합니다
func Concat(v ...interface{}) string {
	return Join("", v)
}

func toUpperByte(c byte) byte {
	if 'a' <= c && c <= 'z' {
		c -= 'a' - 'A'
	}
	return c
}

func toLowerByte(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		c += 'a' - 'A'
	}
	return c
}

func removeUnderLineInternal(b *strings.Builder, name string) string {
	toUpper := false

	for i := 1; i < len(name); i++ {
		switch name[i] {
		case 45: // '-'
		case 95: // '_'
			toUpper = true
		default:
			elem := name[i]
			if toUpper {
				elem = toUpperByte(elem)
				toUpper = false
			}
			b.WriteByte(elem)
		}
	}

	return b.String()
}

func ToPascal(v interface{}) string {
	name := Join(",", v)
	if len(name) == 0 {
		return ""
	}

	b := strings.Builder{}
	firstChar := toUpperByte(name[0])
	b.WriteByte(firstChar)

	return removeUnderLineInternal(&b, name)
}

func ToCamel(v interface{}) string {
	name := Join(",", v)
	if len(name) == 0 {
		return ""
	}

	b := strings.Builder{}
	firstChar := toLowerByte(name[0])
	b.WriteByte(firstChar)

	return removeUnderLineInternal(&b, name)
}

func ToSnake(v interface{}) string {
	name := Join(",", v)
	if len(name) == 0 {
		return ""
	}

	builder := strings.Builder{}
	firstChar := toLowerByte(name[0])
	builder.WriteByte(firstChar)

	for i := 1; i < len(name); i++ {
		if 'A' <= name[i] && name[i] <= 'Z' {
			builder.WriteByte(95)
			builder.WriteByte(toLowerByte(name[i]))
		} else {
			builder.WriteByte(name[i])
		}
	}

	return builder.String()
}
