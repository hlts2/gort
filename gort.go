package gort

import (
	"reflect"
	"sort"
)

// Sort returns sorted array or slice object
func Sort(object interface{}, len int, less func(i, j int) bool) {
	switch reflect.TypeOf(object).Elem().Kind() {
	case reflect.Array:
		rvArray := reflect.Indirect(reflect.ValueOf(&object)).Elem().Elem()

		if len < 2 || len > rvArray.Len() {
			return
		}

		rvSlice := rvArray.Slice(0, len)

		sort.Slice(rvSlice.Interface(), less)

	case reflect.Slice:
		rv := reflect.Indirect(reflect.ValueOf(&object)).Elem().Elem()
		sort.Slice(rv.Slice(0, len).Interface(), less)
	}
}