package common

import (
	"reflect"
	"errors"
	"strconv"
)

// 判断obj是否在target中，target支持的类型arrary,slice,map
func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("not in array")
}

func IntTo64(input int) int64 {
	string := strconv.Itoa(input)
	out, _ := strconv.ParseInt(string, 10, 64)
	return out
}
