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

type wrapperCollection interface {
	Add(r []reflect.Value)
	Get() interface{}
}

type sliceWrapper struct {
	slice []interface{}
}

func (s *sliceWrapper) Add(r []reflect.Value) {
	v := r[0].Interface()
	s.slice = append(s.slice, v)
}

func (s *sliceWrapper) Get() interface{} {
	return s.slice
}

type mapWrapper struct {
	m map[interface{}]interface{}
}

func (s *mapWrapper) Add(r []reflect.Value) {
	k := r[0].Interface()
	v := r[1].Interface()
	s.m[k] = v
}

func (s *mapWrapper) Get() interface{} {
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
	return reflect.Value{}
}

func (r *reflectSliceKeyValueIterator) Value() reflect.Value {
	return r.slice.Index(r.i)
}
