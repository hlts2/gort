package gort

import (
	"reflect"
	"sort"
)

// Sort array or slice object
func Sort(object interface{}, len int, less func(i, j int) bool) {
	switch reflect.TypeOf(object).Elem().Kind() {
	case reflect.Array:
		v := reflect.ValueOf(&object)
		rvArray := reflect.Indirect(v).Elem().Elem()

		if len < 2 || len > rvArray.Len() {
			return
		}

		sort.Slice(rvArray.Slice(0, len).Interface(), less)

	case reflect.Slice:
		v := reflect.ValueOf(&object)
		rvSlice := reflect.Indirect(v).Elem().Elem()
		sort.Slice(rvSlice.Slice(0, len).Interface(), less)
	}
}
