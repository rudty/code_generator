package templatefunction

import (
	"fmt"
	"reflect"
	"strings"
)

// ToString 값을 문자열로 변환합니다
func ToString(v interface{}) string {
	return Join(",", v)
}

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
	name := ToString(v)
	if len(name) == 0 {
		return ""
	}

	b := strings.Builder{}
	firstChar := toUpperByte(name[0])
	b.WriteByte(firstChar)

	return removeUnderLineInternal(&b, name)
}

func ToCamel(v interface{}) string {
	name := ToString(v)
	if len(name) == 0 {
		return ""
	}

	b := strings.Builder{}
	firstChar := toLowerByte(name[0])
	b.WriteByte(firstChar)

	return removeUnderLineInternal(&b, name)
}

func ToSnake(v interface{}) string {
	name := ToString(v)
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

// Contains 입력받은 인자를 문자열로 변환한뒤 해당 문자열이 포함되었는지 검사합니다
func Contains(sub interface{}, v interface{}) bool {
	substr := ToString(sub)
	str := ToString(v)
	return strings.Contains(str, substr)
}

// ContainsThen [(비교, 반환)..., 대상 문자열] 을 입력받고
// 해당 문자열을 가지고 있다면 쌍을 반환합니다.
// 값이 문자열이 아니면 모두 문자열로 변경 후 비교를 수행합니다
func ContainsThen(v ...interface{}) string {
	l := len(v)

	if l == 0 {
		return ""
	}

	if l == 1 {
		return ToString(v[0])
	}

	if l%2 == 0 {
		panic("ContainsThen (comp, return)... , str")
	}

	str := ToString(v[l-1])
	v = v[:l-1]
	for i := 0; i < len(v); i++ {
		if Contains(v[i], str) {
			ret := ToString(v[i+1])
			return ret
		}
	}
	panic("no matching value")
}

// Space 지정한 숫자만큼 공백을 추가합니다
func Space(v int) string {
	return strings.Repeat(" ", v)
}
