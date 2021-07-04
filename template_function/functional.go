package templatefunction

import (
	"fmt"
	"reflect"
)

func getFunction(f interface{}) interface{} {
	rf := reflect.ValueOf(f)

	switch rf.Kind() {
	case reflect.Func:
		return f
	case reflect.String:
		strf := f.(string)
		if v, ok := funcMap[strf]; ok {
			return v
		}
	}

	panic(fmt.Sprint("?", rf.Kind(), f))
}

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

func transform(transformFunction func(functionType reflect.Type, collection reflect.Value) transformReducer, fn interface{}, args ...interface{}) interface{} {
	argsLength := len(args)
	if argsLength < 0 {
		panic("Map must args > 0")
	}

	callFunction := reflect.ValueOf(getFunction(fn))
	inputArgsCount := argsLength - 1
	collection := reflect.ValueOf(args[inputArgsCount])
	callArgs := make([]reflect.Value, inputArgsCount)
	for i := 0; i < inputArgsCount; i++ {
		callArgs[i] = reflect.ValueOf(args[i])
	}

	functionType := callFunction.Type()

	var iter collectionIterable
	var reducer transformReducer = transformFunction(functionType, collection)

	switch collection.Kind() {
	default:
		return callFunction.Call([]reflect.Value{collection})
	case reflect.Map:
		iter = collection.MapRange()
	case reflect.Array,
		reflect.Slice:
		iter = newReflectSliceKeyValueIterator(collection)
	}

	numIn := functionType.NumIn()
	for iter.Next() {
		key := iter.Key()
		val := iter.Value()
		var callResult []reflect.Value
		switch numIn {
		case inputArgsCount:
			callResult = callFunction.Call(callArgs)
		case inputArgsCount + 1:
			callResult = callFunction.Call(append(callArgs, val))
		case inputArgsCount + 2:
			callResult = callFunction.Call(append(callArgs, key, val))
		case 0:
			callResult = callFunction.Call([]reflect.Value{})
		case 1:
			callResult = callFunction.Call([]reflect.Value{val})
		case 2:
			callResult = callFunction.Call([]reflect.Value{key, val})
		default:
			panic(fmt.Sprint("unknown in ", functionType.NumIn()))
		}
		reducer.Add(callResult, key, val)
	}

	return reducer.Get()
}

// Map 함수와 컬렉션을 인자로 받고 각 컬렉션의 원소에 대해서
// 함수를 호출한 값을 새로운 slice를 만들어서 반환합니다.
func Map(fn interface{}, args ...interface{}) interface{} {
	r := func(functionType reflect.Type, collection reflect.Value) transformReducer {
		numOut := functionType.NumOut()
		collectionSize := collection.Len()
		switch numOut {
		default:
			panic(fmt.Sprint("must return 1 or 2 ", numOut))
		case 1:
			return &mapSliceReducer{slice: make([]interface{}, 0, collectionSize)}
		case 2:
			return &mapMapReducer{m: make(map[interface{}]interface{}, collectionSize)}
		}
	}
	return transform(r, fn, args...)
}

// Filter 함수와 컬렉션을 인자로 받고 각 컬렉션의 원소에 대해서
// 함수를 호출한 값을 새로운 slice를 만들어서 반환합니다.
func Filter(fn interface{}, args ...interface{}) interface{} {
	r := func(functionType reflect.Type, collection reflect.Value) transformReducer {
		collectionSize := collection.Len()
		switch collection.Kind() {
		default:
			panic(fmt.Sprint("cannot ", collection.Kind()))
		case reflect.Array,
			reflect.Slice:
			return &filterSliceReducer{slice: make([]interface{}, 0, collectionSize)}
		case reflect.Map:
			return &filterMapReducer{m: make(map[interface{}]interface{}, collectionSize)}
		}
	}
	return transform(r, fn, args...)
}
