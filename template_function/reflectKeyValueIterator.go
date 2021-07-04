package templatefunction

// 같은 인터페이스로 slice 와 map을 순회하기 위해서 만든 구조
// 복붙해도 상관없을것 같긴 한데..
import (
	"reflect"
)

type collectionIterable interface {
	Next() bool
	Key() reflect.Value
	Value() reflect.Value
}

type transformReducer interface {
	Add(returnValue []reflect.Value, key, val reflect.Value)
	Get() interface{}
}

type mapSliceReducer struct {
	slice []interface{}
}

func (s *mapSliceReducer) Add(returnValue []reflect.Value, key, val reflect.Value) {
	v := returnValue[0].Interface()
	s.slice = append(s.slice, v)
}

func (s *mapSliceReducer) Get() interface{} {
	return s.slice
}

type mapMapReducer struct {
	m map[interface{}]interface{}
}

func (s *mapMapReducer) Add(returnValue []reflect.Value, key, val reflect.Value) {
	k := returnValue[0].Interface()
	v := returnValue[1].Interface()
	s.m[k] = v
}

func (s *mapMapReducer) Get() interface{} {
	return s.m
}

func isOk(ret []reflect.Value) bool {
	if len(ret) == 0 {
		return false
	}
	r := ret[0].Interface()
	if b, ok := r.(bool); ok {
		return b
	}

	if i, ok := r.(int); ok {
		return i != 0
	}

	if s, ok := r.(string); ok {
		return len(s) > 0 && s != "false" && s != "0"
	}

	return false
}

type filterSliceReducer struct {
	slice []interface{}
}

func (s *filterSliceReducer) Add(returnValue []reflect.Value, key, val reflect.Value) {
	if isOk(returnValue) {
		v := val.Interface()
		s.slice = append(s.slice, v)
	}
}

func (s *filterSliceReducer) Get() interface{} {
	return s.slice
}

type filterMapReducer struct {
	m map[interface{}]interface{}
}

func (s *filterMapReducer) Add(returnValue []reflect.Value, key, val reflect.Value) {
	if isOk(returnValue) {
		k := key.Interface()
		v := val.Interface()
		s.m[k] = v
	}
}

func (s *filterMapReducer) Get() interface{} {
	return s.m
}

type reflectSliceKeyValueIterator struct {
	slice reflect.Value
	i     int
	len   int
}

func newReflectSliceKeyValueIterator(slice reflect.Value) *reflectSliceKeyValueIterator {
	return &reflectSliceKeyValueIterator{
		slice: slice,
		i:     -1,
		len:   slice.Len(),
	}
}

func (r *reflectSliceKeyValueIterator) Next() bool {
	r.i += 1
	n := r.i
	return n < r.len
}

func (r *reflectSliceKeyValueIterator) Key() reflect.Value {
	return reflect.ValueOf(r.i)
}

func (r *reflectSliceKeyValueIterator) Value() reflect.Value {
	return r.slice.Index(r.i)
}
