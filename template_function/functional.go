package templatefunction

import (
	"reflect"
)

// Select 이름과 값을 입력받고
// 1. 만약 v가 구조체, map 일때는 인자로 받은 이름을 가진 원소를 가져옵니다.
// map은 키가 string 일때만 지원합니다.
// 2. 만약 v가 배열일때는 재귀적으로 모든 원소들에 대해서 1을 수행합니다.
func Select(name string, v interface{}) interface{} {
	k := reflect.ValueOf(name)
	r := reflect.ValueOf(v)

	switch r.Kind() {
	case reflect.Ptr:
		return Select(name, r.Elem().Interface())
	case reflect.Map:
		return r.MapIndex(k).Interface()
	case reflect.Struct:
		e := r.FieldByName(name)
		if !e.IsValid() {
			panic("cannot found " + name)
		}
		return e.Interface()
	case reflect.Slice,
		reflect.Array:
		l := r.Len()
		newArray := make([]interface{}, l)
		for i := 0; i < l; i++ {
			e := r.Index(i).Interface()
			newArray[i] = Select(name, e)
		}
		return newArray
	}
	return v
}

func Case(v interface{}, clauses ...interface{}) interface{} {
	l := len(clauses)

	if l == 0 {
		panic("clauses (compare, return)")
	}

	if l == 1 {
		return clauses[0]
	}

	for i := 0; i < l; i += 2 {
		if reflect.DeepEqual(v, clauses[i]) {
			if i+1 < l {
				return clauses[i+1]
			}
		}
	}

	if l%2 == 0 {
		panic("no matching clause")
	}
	return clauses[l-1]
}

// RemoveLast 마지막 원소를 제거하고 반환합니다
func RemoveLast(v interface{}) interface{} {
	r := reflect.ValueOf(v)
	l := r.Len()

	switch r.Kind() {
	case reflect.String:
		if l == 0 {
			return ""
		}
		s := r.String()
		return s[:l-1]
	case reflect.Array,
		reflect.Slice:
		s := make([]interface{}, l-1)
		for i := 0; i < l-1; i++ {
			e := r.Index(i).Interface()
			s[i] = e
		}
		return s
	}
	panic("not support type")
}

// RemoveFirst 첫 원소를 제거하고 반환합니다
func RemoveFirst(v interface{}) interface{} {
	r := reflect.ValueOf(v)
	l := r.Len()

	switch r.Kind() {
	case reflect.String:
		if l == 0 {
			return ""
		}
		s := r.String()
		return s[:l-1]
	case reflect.Array,
		reflect.Slice:
		s := make([]interface{}, l-1)
		for i := 1; i < l; i++ {
			e := r.Index(i).Interface()
			s[i-1] = e
		}
		return s
	}

	panic("not support type")
}

// Map 함수와 컬렉션을 인자로 받고 각 컬렉션의 원소에 대해서
// 함수를 호출한 값을 새로운 slice를 만들어서 반환합니다.
func Map(fn interface{}, collection interface{}) interface{} {
	f := reflect.ValueOf(fn)
	r := reflect.ValueOf(collection)

	if f.Kind() == reflect.String {
		f = reflect.ValueOf(funcMap[fn.(string)])
	}

	switch r.Kind() {
	case reflect.Map:
		l := r.Len()
		iter := r.MapRange()
		functionType := f.Type()
		if functionType.NumOut() == 2 {
			s := make(map[interface{}]interface{})
			for iter.Next() {
				key := iter.Key()
				val := iter.Value()
				mapResult := f.Call([]reflect.Value{key, val})
				s[mapResult[0].Interface()] = mapResult[1].Interface()
			}
			return s
		} else if functionType.NumOut() == 1 {
			s := make([]interface{}, l)
			for i := 0; iter.Next(); i++ {
				key := iter.Key()
				val := iter.Value()
				var callResult []reflect.Value
				switch functionType.NumIn() {
				case 0:
					callResult = f.Call([]reflect.Value{})
				case 1:
					callResult = f.Call([]reflect.Value{val})
				case 2:
					callResult = f.Call([]reflect.Value{key, val})
				}
				s[i] = callResult[0].Interface()
			}
			return s
		} else {
			panic("must return")
		}
	case reflect.Array,
		reflect.Slice:
		l := r.Len()
		s := make([]interface{}, l)
		for i := 0; i < l; i++ {
			e := r.Index(i)
			mapResult := f.Call([]reflect.Value{e})
			s[i] = mapResult[0].Interface()
		}
		return s
	}
	return f.Call([]reflect.Value{r})
}
