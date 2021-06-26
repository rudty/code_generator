package templatefunction

import "reflect"

// Select 이름과 값을 입력받고
// 1. 만약 v가 구조체, map 일때는 인자로 받은 이름을 가진 원소를 가져옵니다.
// map은 키가 string 일때만 지원합니다.
// 2. 만약 v가 배열일때는 재귀적으로 모든 원소들에 대해서 1을 수행합니다.
func Select(name string, v interface{}) interface{} {
	k := reflect.ValueOf(name)
	r := reflect.ValueOf(v)

	switch r.Kind() {
	case reflect.Map:
		return r.MapIndex(k).Interface()
	case reflect.Struct:
		return r.FieldByName(name).Interface()

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

	panic("not support type")
}
